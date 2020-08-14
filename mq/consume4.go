package main

import (
	"go-crawler-distributed/mq/mqTools"
	"log"
)

/**
* @Author: super
* @Date: 2020-08-14 13:43
* @Description:
**/

func main() {
	testOne := mqTools.NewRabbitMQSimple("testTwo")
	messages := testOne.GetMsgs()
	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message : %s", d.Body)
		}
	}()

	<-forever
}
