package mq

import (
	"fmt"
	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/setting"
	"strings"
	"testing"
)

/**
* @Author: super
* @Date: 2020-12-29 13:58
* @Description:
**/

func TestPublish(t *testing.T) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-frontground-backend/configs", ",")...)
	if err != nil {
		t.Error(err)
	}
	err = newSetting.ReadSection("RabbitMQ", &global.RabbitMQSetting)
	if err != nil {
		t.Error(err)
	}
	global.RabbitMQEngine, err = NewRabbitMQEngine(global.RabbitMQSetting)
	if err != nil {
		t.Error(err)
	}
	err = Publish("test.oss", []byte("dddddddwedad"))
	if err != nil {
		t.Log(err)
	}
}

func TestConsume(t *testing.T) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-frontground-backend/configs", ",")...)
	if err != nil {
		t.Error(err)
	}
	err = newSetting.ReadSection("RabbitMQ", &global.RabbitMQSetting)
	if err != nil {
		t.Error(err)
	}
	global.RabbitMQEngine, err = NewRabbitMQEngine(global.RabbitMQSetting)
	if err != nil {
		t.Error(err)
	}
	msgs, err := Consume("test.oss")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println(d.Body)
			//实现其他的逻辑函数
		}
	}()
	<-forever
}
