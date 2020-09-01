package parser

import (
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/mq/mqTools"
	"go.uber.org/zap"
	"strings"
)

/**
* @Author: super
* @Date: 2020-09-01 19:09
* @Description:
**/

func ParseArticleDetail(contents []byte, queueName string, url string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		logger.Error("new doc reader error", zap.Error(err))
	}

	//初始化消息队列
	articleUrlList := mqTools.NewRabbitMQSimple(queueName)

	result := dom.Find("a[rel=bookmark]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, exist := selection.Attr("href")
		if exist{
			//将解析到的图书详细信息URL放到消息队列
			//不加延迟会出现问题
			articleUrlList.PublishSimple(href)
		}
	})
}