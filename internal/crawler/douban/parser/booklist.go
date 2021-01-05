package parser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/mq"
	"go-crawler-distributed/service/cache/client"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-14 13:54
* @Description:
**/

func ParseBookList(contents []byte, queueName string, url string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}

	result := dom.Find("a[title]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		global.Logger.Infof(context.Background(), "url: %s", href)

		//redis去重
		boolean, _ := client.ElementIsInSet(queueName, href)
		if !boolean {
			//不再redis中就添加
			_, _ = client.AddElementToSet(queueName, href)
			//将解析到的图书详细信息URL放到消息队列
			err = mq.Publish(queueName, []byte(href))
			if err != nil {
				global.Logger.Error(context.Background(), err)
			}
		}
	})
}
