package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/fetcher"
	"go-crawler-distributed/mq/mqTools"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-14 21:22
* @Description:
**/

func main() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BookDetailUrl)
	messages := bookDetailURL.GetMsgs()

	forever := make(chan bool)

	go func() {
		log.Println("Ready to fetching BookDetail")
		for d := range messages {
			go func() {
				url := string(d.Body)
				log.Printf("Fetching BookDetail: %s", url)
				contents, _ := fetcher.Fetch(url)
				parser.ParseBookDetail(contents, crawerConfig.BookDetail, url)
			}()
			time.Sleep(5 * time.Second)
		}
	}()

	<-forever
}
