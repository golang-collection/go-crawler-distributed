package elasticTools

import (
	"github.com/olivere/elastic/v7"
	"go-crawler-distributed/service/watchConfig"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-09-01 16:27
* @Description:
**/

var client *elastic.Client
var logger = unifiedLog.GetLogger()

func init() {
	var err error
	url, _ := watchConfig.GetElasticUrl()
	client, err = elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false))
	if err != nil{
		logger.Error("elastic new client error", zap.Error(err))
		panic(err)
	}
}


func GetClient() *elastic.Client{
	return client
}