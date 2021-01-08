package main

import (
	_ "go-crawler-distributed/init"
	"go-crawler-distributed/internal/crawler"
	"go-crawler-distributed/internal/crawler/crawerConfig"
	"go-crawler-distributed/internal/crawler/meituan/parser"
)

/**
* @Author: super
* @Date: 2020-09-01 19:03
* @Description:
**/

func main() {
	crawler.Crawl(crawerConfig.ArticleList, crawerConfig.ArticleUrlList, "ArticleUrlList", parser.ParseArticleUrlList)
}
