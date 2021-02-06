package main

import (
	"go-crawler-distributed/internal/crawler"
	"go-crawler-distributed/internal/crawler/crawerConfig"
	"go-crawler-distributed/internal/crawler/douban/parser"
)

/**
* @Author: super
* @Date: 2020-08-31 16:51
* @Description:
**/

func main() {
	crawler.Crawl(crawerConfig.TagUrl, crawerConfig.BookDetailUrl, "tagList", parser.ParseBookList)
}
