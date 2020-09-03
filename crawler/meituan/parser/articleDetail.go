package parser

import (
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/model"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/tools"
	"go.uber.org/zap"
	"strings"
)

/**
* @Author: super
* @Date: 2020-09-01 19:09
* @Description:
**/

func ParseArticleDetail(contents []byte, queueName string, url string) {
	//初始化消息队列
	articleDetail := mqTools.NewRabbitMQSimple(queueName)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		logger.Error("new doc reader error", zap.Error(err))
	}


	article := &model.Article{}

	result := dom.Find("a[rel=bookmark]")
	article.Url = url

	title := result.Text()
	article.Title = title

	s, err := tools.ZipString(contents)
	if err != nil{
		logger.Error("zipString error", zap.Error(err))
	}
	article.Content = s

	result = dom.Find("a[rel=tag]")
	result.Each(func(i int, selection *goquery.Selection) {
		tag := selection.Text()
		article.Genres = append(article.Genres, tag)
	})

	//Article结构体转json
	bytes, err := article.MarshalJSON()
	if err != nil {
		logger.Error("article to json error", zap.Error(err))
	} else {
		articleJson := string(bytes)
		//将解析到的图书详细信息URL放到消息队列
		articleDetail.PublishSimple(articleJson)
	}
}