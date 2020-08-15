package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/mq/mqTools"
	"log"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-14 20:49
* @Description:
**/
func ParseTagList(contents []byte, queueName string, url string) {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatal(err)
	}

	//初始化消息队列
	bookDetailURL := mqTools.NewRabbitMQSimple(queueName)

	result := dom.Find("table[class=tagCol]").Find("a")
	href := ""
	result.Each(func(i int, selection *goquery.Selection) {
		href = url + selection.Text()
		fmt.Printf("Fetching: %s\n", href)

		//将解析到的图书详细信息URL放到消息队列
		bookDetailURL.PublishSimple(href)

	})

}
