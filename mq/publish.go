package main

import (
	"go-crawler-distributed/mq/mqTools"
	"log"
	"strconv"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-13 08:59
* @Description:
**/

func main() {
	two := mqTools.NewRabbitMQSimple("testTwo")

	for i := 0; i < 100; i++ {
		two.PublishSimple("1-Routing模式testTwo第" + strconv.Itoa(i) + "条数据")
		log.Println("Routing模式生产第" + strconv.Itoa(i) + "条数据")
		time.Sleep(1 * time.Second)
	}
}
