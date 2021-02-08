package worker

import (
	"fmt"
	"go-crawler-distributed/internal/crontab/common"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-08 19:47
* @Description:
**/

type Scheduler struct {
	JobEventChan chan *common.JobEvent
	JobPlanTable map[string]*common.JobSchedulePlan //任务调度计划表
}

var (
	GlobalScheduler *Scheduler
)

//处理任务事件
func (s *Scheduler) handleJobEvent(jobEvent *common.JobEvent) {
	var (
		jobSchedulePlan *common.JobSchedulePlan
		jobExisted      bool
		err             error
	)
	switch jobEvent.EventType {
	//保存任务事件
	case common.JOB_EVENT_SAVE:
		if jobSchedulePlan, err = common.BuildJobSchedulePlan(jobEvent.Job); err != nil {
			return
		}
		s.JobPlanTable[jobEvent.Job.Name] = jobSchedulePlan
	//删除任务事件
	case common.JOB_EVENT_DELETE:
		if jobSchedulePlan, jobExisted = s.JobPlanTable[jobEvent.Job.Name]; jobExisted {
			delete(s.JobPlanTable, jobEvent.Job.Name)
		}
	}
}

func (s *Scheduler) TrySchedule() (scheduleAfter time.Duration) {
	var(
		jobPlan *common.JobSchedulePlan
		now time.Time
		nearTime *time.Time
	)

	if len(s.JobPlanTable) == 0{
		scheduleAfter = 1 * time.Second
		return
	}

	now = time.Now()
	for _, jobPlan = range s.JobPlanTable{
		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now){
			//todo:执行任务
			fmt.Println("执行任务", jobPlan.Job.Name)
			jobPlan.NextTime = jobPlan.Schedule.Next(now)
		}

		if nearTime == nil || jobPlan.NextTime.Before(*nearTime){
			nearTime = &jobPlan.NextTime
		}
	}
	scheduleAfter = (*nearTime).Sub(now)
	return
}

func (s *Scheduler) schedulerLoop() {
	var (
		jobEvent *common.JobEvent
		scheduleAfter time.Duration
		scheduleTimer *time.Timer
	)

	scheduleAfter = s.TrySchedule()
	scheduleTimer = time.NewTimer(scheduleAfter)

	for {
		select {
		//监听任务变化
		case jobEvent = <-s.JobEventChan:
			//对内存中的任务进行增删改查
			s.handleJobEvent(jobEvent)
		case <-scheduleTimer.C:

		}
		scheduleAfter = s.TrySchedule()
		scheduleTimer.Reset(scheduleAfter)
	}
}

func (s *Scheduler) PushJobEvent(jobEvent *common.JobEvent) {
	s.JobEventChan <- jobEvent
}

func NewScheduler() (err error) {
	GlobalScheduler = &Scheduler{
		JobEventChan: make(chan *common.JobEvent, 10000),
		JobPlanTable: make(map[string]*common.JobSchedulePlan),
	}
	go GlobalScheduler.schedulerLoop()
	return
}
