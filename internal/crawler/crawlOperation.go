package crawler

import (
	"context"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crawler/crawerConfig"
	"go-crawler-distributed/internal/crawler/worker"
	"go-crawler-distributed/pkg/mq"
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
func Crawl(sourceMQ string, targetMQ string, name string, function worker.ParserFunc) {
	funcParser := worker.NewFuncParser(function, targetMQ, name)
	if sourceMQ == "" {
		//代表开始模块
		url := crawerConfig.StartUrl
		doCrawler(url, funcParser)
	} else if targetMQ == "" {
		//存储模块
		getMessage(sourceMQ, funcParser, true)
	} else {
		getMessage(sourceMQ, funcParser, false)
	}
}

func getMessage(sourceMQ string, funcParser *worker.FuncParser, isStorage bool) {
	messages, err := mq.Consume(sourceMQ)
	if err != nil {
		global.Logger.Error(context.Background(), err)
		return
	}
	global.Logger.Infof(context.Background(), "parser name: %s", funcParser.Name)

	var wg sync.WaitGroup
	for d := range messages {
		d.Ack(false)
		if string(d.Body) == crawerConfig.StopTAG {
			break
		} else {
			wg.Add(1)
			go func(data []byte) {
				defer wg.Done()
				//是否是保存操作
				if isStorage {
					doStorage(data, funcParser)
				} else {
					d := string(data)
					doCrawler(d, funcParser)
				}
			}(d.Body)
		}
		time.Sleep(time.Second * 2)
	}
	wg.Wait()
	global.Logger.Infof(context.Background(), "finish fetching parser name: %s", funcParser.Name)
}

func doCrawler(url string, funcParser *worker.FuncParser) {
	global.Logger.Infof(context.Background(), "fetching: %s", url)
	r := worker.Request{
		Url:    url,
		Parser: funcParser,
	}

	worker.Worker(r)
}

func doStorage(data []byte, funcParser *worker.FuncParser) {
	global.Logger.Infof(context.Background(), "saving: %s", data)
	funcParser.Parse(data, "")
}
