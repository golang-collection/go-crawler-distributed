package mq

import (
	"github.com/streadway/amqp"

	"go-crawler-distributed/global"
)

/**
* @Author: super
* @Date: 2020-12-29 13:43
* @Description:
**/

//发布消息
func Publish(queueName string, msg []byte) error {
	//检查channel是否正常
	if err := initChannel(global.RabbitMQEngine, global.RabbitMQSetting); err != nil {
		return err
	}

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
		return err
	}

	//2. 发送消息到队列中
	err = global.RabbitMQEngine.Channel.Publish(
		"",
		queueName,
		// 如果为true， 则根据exchange类型和routkey规则，如果无法找到符合条件的队列
		// 那么会把发送的消息回退给publish
		false,
		//如果为true，当exchange发送消息到队列后发现没有consume，则会把发送的消息返回给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
