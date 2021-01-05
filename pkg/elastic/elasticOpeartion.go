package elastic

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-crawler-distributed/internal/model"
	"reflect"

	"go-crawler-distributed/global"
)

/**
* @Author: super
* @Date: 2020-12-29 11:58
* @Description:
**/
const Mapping = `
{
    "mappings": {
        "properties": {
            "title": {
                "type": "text"
            },
            "url": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            },
            "content": {
                "type": "text"
            }
        }
    }
}`

//判断index是否存在
func IndexExist(index string) (bool, error) {
	client := global.ElasticEngine

	exist, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return exist, err
	}
	if !exist {
		_, err := client.CreateIndex(index).BodyString(Mapping).Do(context.Background())
		if err != nil {
			return false, err
		}
	}
	return true, err
}

//保存信息
func SaveInfo(table string, data interface{}) (string, error) {
	client := global.ElasticEngine
	// https://www.letianbiji.com/elasticsearch/es7-quick-start.html
	// 在v7中Type被注释
	// ES 实例：对应 MySQL 实例中的一个 Database。
	// Index 对应 MySQL 中的 Table
	// Document 对应 MySQL 中表的记录。
	response, err := client.Index().Index(table).BodyJson(data).Do(context.Background())
	if err != nil {
		return "", err
	}
	return response.Id, nil
}

//获取信息
func GetInfo(table string, id string) (*model.Article, error) {
	client := global.ElasticEngine
	result, err := client.Get().Index(table).Id(id).Do(context.Background())
	if err != nil {
		return nil, err
	}
	article := &model.Article{}
	err = article.UnmarshalJSON(result.Source)
	if err != nil {
		return nil, err
	}
	return article, nil
}

//搜索信息
func SearchInfo(table string, fieldName string, fieldValue string) ([]*model.Article, error) {
	query := elastic.NewTermQuery(fieldName, fieldValue)
	client := global.ElasticEngine
	result, err := client.Search().Index(table).Query(query).Do(context.Background())
	if err != nil {
		return nil, err
	}
	articles := make([]*model.Article, 0)
	article := model.Article{}
	total := result.TotalHits()
	if total > 0 {
		for _, item := range result.Each(reflect.TypeOf(article)) {
			if t, ok := item.(model.Article); ok {
				articles = append(articles, &t)
			}
		}
	}
	return articles, nil
}
