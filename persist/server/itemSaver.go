package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"go-crawler-distributed/config"
	"go-crawler-distributed/mylog"
	"go-crawler-distributed/persist"
	"go-crawler-distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	err := serverRpc(fmt.Sprintf(":%d", *port),
		config.ElasticIndex)

	if err != nil {
		mylog.LogError("persist server: ServeRpc", err)
		panic(err)
	}
}

func serverRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		mylog.LogError("persist server: NewClient", err)
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
