package main

import (
	"github.com/micro/go-micro/v2"
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
	service := micro.NewService(
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