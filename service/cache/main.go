package main

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"

	"go-crawler-distributed/configs"
	"go-crawler-distributed/global"
	"go-crawler-distributed/service/cache/proto"
	"go-crawler-distributed/service/cache/server"
)

/**
* @Author: super
* @Date: 2020-08-17 20:21
* @Description:
**/

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			configs.ConsulURL,
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
		global.Logger.Error(context.Background(), err)
	}

	// 运行服务
	if err := service.Run(); err != nil {
		global.Logger.Error(context.Background(), err)
	}
}
