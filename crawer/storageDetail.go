package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/storage"
	"go-crawler-distributed/crawer/persistence"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-14 16:27
* @Description:
**/
var logger = unifiedLog.GetLogger()


func main() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BookDetail)
	messages := bookDetailURL.GetMsgs()

	forever := make(chan bool)

	funcStorage := persistence.FuncStorage{
		Name:      "BookDetail",
		ParseFunc: storage.ParseAndStorage,
	}


	go func() {
		logger.Info("Ready to storage", zap.String("name", funcStorage.Name))
		for d := range messages {
			go func() {
				logger.Info("storage", zap.String(funcStorage.Name, string(d.Body)))

				err := funcStorage.ParseFunc(d.Body)
				if err != nil {
					logger.Error("storage parse error", zap.Error(err))
				}
			}()
		}
	}()

	<-forever
}
