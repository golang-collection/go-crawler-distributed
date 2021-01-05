package mq

import (
	"github.com/streadway/amqp"

	"go-crawler-distributed/global"
)

/**
* @Author: super
* @Date: 2020-12-29 13:47
* @Description:
**/

func Consume(queueName string) (<-chan amqp.Delivery, error) {
	//1. 申请队列，如果队列不存在则自动创建，如果存在则获取存在的队列
	//保证队列存在，使消息发送到队列中
	_, err := global.RabbitMQEngine.Channel.QueueDeclare(queueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性，独占队列
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		return nil, err
	}

	//接受消息
	msgs, err := global.RabbitMQEngine.Channel.Consume(
		queueName,
		//用于区分多个不同的消费者
		"",
		//是否自动应答，也就是消费者消费一个队列后是否主动告知rabbitmq当前的消息我已经消费完
		//rabbitmq会根据这个判断是否可以删除该消息
		//为false的话要手动实现
		false,
		//是否具有排他性
		false,
		//如果为true不能在同一个connection中发送消息传递给当前conn的消费者
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
