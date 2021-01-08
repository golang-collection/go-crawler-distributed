package client

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-crawler-distributed/global"
	_ "go-crawler-distributed/init"
	"go-crawler-distributed/service/cache/proto"
)

/**
* @Author: super
* @Date: 2021-01-05 19:33
* @Description:
**/

var redisOP proto.RedisOperationService

func init() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			global.ConsulSetting.Url,
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.redis.client"),
	)
	service.Init()
	redisOP = proto.NewRedisOperationService("go.micro.service.redis", service.Client())
}

func AddElementToSet(key string, value string) (int32, error) {
	res, err := redisOP.AddElementToSet(context.TODO(), &proto.Request{Key: key, Value: value})
	if err != nil {
		global.Logger.Error(context.Background(), err)
		return -1, err
	}
	return res.Result, nil
}

func ElementIsInSet(key string, value string) (bool, error) {
	rsp, err := redisOP.ElementIsInSet(context.TODO(), &proto.Request{Key: key, Value: value})
	if err != nil {
		global.Logger.Error(context.Background(), err)
		return false, err
	}
	return rsp.Result, err
}
