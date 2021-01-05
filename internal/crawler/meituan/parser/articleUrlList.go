package parser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/mq"
	"strings"
)

/**
* @Author: super
* @Date: 2020-09-01 18:59
* @Description:
**/

func ParseArticleUrlList(contents []byte, queueName string, _ string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}

	result := dom.Find("a[rel=bookmark]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, exist := selection.Attr("href")
		if exist {
			global.Logger.Infof(context.Background(), "url: %s", href)
			//将解析到的图书详细信息URL放到消息队列
			err = mq.Publish(queueName, []byte(href))
			if err != nil {
				global.Logger.Error(context.Background(), err)
			}
		}
	})
}
