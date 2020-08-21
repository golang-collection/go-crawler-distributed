package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/worker"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-14 20:46
* @Description:
**/
func main() {
	url := "https://book.douban.com/tag/"

	funcParser := worker.NewFuncParser(parser.ParseTagList, crawerConfig.TagUrl, "tags")


    logger.Info("fetching", zap.String(funcParser.Name, url))

	r := worker.Request{
		Url:    url,
		Parser: funcParser,
	}
	worker.Worker(r)
}
