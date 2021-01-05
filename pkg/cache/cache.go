package cache

import (
	"github.com/garyburd/redigo/redis"

	"go-crawler-distributed/pkg/setting"

	"time"
)

/**
* @Author: super
* @Date: 2020-11-18 11:35
* @Description: 根据配置创建redis连接池
**/

func NewRedisEngine(cacheSetting *setting.CacheSettingS) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     cacheSetting.MaxIdle,
		MaxActive:   cacheSetting.MaxActive,
		IdleTimeout: 300 * time.Second,
		// 如果空闲列表中没有可用的连接,且当前Active连接数 < MaxActive, 则等待
		Wait: true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", cacheSetting.Host)
			if err != nil {
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
	}, nil
}
