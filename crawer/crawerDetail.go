package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/worker"
	"go-crawler-distributed/mq/mqTools"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-14 15:50
* @Description:
**/
func main() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BookDetailUrl)
	messages := bookDetailURL.GetMsgs()

	forever := make(chan bool)

	funcParser := worker.NewFuncParser(parser.ParseBookDetail, crawerConfig.BookDetail, "BookDetail")

	go func() {
		log.Println("Ready to fetching " + funcParser.Name)
		for d := range messages {
			go func() {
				url := string(d.Body)
				log.Printf("Fetching "+funcParser.Name+": %s", url)

				r := worker.Request{
					Url:    url,
					Parser: funcParser,
				}

				worker.Worker(r)

				//contents, _ := fetcher.Fetch(url)
				//parser.ParseBookDetail(contents, crawerConfig.BookDetail, url)
			}()
			time.Sleep(5 * time.Second)
		}
	}()

	<-forever
}
