package main

import (
	"go-crawler-distributed/mq/mqTools"
	"log"
)

/**
* @Author: super
* @Date: 2020-08-13 09:00
* @Description:
**/
func main() {
	testOne := mqTools.NewRabbitMQRouting("testRouting", "testOne")
	messages := testOne.BindConsumerRouting()

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message : %s", d.Body)
		}
	}()

	<-forever
}
