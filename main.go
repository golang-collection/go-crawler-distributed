package main

import (
	"errors"
	"flag"
	"go-crawler-distributed/config"
	"go-crawler-distributed/engine"
	"go-crawler-distributed/persist/client"
	"go-crawler-distributed/rpcsupport"
	"go-crawler-distributed/scheduler"
	worker "go-crawler-distributed/worker/client"
	"go-crawler-distributed/zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := client.ItemSaver(
		*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool, err := createClientPool(
		strings.Split(*workerHosts, ","))
	if err != nil {
		panic(err)
	}

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
			config.ParseCityList),
	})
}

func createClientPool(
	hosts []string) (chan *rpc.Client, error) {
	var clients []*rpc.Client
	for _, h := range hosts {
		c, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, c)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf(
				"Error connecting to %s: %v",
				h, err)
		}
	}

	if len(clients) == 0 {
		return nil, errors.New(
			"no connections available")
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out, nil
}
