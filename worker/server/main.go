package main

import (
	"fmt"
	"go-crawler-distributed/config"
	"go-crawler-distributed/rpcsupport"
	"go-crawler-distributed/worker"
)

func main() {
	err := rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", config.WorkerPort0),
		worker.CrawlService{})
	if err != nil {
		panic(err)
	}
}
