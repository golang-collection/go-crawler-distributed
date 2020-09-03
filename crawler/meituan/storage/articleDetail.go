package storage

import (
	"go-crawler-distributed/elastic/client"
	"go-crawler-distributed/elastic/elasticOperation"
	"go-crawler-distributed/model"
	"go-crawler-distributed/service/watchConfig"
	"go-crawler-distributed/tools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-01 19:29
* @Description:
**/


func StorageArticle(contents []byte, _ string, _ string) {
	article := &model.Article{}
	err := article.UnmarshalJSON(contents)
	if err != nil {
		unifiedLog.GetLogger().Error("article unmarshalJSON error", zap.Error(err))
		return
	}
	article.Content = tools.UnzipString(article.Content)

	index, _ := watchConfig.GetElasticIndex()
	_, _ = elasticOperation.IndexExist(index)

	_, err = client.SaveInfo(index, article)
	//_, err = elasticOperation.SaveInfo(index, article)
	if err != nil {
		unifiedLog.GetLogger().Error("article save info error", zap.Error(err))
	}
}