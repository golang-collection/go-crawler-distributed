package elasticOperation

import (
	"context"
	"fmt"
	"go-crawler-distributed/crawler/meituan/storage"
	"go-crawler-distributed/elastic/elasticTools"
	"go-crawler-distributed/model"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-01 16:37
* @Description:
**/

//判断index是否存在
func IndexExist(index string){
	client := elasticTools.GetClient()

	exist, err := client.IndexExists(index).Do(context.Background())
	fmt.Println(exist)
	if err != nil{
		unifiedLog.GetLogger().Error("elastic index exist error", zap.Error(err))
		return
	}
	if !exist{
		result, err := client.CreateIndex(index).BodyString(storage.Mapping).Do(context.Background())
		if err != nil{
			unifiedLog.GetLogger().Error("elastic create index error", zap.Error(err))
			return
		}
		fmt.Println(result)
	}
}

//保存信息
func SaveInfo(table string, data model.Article) (string, error){
	client := elasticTools.GetClient()
	// https://www.letianbiji.com/elasticsearch/es7-quick-start.html
	// 在v7中Type被注释
	// ES 实例：对应 MySQL 实例中的一个 Database。
	// Index 对应 MySQL 中的 Table 。
	// Document 对应 MySQL 中表的记录。
	response, err := client.Index().Index(table).BodyJson(data).Do(context.Background())
	if err != nil{
		unifiedLog.GetLogger().Error("elastic save error", zap.Error(err))
		return "", err
	}
	return response.Id, nil
}

//获取信息
func GetInfo(table string, id string) (*model.Article, error) {
	client := elasticTools.GetClient()
	result, err := client.Get().Index(table).Id(id).Do(context.Background())
	if err != nil{
		return nil, err
	}
	article := &model.Article{}
	err = article.UnmarshalJSON(result.Source)
	if err != nil{
		return nil, err
	}
	return article, nil
}
