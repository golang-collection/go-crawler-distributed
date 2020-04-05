package main

import (
	"go-crawler-distributed/engine"
	"go-crawler-distributed/model"
	"go-crawler-distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	//开启Rpc服务
	go serverRpc(host, "test1")
	time.Sleep(time.Second)

	//创建客户端连接
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	//调用存储服务
	item := engine.Item{
		Url:  "http://localhost:8080/mock/album.zhenai.com/u/6223663451521583212",
		Type: "zhenai",
		Id:   "6223663451521583212",
		Payload: model.Profile{

			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		t.Errorf("result: %s; err %s", result, err)
	}
}
