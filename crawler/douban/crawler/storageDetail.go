package crawler

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
* @Date: 2020-08-14 16:27
* @Description:
**/

var logger = unifiedLog.GetLogger()


func StorageDetail() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BookDetail)
	messages := bookDetailURL.GetMsgs()


	funcStorage := persistence.FuncStorage{
		Name:      "BookDetail",
		ParseFunc: storage.ParseAndStorage,
	}


	logger.Info("Ready to storage", zap.String("name", funcStorage.Name))
	for d := range messages {
		go func(data []byte) {
			logger.Info("storage", zap.String(funcStorage.Name, string(data)))

			err := funcStorage.ParseFunc(data)
			if err != nil {
				logger.Error("storage parse error", zap.Error(err))
			}
		}(d.Body)
	}
}
