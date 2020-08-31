package crawler

import (
	"go-crawler-distributed/crawler"
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/douban/parser"
)

/**
* @Author: super
* @Date: 2020-08-14 20:46
* @Description:
**/
func CrawlTags() {
	crawler.Crawl("", crawerConfig.TagUrl, "tags", parser.ParseTagList)
}
