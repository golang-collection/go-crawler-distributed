package crawler

import (
	"go-crawler-distributed/crawler/crawerConfig"
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

//sourceMQ: 配置从哪里读取消息
//targetMQ: 配置将解析好的消息发送到什么位置
//name: 当前工作节点的名称
//function: 页面的具体解析函数
func Crawl(sourceMQ string, targetMQ string, name string, function worker.ParserFunc){
	funcParser := worker.NewFuncParser(function, targetMQ, name)
	if sourceMQ == ""{
		//代表开始模块
		url := crawerConfig.StartUrl
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