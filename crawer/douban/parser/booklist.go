package parser

import (
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/cache/client"
	"go-crawler-distributed/mq/mqTools"
	"go.uber.org/zap"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-14 13:54
* @Description:
**/


func ParseBookList(contents []byte, queueName string, _ string) {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		logger.Error("new doc reader error", zap.Error(err))
	}

	//初始化消息队列
	bookDetailURL := mqTools.NewRabbitMQSimple(queueName)

	result := dom.Find("a[title]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		logger.Info("fetching", zap.String("url", href))

		//redis去重
		boolean, _ := client.ElementIsInSet(queueName, href)
		if !boolean {
			//不再redis中就添加
			_, _ = client.AddElementToSet(queueName, href)
			//将解析到的图书详细信息URL放到消息队列
			bookDetailURL.PublishSimple(href)
		}
	})

}
