package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/fetcher"
)

/**
* @Author: super
* @Date: 2020-08-14 16:50
* @Description:
**/
func main() {
	contents, _ := fetcher.Fetch("https://book.douban.com/subject/6781808/")
	parser.ParseBookDetail(contents, crawerConfig.BOOK_DETAIL_URL)
}
