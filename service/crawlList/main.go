package main

import (
	"go-crawler-distributed/crawler"
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/douban/parser"
)

/**
* @Author: super
* @Date: 2020-08-31 16:51
* @Description:
**/

func main() {
	crawler.Crawl(crawerConfig.TagUrl, crawerConfig.BookDetailUrl, "tagList", parser.ParseBookList)
}