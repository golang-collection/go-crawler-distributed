package main

import (
	"github.com/micro/go-micro/v2"
	"go-crawler-distributed/service/cache/proto"
	"go-crawler-distributed/service/cache/server"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-17 20:21
* @Description:
**/

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.redis"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()

	// 注册处理器
	proto.RegisterRedisOperationHandler(service.Server(), new(server.CacheStruct))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}