package main

import (
	"fmt"
	"go-crawler-distributed/config"
	"go-crawler-distributed/mylog"
	"go-crawler-distributed/rpcsupport"
	"go-crawler-distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go func() {
		err := rpcsupport.ServeRpc(
			host, worker.CrawlService{})
		if err != nil {
			t.Errorf("%v", err)
		}
	}()
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		mylog.LogError("persist.client_test", err)
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "安静的雪",
		},
	}

	var result worker.ParseResult
	err = client.Call(
		config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Errorf("%v", err)
	} else {
		fmt.Println(result)
	}

	// TODO: Verify results
}
