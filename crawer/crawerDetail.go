package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/worker"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"sync"
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

				r := worker.Request{
					Url:    url,
					Parser: funcParser,
				}

				worker.Worker(r)
			}(d.Body)
		}
		wg.Wait()
	}()

	<-forever
}
