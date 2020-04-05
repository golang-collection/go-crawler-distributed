package main

import (
	"flag"
	"fmt"
	"go-crawler-distributed/engine"
	"go-crawler-distributed/mylog"
	itemSaver "go-crawler-distributed/persist/client"
	"go-crawler-distributed/rpcsupport"
	"go-crawler-distributed/scheduler"
	worker "go-crawler-distributed/worker/client"
	"go-crawler-distributed/zhenai/parser"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String(
		"worker_hosts", "",
		"workrt hosts(comma separated)")
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	flag.Parse()
	itemChan, err := itemSaver.ItemSaver(
		fmt.Sprintf(":%d", *itemSaverHost))
	if err != nil {
		panic(err)
	}

	pool := createClientPool(
		strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			"ParseCityList"),
	})

}

func createClientPool(host []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range host {
		c, err := rpcsupport.NewClient(h)
		if err != nil {
			clients = append(clients, c)
			mylog.LogInfo("main.createClientPool", fmt.Sprintf("connect %s", h))
		} else {
			mylog.LogError("main.createClientPool", err)
		}
	}
	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()

	return out
}
