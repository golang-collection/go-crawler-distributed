package parser

import (
	"context"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crawler/crawerConfig"
	"go-crawler-distributed/pkg/mq"
	"strconv"
)

/**
* @Author: super
* @Date: 2020-09-01 16:00
* @Description:
**/

func ParseArticleList(contents []byte, queueName string, url string) {
	err := mq.Publish(queueName, []byte(url))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}
	global.Logger.Infof(context.Background(), "url: %s", url)

	for i := 2; i < 22; i++ {
		url := "https://tech.meituan.com//page/" + strconv.Itoa(i) + ".html"
		global.Logger.Infof(context.Background(), "url: %s", url)
		err = mq.Publish(queueName, []byte(url))
		if err != nil {
			global.Logger.Error(context.Background(), err)
		}
	}
	err = mq.Publish(queueName, []byte(crawerConfig.StopTAG))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}
}
