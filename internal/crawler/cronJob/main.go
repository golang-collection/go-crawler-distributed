package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

/**
* @Author: super
* @Date: 2021-01-19 13:52
* @Description: 定时任务
**/

func main() {
	c := cron.New()
	//c.AddFunc("1 * * * *", func() { fmt.Println("Every hour on the half hour") })
	//c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	//c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	//c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.AddFunc("@every 2s", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
