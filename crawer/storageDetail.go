package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/db/DBOperation"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/tools"
	"log"
)

/**
* @Author: super
* @Date: 2020-08-14 16:27
* @Description:
**/

func main() {
	bookDetailURL := mqTools.NewRabbitMQSimple(crawerConfig.BOOK_DETAIL)
	messages := bookDetailURL.GetMsgs()

	forever := make(chan bool)

	go func() {
		for d := range messages {
			go func() {
				log.Printf("Storage BookDetail: %s", d.Body)
				book, err := tools.JsonToBook(string(d.Body))
				err = DBOperation.InsertBook(book)
				if err != nil {
					log.Fatalln(err)
				}
			}()
		}
	}()

	<-forever
}
