package util

import (
	"go-crawler-distributed/internal/model"
	"go-crawler-distributed/service/elastic/proto"
)

/**
* @Author: super
* @Date: 2021-01-05 19:27
* @Description:
**/

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
