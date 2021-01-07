package mq

import (
	"github.com/streadway/amqp"

	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/setting"
)

/**
* @Author: super
* @Date: 2020-11-18 11:50
* @Description: rabbitMQ连接池
**/

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

// Init : 初始化MQ连接信息
func NewRabbitMQEngine(rabbitMQSetting *setting.RabbitMQSettingS) (*global.RabbitMQ, error) {
	rabbit := &global.RabbitMQ{}
	if err := initChannel(rabbit, rabbitMQSetting); err!= nil {
		if rabbit.Channel != nil{
			rabbit.Channel.NotifyClose(notifyClose)
		}
		return nil, err
	}
	// 断线自动重连
	go func(rabbitMQ *global.RabbitMQ, rabbitMQSetting *setting.RabbitMQSettingS) {
		for {
			select {
			case _ = <-notifyClose:
				rabbit.Conn = nil
				rabbit.Channel = nil
				_ = initChannel(rabbitMQ, rabbitMQSetting)
			}
		}
	}(rabbit, rabbitMQSetting)
	return rabbit, nil
}

//初始化channel
func initChannel(rabbitMQ *global.RabbitMQ, rabbitMQSetting *setting.RabbitMQSettingS) error {
	if rabbitMQ.Channel != nil {
		return nil
	}
	var err error
	rabbitHost := "amqp://" + rabbitMQSetting.UserName + ":" + rabbitMQSetting.Password + "@" + rabbitMQSetting.Host + "/"
	rabbitMQ.Conn, err = amqp.Dial(rabbitHost)
	if err != nil {
		return err
	}

	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	if err != nil {
		return err
	}

	return nil
}
