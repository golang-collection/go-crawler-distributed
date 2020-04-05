package main

import (
	"flag"
	"fmt"
	"go-crawler-distributed/rpcsupport"
	"go-crawler-distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {

	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	err := rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port), worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
