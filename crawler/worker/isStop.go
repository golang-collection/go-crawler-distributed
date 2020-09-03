package worker

import (
	"go-crawler-distributed/crawler/crawerConfig"
	"go-crawler-distributed/mq/mqTools"
)

/**
* @Author: super
* @Date: 2020-09-03 09:02
* @Description:
**/

func IsStop(contents []byte, mq *mqTools.RabbitMQ) bool{
	if string(contents) == crawerConfig.StopTAG{
		mq.PublishSimple(crawerConfig.StopTAG)
		return true
	}else{
		return false
	}
}