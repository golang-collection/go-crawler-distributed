package storage

import (
	"go-crawler-distributed/elastic/client"
	"go-crawler-distributed/model"
	"go-crawler-distributed/service/watchConfig"
	"go-crawler-distributed/tools"
)

/**
* @Author: super
* @Date: 2020-09-01 19:29
* @Description:
**/


func StorageArticle(contents []byte) error {
	article := &model.Article{}
	err := article.UnmarshalJSON(contents)
	if err != nil {
		return err
	}
	article.Content = tools.UnzipString(article.Content)

	index, _ := watchConfig.GetElasticIndex()
	_, err = client.SaveInfo(index, article)
	//_, err = elasticOperation.SaveInfo(index, article)
	return err
}