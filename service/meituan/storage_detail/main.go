package main

import (
	"go-crawler-distributed/crawler"
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/meituan/storage"
)

/**
* @Author: super
* @Date: 2020-09-01 19:37
* @Description:
**/

func main() {
	crawler.Crawl(crawerConfig.ArticleDetail, "", "storageArticleDetail", storage.StorageArticle)
}