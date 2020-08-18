package server

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"go-crawler-distributed/cache/cacheTools"
	"go-crawler-distributed/service/cache/proto"
)

/**
* @Author: super
* @Date: 2020-08-17 20:08
* @Description: redis微服务化
**/

type CacheStruct struct {
}

func (cache *CacheStruct) SetString(ctx context.Context, req *proto.Request, res *proto.StringResponse) error {
	c := cacheTools.GetConn()

	str, err := redis.String(c.Do("set", req.Key, req.Value))
	if err != nil {
		return err
	}
	res.Result = str
	return nil
}

func (cache *CacheStruct) GetString(ctx context.Context, req *proto.Request, res *proto.StringResponse) error {
	c := cacheTools.GetConn()

	str, err := redis.String(c.Do("get", req.Key))
	if err != nil {
		return err
	}
	res.Result = str
	return nil
}

func (cache *CacheStruct) AddElementToSet(ctx context.Context, req *proto.Request, res *proto.IntResponse) error {
	c := cacheTools.GetConn()
	defer c.Close()

	result, err := redis.Int(c.Do("sadd", req.Key, req.Value))
	if err != nil {
		return err
	}
	res.Result = int32(result)
	return nil
}

func (cache *CacheStruct) ElementIsInSet(ctx context.Context, req *proto.Request, res *proto.BoolResponse) error {
	c := cacheTools.GetConn()
	defer c.Close()

	result, err := redis.Int(c.Do("sismember", req.Key, req.Value))
	if err != nil {
		return err
	}
	if result == 1{
		res.Result = true
	}else{
		res.Result = false
	}
	return nil
}

func (cache *CacheStruct) GetAllElementFromSet(ctx context.Context, req *proto.Request, res *proto.StringsResponse) error {
	c := cacheTools.GetConn()
	strs, err := redis.Strings(c.Do("smembers", req.Key))
	if err != nil {
		return err
	}
	res.Result = strs
	return nil
}
