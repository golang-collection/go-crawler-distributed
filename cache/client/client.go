package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"go-crawler-distributed/service/cache/proto"
)

/**
* @Author: super
* @Date: 2020-08-17 20:26
* @Description:
**/

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.redis.client"),
	)
	service.Init()
	redisOP := proto.NewRedisOperationService("go.micro.service.redis", service.Client())

	SetString(redisOP)
	GetString(redisOP)

}

func SetString(redisOP proto.RedisOperationService) {
	rsp, err := redisOP.SetString(context.TODO(), &proto.Request{Key: "s1", Value: "lalalala"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Result)
}

func GetString(redisOP proto.RedisOperationService) {
	rsp, err := redisOP.GetString(context.TODO(), &proto.Request{Key: "s1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Result)
}
