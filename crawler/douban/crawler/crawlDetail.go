package crawler

import (
	"go-crawler-distributed/crawler"
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/douban/parser"
)

/**
* @Author: super
* @Date: 2020-08-14 15:50
* @Description:
**/
func CrawlDetail() {
	crawler.Crawl(crawerConfig.BookDetailUrl, crawerConfig.BookDetail, "BookDetail", parser.ParseBookDetail)
}
