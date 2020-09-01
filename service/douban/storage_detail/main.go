package main

import (
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/douban/storage"
	"go-crawler-distributed/crawler/persistence"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-31 16:52
* @Description:
**/

func main() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BookDetail)
	messages := bookDetailURL.GetMsgs()


	funcStorage := persistence.FuncStorage{
		Name:      "BookDetail",
		ParseFunc: storage.ParseAndStorage,
	}


	unifiedLog.GetLogger().Info("Ready to storage", zap.String("name", funcStorage.Name))
	for d := range messages {
		go func(data []byte) {
			unifiedLog.GetLogger().Info("storage", zap.String(funcStorage.Name, string(data)))

			err := funcStorage.ParseFunc(data)
			if err != nil {
				unifiedLog.GetLogger().Error("storage parse error", zap.Error(err))
			}
		}(d.Body)
	}
}