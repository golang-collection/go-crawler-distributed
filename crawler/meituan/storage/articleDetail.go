package storage

import (
	"go-crawler-distributed/elastic/elasticOperation"
	"go-crawler-distributed/model"
	"go-crawler-distributed/service/watchConfig"
)

/**
* @Author: super
* @Date: 2020-09-01 19:29
* @Description:
**/



func StorageArticle(data interface{}) error {

	article := data.(*model.Article)
	index, _ := watchConfig.GetElasticIndex()

	_, err := elasticOperation.SaveInfo(index, article)
	if err != nil{
		return err
	}

	return nil
}