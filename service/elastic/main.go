package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-crawler-distributed/service/elastic/proto"
	"go-crawler-distributed/service/elastic/server"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-01 20:44
* @Description:
**/

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.elastic"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
	)
	service.Init()

	// 注册处理器
	err := proto.RegisterElasticOperationHandler(service.Server(), new(server.Elastic))
	if err != nil{
		unifiedLog.GetLogger().Error("elastic service register error", zap.Error(err))
	}

	// 运行服务
	if err := service.Run(); err != nil {
		unifiedLog.GetLogger().Error("elastic service error", zap.Error(err))
	}
}