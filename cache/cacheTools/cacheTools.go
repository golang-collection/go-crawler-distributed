package cacheTools

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go-crawler-distributed/service/watchConfig"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-13 08:40
* @Description:
**/

var pool *redis.Pool

//创建redis连接池
func newRedisPool() *redis.Pool {
	redisUrl, err := watchConfig.GetRedisUrl()
	if err != nil{
		panic(err)
	}

	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   100,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", redisUrl)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func GetConn() redis.Conn {
	return pool.Get()
}
