package client

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

var redisOP proto.RedisOperationService

func init(){
	service := micro.NewService(
		micro.Name("go.micro.service.redis.client"),
	)
	service.Init()
	redisOP = proto.NewRedisOperationService("go.micro.service.redis", service.Client())
}

//func SetString(key, value string) {
//	rsp, err := redisOP.SetString(context.TODO(), &proto.Request{Key: key, Value: value})
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Println(rsp.Result)
//}
//
//func GetString(key string) {
//	rsp, err := redisOP.GetString(context.TODO(), &proto.Request{Key: key})
//	if err != nil {
//		log.Println(err)
//	}
//	fmt.Println(rsp.Result)
//}

func AddElementToSet(key string, value string) (int32 ,error) {
	res, err :=  redisOP.AddElementToSet(context.TODO(), &proto.Request{Key: key, Value:value})
	if err != nil{
		fmt.Println(err)
		return -1, err
	}
	return res.Result, err
}

func ElementIsInSet(key string, value string) (bool, error){
	rsp, err := redisOP.ElementIsInSet(context.TODO(), &proto.Request{Key: key, Value:value})
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return rsp.Result, err
}