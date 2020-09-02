package tools

import (
	"go-crawler-distributed/model"
	"go-crawler-distributed/service/elastic/proto"
)

/**
* @Author: super
* @Date: 2020-09-02 08:12
* @Description:
**/

//TODO 使用反射操作
func ProtoToArticle(article *proto.Article) *model.Article {
	result := &model.Article{}
	result.Title = article.Title
	result.Genres = article.Genres
	result.Url = article.Url
	result.Content = article.Content
	return result
}

func ArticleToProto(article *model.Article) *proto.Article {
	result := &proto.Article{}
	result.Title = article.Title
	result.Genres = article.Genres
	result.Url = article.Url
	result.Content = article.Content
	return result
}