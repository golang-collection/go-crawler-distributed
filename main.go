package main

import (
	"fmt"
	"go-crawler-distributed/config"
	"go-crawler-distributed/engine"
	itemSaver "go-crawler-distributed/persist/client"
	"go-crawler-distributed/scheduler"
	worker "go-crawler-distributed/worker/client"
	"go-crawler-distributed/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	itemChan, err := itemSaver.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor()
	if err != nil {
		panic(err)
	}

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
