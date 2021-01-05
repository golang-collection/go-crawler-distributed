package cache

import (
	"github.com/garyburd/redigo/redis"

	"go-crawler-distributed/global"
)

/**
* @Author: super
* @Date: 2020-12-29 11:20
* @Description:
**/

func SetString(key string, value string) (string, error) {
	c := global.RedisEngine.Get()
	defer c.Close()
	str, err := redis.String(c.Do("set", key, value))
	if err != nil {
		return "", err
	}
	return str, err
}

func GetString(key string) (string, error) {
	c := global.RedisEngine.Get()
	defer c.Close()
	str, err := redis.String(c.Do("get", key))
	if err != nil {
		return "", err
	}
	return str, err
}

func AddElementToSet(key string, value string) (int, error) {
	c := global.RedisEngine.Get()
	defer c.Close()
	result, err := redis.Int(c.Do("sadd", key, value))
	if err != nil {
		return -1, err
	}
	return result, err
}

func ElementIsInSet(key string, value string) (bool, error) {
	c := global.RedisEngine.Get()
	defer c.Close()
	result, err := redis.Int(c.Do("sismember", key, value))
	if err != nil {
		return false, err
	}
	if result == 1 {
		return true, err
	}
	return false, err
}

func GetAllElementFromSet(key string) ([]string, error) {
	c := global.RedisEngine.Get()
	defer c.Close()
	return redis.Strings(c.Do("smembers", key))
}

func DelAllElementFromSet(key string) (int, error) {
	c := global.RedisEngine.Get()
	defer c.Close()
	return redis.Int(c.Do("DEL", key))
}
