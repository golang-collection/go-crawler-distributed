package main

import (
	"go-crawler-distributed/internal/crawler"
	"go-crawler-distributed/internal/crawler/crawerConfig"
	"go-crawler-distributed/internal/crawler/douban/storage"
)

/**
* @Author: super
* @Date: 2020-08-31 16:52
* @Description:
**/

func main() {
	crawler.Crawl(crawerConfig.BookDetail, "", "storageBookDetail", storage.ParseAndStorage)
}
