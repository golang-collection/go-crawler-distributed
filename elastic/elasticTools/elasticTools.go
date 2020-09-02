package elasticTools

import (
	"errors"
	"github.com/olivere/elastic/v7"
	"go-crawler-distributed/service/watchConfig"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-01 16:27
* @Description:
**/

var client *elastic.Client
var logger = unifiedLog.GetLogger()

func init() {
	url, _ := watchConfig.GetElasticUrl()
	err := pingServer(url)
	if err != nil{
		logger.Error("elastic new client error", zap.Error(err))
		panic(err)
	}
}


func GetClient() *elastic.Client{
	return client
}

// pingServer pings the http server to make sure the router is working.
func pingServer(url string) error {
	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		var err error
		client, err = elastic.NewClient(
			elastic.SetURL(url),
			elastic.SetSniff(false))
		if err == nil{
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the elastic server")
}