package main

import (
	"go-crawler-distributed/internal/crawler"
	"go-crawler-distributed/internal/crawler/crawerConfig"
	"go-crawler-distributed/internal/crawler/meituan/parser"
)

/**
* @Author: super
* @Date: 2020-09-01 18:30
* @Description:
**/

func main() {
	crawler.Crawl("", crawerConfig.ArticleList, "ArticleList", parser.ParseArticleList)
}
