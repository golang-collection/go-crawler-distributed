package main

import (
	"github.com/olivere/elastic"
	"go-crawler-distributed/mylog"
	"go-crawler-distributed/persist"
	"go-crawler-distributed/rpcsupport"
)

func main() {
	err := serverRpc(":1234", "dating_profile")
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
