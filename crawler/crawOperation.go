package crawler

import (
	"go-crawler-distributed/crawler/worker"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-31 15:20
* @Description:
**/

func Crawl(sourceMQ string, targetMQ string, name string, function worker.ParserFunc){
	funcParser := worker.NewFuncParser(function, targetMQ, name)
	if sourceMQ == ""{
		//代表开始模块
		url := "https://book.douban.com/tag/"
		doCrawler(url, funcParser)
	}else{
		mq := mqTools.NewRabbitMQSimple(sourceMQ)
		messages := mq.GetMsgs()
		unifiedLog.GetLogger().Info("Ready to fetching", zap.String("parser name", funcParser.Name))
		for d := range messages {
			go func(data []byte) {
				url := string(data)
				doCrawler(url, funcParser)
			}(d.Body)
			time.Sleep(time.Second * 2)
		}
	}
}

func doCrawler(url string, funcParser *worker.FuncParser){
	unifiedLog.GetLogger().Info("fetching", zap.String(funcParser.Name, url))

	r := worker.Request{
		Url:    url,
		Parser: funcParser,
	}

	worker.Worker(r)
}