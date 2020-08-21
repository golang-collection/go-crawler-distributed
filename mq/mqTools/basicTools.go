package mqTools

import (
	"fmt"
	"github.com/streadway/amqp"
	"go-crawler-distributed/service/watchConfig"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-13 08:44
* @Description:
**/

var logger = unifiedLog.GetLogger()

type RabbitMQ struct {
	conn    *amqp.Connection
	Channel *amqp.Channel

	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	Mqurl string
}

// 创建RabbitMQ实例
func NewRabbitMQ(queuqName string,
	exchange string, key string) *RabbitMQ {

	rabbitMQUrl, err := watchConfig.GetRabbitMQUrl()
	if err != nil{
		logger.Error("get rabbitmq url error", zap.Error(err))
		panic(err)
	}
	fmt.Println(rabbitMQUrl)
	rabbitmq := &RabbitMQ{
		QueueName: queuqName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     rabbitMQUrl,
	}
	//创建rabbitmq连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.FailOnErr(err, "创建连接错误")
	rabbitmq.Channel, err = rabbitmq.conn.Channel()
	rabbitmq.FailOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开channel和connection的连接释放资源
func (r *RabbitMQ) Destory() {
	r.Channel.Close()
	r.conn.Close()
}

//自定义错误处理函数
func (r *RabbitMQ) FailOnErr(err error, message string) {
	if err != nil {
		logger.Error(message, zap.Error(err))
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
