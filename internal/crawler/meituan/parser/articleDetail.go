package parser

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/model"
	"go-crawler-distributed/pkg/mq"
	"go-crawler-distributed/pkg/util"
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
		global.Logger.Error(context.Background(), err)
	}

	article := &model.Article{}

	result := dom.Find("a[rel=bookmark]")
	article.Url = url

	title := result.Text()
	article.Title = title

	s, err := util.ZipString(contents)
	if err != nil {
		global.Logger.Error(context.Background(), err)
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
		global.Logger.Error(context.Background(), err)
	} else {
		//将解析到的图书详细信息URL放到消息队列
		err = mq.Publish(queueName, bytes)
		if err != nil {
			global.Logger.Error(context.Background(), err)
		}
	}
}
