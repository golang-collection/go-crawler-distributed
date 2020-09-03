package parser

import (
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/mq/mqTools"
	"go.uber.org/zap"
	"strings"
)

/**
* @Author: super
* @Date: 2020-09-01 18:59
* @Description:
**/

func ParseArticleUrlList(contents []byte, queueName string, _ string) {
	//初始化消息队列
	articleUrlList := mqTools.NewRabbitMQSimple(queueName)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		logger.Error("new doc reader error", zap.Error(err))
	}

	result := dom.Find("a[rel=bookmark]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, exist := selection.Attr("href")
		if exist{
			logger.Info("fetching", zap.String("url", href))
			//将解析到的图书详细信息URL放到消息队列
			//不加延迟会出现问题
			articleUrlList.PublishSimple(href)
		}
	})
}