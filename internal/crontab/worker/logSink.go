package worker

import (
	"context"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crontab/common"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-09 14:06
* @Description:
**/

type LogSink struct {
	LogCollection  *mongo.Collection
	LogChan        chan *common.JobLog
	AutoCommitChan chan *common.LogBatch
}

var (
	GlobalLogSink *LogSink
)

func (l *LogSink) SaveLogs(batch *common.LogBatch) {
	_, err := l.LogCollection.InsertMany(context.TODO(), batch.Logs)
	if err != nil {
		log.Println("saveLogs", err)
	}
	log.Println("saveLogs")
}

func (l *LogSink) writeLoop() {
	var (
		jobLog       *common.JobLog
		logBatch     *common.LogBatch // 当前的批次
		commitTimer  *time.Timer
		timeoutBatch *common.LogBatch // 超时批次
	)
	for {
		select {
		case jobLog = <-l.LogChan:
			if logBatch == nil {
				logBatch = &common.LogBatch{}
				// 让这个批次超时自动提交(给1秒的时间）
				commitTimer = time.AfterFunc(
					time.Duration(1000)*time.Millisecond,
					func(batch *common.LogBatch) func() {
						return func() {
							l.AutoCommitChan <- batch
						}
					}(logBatch),
				)
			}

			// 把新日志追加到批次中
			logBatch.Logs = append(logBatch.Logs, jobLog)

			// 如果批次满了, 就立即发送
			if len(logBatch.Logs) >= 100 {
				// 发送日志
				l.SaveLogs(logBatch)
				// 清空logBatch
				logBatch = nil
				// 取消定时器
				commitTimer.Stop()
			}
		case timeoutBatch = <-l.AutoCommitChan: // 过期的批次
			// 判断过期批次是否仍旧是当前的批次
			if timeoutBatch != logBatch {
				continue // 跳过已经被提交的批次
			}
			// 把批次写入到mongo中
			l.SaveLogs(timeoutBatch)
			// 清空logBatch
			logBatch = nil
		}
	}
}

// 发送日志
func (l *LogSink) Append(jobLog *common.JobLog) {
	select {
	case l.LogChan <- jobLog:
	default:
		// 队列满了就丢弃
	}
}

func NewLogSink() (err error) {
	GlobalLogSink = &LogSink{
		LogCollection:  global.MongoDBEngine.Database("cron").Collection("log"),
		LogChan:        make(chan *common.JobLog, 1000),
		AutoCommitChan: make(chan *common.LogBatch, 1000),
	}
	go GlobalLogSink.writeLoop()
	return
}
