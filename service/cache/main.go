package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
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
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.redis"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()

	// 注册处理器
	err := proto.RegisterRedisOperationHandler(service.Server(), new(server.CacheStruct))
	if err != nil {
		logger.Error("cache service register error", zap.Error(err))
	}

	// 运行服务
	if err := service.Run(); err != nil {
		logger.Error("cache service error", zap.Error(err))
	}
}
