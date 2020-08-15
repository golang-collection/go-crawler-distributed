package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/fetcher"
	"go-crawler-distributed/mq/mqTools"
	"log"
	"strconv"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-12 19:47
* @Description:
**/

func main() {
	tagUrl := mqTools.NewRabbitMQSimple(crawerConfig.TagUrl)
	messages := tagUrl.GetMsgs()

	forever := make(chan bool)

	go func() {
		log.Println("Ready to fetching tagList")
		for d := range messages {
			go func() {
				url := string(d.Body)
				log.Printf("Fetching tagList: %s", url)
				for i := 0; i <= 1000; i = i + 20 {
					go func() {
						contents, _ := fetcher.Fetch(url + "?start=" + strconv.Itoa(i) + "&type=T")
						parser.ParseBookList(contents, crawerConfig.BookDetailUrl)
					}()
					time.Sleep(5 * time.Second)
				}
			}()
			time.Sleep(5 * time.Second)
		}
	}()

	<-forever
}
