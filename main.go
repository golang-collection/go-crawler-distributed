package main

import (
	"fmt"
	"go-crawler-distributed/config"
	"go-crawler-distributed/engine"
	"go-crawler-distributed/persist/client"
	"go-crawler-distributed/scheduler"
	"go-crawler-distributed/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			"ParseCityList"),
	})

}
