package storage

import (
	"context"

	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/model"
	"go-crawler-distributed/pkg/util"
	"go-crawler-distributed/service/elastic/client"
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
		global.Logger.Error(context.Background(), err)
		return
	}
	article.Content = util.UnzipString(article.Content)


	index := global.ElasticSetting.Index
	_, _ = client.IndexExist(index)
	_, err = client.SaveInfo(index, article)
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}
}
