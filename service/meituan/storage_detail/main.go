package main

import (
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/meituan/storage"
	"go-crawler-distributed/crawler/persistence"
	"go-crawler-distributed/elastic/elasticOperation"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/service/watchConfig"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-01 19:37
* @Description:
**/

func main() {
	articleDetail := mqTools.NewRabbitMQSimple(crawerConfig.ArticleDetail)
	messages := articleDetail.GetMsgs()


	funcStorage := persistence.FuncStorage{
		Name:      "ArticleDetail",
		ParseFunc: storage.StorageArticle,
	}

	index, _ := watchConfig.GetElasticIndex()
	elasticOperation.IndexExist(index)


	unifiedLog.GetLogger().Info("Ready to storage", zap.String("name", funcStorage.Name))
	for d := range messages {
		go func(data []byte) {
			unifiedLog.GetLogger().Info("storage", zap.String(funcStorage.Name, string(data)))

			err := funcStorage.ParseFunc(data)
			if err != nil {
				unifiedLog.GetLogger().Error("storage parse error", zap.Error(err))
			}
		}(d.Body)
		time.Sleep(time.Millisecond * 100)
	}
}