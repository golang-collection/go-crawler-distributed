package main

import (
	"context"
	"fmt"
	"go-crawler-distributed/initConf"
	"go-crawler-distributed/internal/crontab/worker"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-08 23:04
* @Description:
**/

func main() {
	initConf.Init("/Users/super/develop/go-crawler-distributed/configs/")
	if err := worker.NewScheduler(); err != nil{
		fmt.Println(err)
		return
	}
	if err := worker.WatchJobs(context.Background()); err != nil{
		fmt.Println(err)
		return
	}
	worker.WatchKiller(context.Background())

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}

}