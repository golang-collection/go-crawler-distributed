package global

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"
)

/**
* @Author: super
* @Date: 2020-09-18 08:51
* @Description: 全局配置DB
**/

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

var (
	DBEngine       *gorm.DB
	RedisEngine    *redis.Pool
	RabbitMQEngine *RabbitMQ
	ElasticEngine  *elastic.Client
)
