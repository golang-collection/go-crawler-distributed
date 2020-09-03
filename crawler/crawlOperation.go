package crawler

import (
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/crawler/worker"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"sync"
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
	} else if targetMQ == ""{
		//存储模块
		getMessage(sourceMQ, funcParser, true)
	}else{
		getMessage(sourceMQ, funcParser, false)
	}
}

func getMessage(sourceMQ string, funcParser *worker.FuncParser, isStorage bool){
	mq := mqTools.NewRabbitMQSimple(sourceMQ)
	messages := mq.GetMsgs()
	unifiedLog.GetLogger().Info("Ready to fetching", zap.String("parser name", funcParser.Name))
	var wg sync.WaitGroup
	for d := range messages {
		if string(d.Body) == crawerConfig.StopTAG{
			break
		}else{
			wg.Add(1)
			go func(data []byte) {
				defer wg.Done()
				//是否是保存操作
				if isStorage{
					doStorage(data, funcParser)
				}else{
					d := string(data)
					doCrawler(d, funcParser)
				}
			}(d.Body)
		}
		time.Sleep(time.Second * 2)
	}
	wg.Wait()
	mq.Destory()
	unifiedLog.GetLogger().Info("finish fetching", zap.String("parser name", funcParser.Name))
}

func doCrawler(url string, funcParser *worker.FuncParser){
	unifiedLog.GetLogger().Info("fetching", zap.String(funcParser.Name, url))

	r := worker.Request{
		Url:    url,
		Parser: funcParser,
	}

	worker.Worker(r)
}

func doStorage(data []byte, funcParser *worker.FuncParser){
	unifiedLog.GetLogger().Info("saving", zap.String(funcParser.Name, string(data)))
	funcParser.Parse(data, "")
}