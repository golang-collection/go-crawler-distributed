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
* @Date: 2020-08-14 15:50
* @Description:
**/
func main() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BOOK_DETAIL_URL)
	messages := bookDetailURL.GetMsgs()

	forever := make(chan bool)

	go func() {
		for d := range messages {
			go func() {
				log.Printf("Fetching BookDetail: %s", d.Body)
				contents, _ := fetcher.Fetch(string(d.Body))
				parser.ParseBookDetail(contents, crawerConfig.BOOK_DETAIL)
			}()
			time.Sleep(5 * time.Second)
		}
	}()

	<-forever
}
