package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/cache/cacheOperation"
	"go-crawler-distributed/mq/mqTools"
	"log"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-14 13:54
* @Description:
**/

func ParseBookList(contents []byte, queueName string) {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatal(err)
	}

	//初始化消息队列
	bookDetailURL := mqTools.NewRabbitMQSimple(queueName)

	result := dom.Find("a[title]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		fmt.Printf("Fetching: %s\n", href)

		//redis去重
		boolean, _ := cacheOperation.ElementIsInSet(queueName, href)
		if !boolean {
			//不再redis中就添加
			_, _ = cacheOperation.AddElementToSet(queueName, href)
			//将解析到的图书详细信息URL放到消息队列
			bookDetailURL.PublishSimple(href)
		}
	})

}
