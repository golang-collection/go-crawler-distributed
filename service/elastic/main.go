package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-crawler-distributed/global"
	"go-crawler-distributed/service/elastic/proto"
	"go-crawler-distributed/service/elastic/server"
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
			global.ConsulSetting.Url,
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
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}

	// 运行服务
	if err := service.Run(); err != nil {
		global.Logger.Error(context.Background(), err)
	}
}
