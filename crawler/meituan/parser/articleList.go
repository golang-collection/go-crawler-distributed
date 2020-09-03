package parser

import (
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"strconv"
)

/**
* @Author: super
* @Date: 2020-09-01 16:00
* @Description:
**/

var logger = unifiedLog.GetLogger()

func ParseArticleList(contents []byte, queueName string, url string) {
	//初始化消息队列
	articleList := mqTools.NewRabbitMQSimple(queueName)
	articleList.PublishSimple(url)
	logger.Info("fetching", zap.String("url", url))

	for i := 2; i<22;i++{
		url := "https://tech.meituan.com//page/"+ strconv.Itoa(i) +".html"
		logger.Info("fetching", zap.String("url", url))
		articleList.PublishSimple(url)
	}

	articleList.PublishSimple(crawerConfig.StopTAG)
}