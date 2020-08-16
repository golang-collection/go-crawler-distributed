package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/storage"
	"go-crawler-distributed/crawer/persistence"
	"go-crawler-distributed/mq/mqTools"
	"log"
)

/**
* @Author: super
* @Date: 2020-08-14 16:27
* @Description:
**/

func main() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BookDetail)
	messages := bookDetailURL.GetMsgs()

	forever := make(chan bool)

	funcStorage := persistence.FuncStorage{
		Name:      "BookDetail",
		ParseFunc: storage.ParseAndStorage,
	}

	go func() {
		log.Println("Ready to storage " + funcStorage.Name)
		for d := range messages {
			go func() {
				log.Printf("Storage "+funcStorage.Name+": %s", d.Body)

				err := funcStorage.ParseFunc(d.Body)
				if err != nil {
					log.Fatalln(err)
				}
			}()
		}
	}()

	<-forever
}
