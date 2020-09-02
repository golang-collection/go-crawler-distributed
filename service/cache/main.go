package main

import (
	"github.com/micro/go-micro/v2"
	"go-crawler-distributed/service/cache/proto"
	"go-crawler-distributed/service/cache/server"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-17 20:21
* @Description:
**/

var logger = unifiedLog.GetLogger()

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.redis"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()

	// 注册处理器
	err := proto.RegisterRedisOperationHandler(service.Server(), new(server.CacheStruct))
	if err != nil{
		logger.Error("cache service register error", zap.Error(err))
	}

	// 运行服务
	if err := service.Run(); err != nil {
		logger.Error("cache service error", zap.Error(err))
	}
}