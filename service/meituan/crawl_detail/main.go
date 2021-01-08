package main

import (
	_ "go-crawler-distributed/init"
	"go-crawler-distributed/internal/crawler"
	"go-crawler-distributed/internal/crawler/crawerConfig"
	"go-crawler-distributed/internal/crawler/meituan/parser"
)

/**
* @Author: super
* @Date: 2020-09-01 19:10
* @Description:
**/

func main() {
	crawler.Crawl(crawerConfig.ArticleUrlList, crawerConfig.ArticleDetail, "ArticleDetail", parser.ParseArticleDetail)
}
