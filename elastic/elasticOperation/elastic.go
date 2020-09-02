package elasticOperation

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-crawler-distributed/crawler/meituan/conf"
	"go-crawler-distributed/elastic/elasticTools"
	"go-crawler-distributed/model"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"reflect"
)

/**
* @Author: super
* @Date: 2020-09-01 16:37
* @Description:
**/

//判断index是否存在
func IndexExist(index string) (bool, error) {
	client := elasticTools.GetClient()

	exist, err := client.IndexExists(index).Do(context.Background())
	if err != nil{
		unifiedLog.GetLogger().Error("elastic index exist error", zap.Error(err))
		return exist, err
	}
	if !exist{
		_, err := client.CreateIndex(index).BodyString(conf.Mapping).Do(context.Background())
		if err != nil{
			unifiedLog.GetLogger().Error("elastic create index error", zap.Error(err))
			return false, err
		}
	}
	return true, err
}

//保存信息
func SaveInfo(table string, data *model.Article) (string, error){
	client := elasticTools.GetClient()
	// https://www.letianbiji.com/elasticsearch/es7-quick-start.html
	// 在v7中Type被注释
	// ES 实例：对应 MySQL 实例中的一个 Database。
	// Index 对应 MySQL 中的 Table
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

//搜索信息
func SearchInfo(table string, fieldName string, fieldValue string)([]*model.Article, error){
	query := elastic.NewTermQuery(fieldName, fieldValue)
	client := elasticTools.GetClient()
	result, err := client.Search().Index(table).Query(query).Do(context.Background())
	if err != nil{
		return nil, err
	}
	articles := make([]*model.Article, 0)
	article := model.Article{}
	total := result.TotalHits()
	if total > 0{
		for _, item := range result.Each(reflect.TypeOf(article)){
			if t, ok := item.(model.Article); ok {
				articles = append(articles, &t)
			}
		}
	}
	return articles, nil
}