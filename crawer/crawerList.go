package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/worker"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"strconv"
	"sync"
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
		wg := sync.WaitGroup{}
		unifiedLog.GetLogger().Info("Ready to fetching", zap.String("parser name", funcParser.Name))
		for d := range messages {
			go func(data []byte) {
				defer func() {
					wg.Done()
				}()
				wg.Add(1)
				url := string(data)
				unifiedLog.GetLogger().Info("fetching", zap.String(funcParser.Name, url))
				minWg := sync.WaitGroup{}
				for i := 0; i <= 1000; i = i + 20 {
					go func(i int) {
						defer func() {
							minWg.Done()
						}()
						minWg.Add(1)
						url := url + "?start=" + strconv.Itoa(i) + "&type=T"

						r := worker.Request{
							Url:    url,
							Parser: funcParser,
						}
						worker.Worker(r)
					}(i)
					time.Sleep(time.Second * 2)
				}
				minWg.Wait()
			}(d.Body)
			time.Sleep(time.Second * 2)
		}
		wg.Wait()
	}()

	<-forever
}
