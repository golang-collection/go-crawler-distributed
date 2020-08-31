package crawler

import (
	"go-crawler-distributed/crawler"
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/douban/parser"
)

/**
* @Author: super
* @Date: 2020-08-12 19:47
* @Description:
**/

func CrawlList() {
	crawler.Crawl(crawerConfig.TagUrl, crawerConfig.BookDetailUrl, "tagList", parser.ParseBookList)
}
