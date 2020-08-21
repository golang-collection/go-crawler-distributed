package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/worker"
	"go-crawler-distributed/mq/mqTools"
	"go.uber.org/zap"
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

	funcParser := worker.NewFuncParser(parser.ParseBookList, crawerConfig.BookDetailUrl, "tagList")

	go func() {
		logger.Info("Ready to fetching", zap.String("parser name", funcParser.Name))
		for d := range messages {
			go func() {
				url := string(d.Body)
				logger.Info("fetching", zap.String(funcParser.Name, url))
				for i := 0; i <= 1000; i = i + 20 {
					go func() {
						url := url + "?start=" + strconv.Itoa(i) + "&type=T"

						r := worker.Request{
							Url:    url,
							Parser: funcParser,
						}
						worker.Worker(r)
					}()
					time.Sleep(5 * time.Second)
				}
			}()
			time.Sleep(5 * time.Second)
		}
	}()

	<-forever
}
