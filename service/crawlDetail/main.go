package main

import (
	"go-crawler-distributed/crawler"
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/douban/parser"
)

/**
* @Author: super
* @Date: 2020-08-31 16:52
* @Description:
**/

func main() {
	crawler.Crawl(crawerConfig.BookDetailUrl, crawerConfig.BookDetail, "BookDetail", parser.ParseBookDetail)
}