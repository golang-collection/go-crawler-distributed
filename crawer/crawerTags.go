package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/fetcher"
)

/**
* @Author: super
* @Date: 2020-08-14 20:46
* @Description:
**/
func main() {
	//交叉编译
	//CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
	url := "https://book.douban.com/tag/"
	contents, _ := fetcher.Fetch(url)
	parser.ParseTagList(contents, crawerConfig.TagUrl, url)
}
