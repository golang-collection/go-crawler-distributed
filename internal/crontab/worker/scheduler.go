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
	JobEventChan      chan *common.JobEvent
	JobPlanTable      map[string]*common.JobSchedulePlan //任务调度计划表
	JobExecutingTable map[string]*common.JobExecuteInfo
	JobResultChan     chan *common.JobExecuteResult // 任务结果队列
}

var (
	GlobalScheduler *Scheduler
)

//处理任务事件
func (s *Scheduler) handleJobEvent(jobEvent *common.JobEvent) {
	var (
		jobSchedulePlan *common.JobSchedulePlan
		jobExcuteInfo *common.JobExecuteInfo
		jobExcuting bool
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
	case common.JOB_EVENT_KILL:
		//通过context取消任务
		if jobExcuteInfo, jobExcuting = s.JobExecutingTable[jobEvent.Job.Name]; jobExcuting{
			jobExcuteInfo.CancelFunc()
		}
	}
}

// 任务虽然被调度了，但是可能因为一些原因执行很久，加入1s执行一次的任务，单次任务执行了1分钟
// 当前任务就会被调度60次却只执行1次
func (s *Scheduler) TryStartJob(jobPlan *common.JobSchedulePlan) {
	var (
		jobExcuteInfo *common.JobExecuteInfo
		jobExcuting   bool
	)
	if jobExcuteInfo, jobExcuting = s.JobExecutingTable[jobPlan.Job.Name]; jobExcuting {
		return
	}
	jobExcuteInfo = common.BuildJobExecuteInfo(jobPlan)
	s.JobExecutingTable[jobPlan.Job.Name] = jobExcuteInfo
	fmt.Println("执行任务", jobExcuteInfo.Job.Name, jobExcuteInfo.PlanTime, jobExcuteInfo.RealTime)
	GlobalExecutor.ExecuteJob(jobExcuteInfo)
}

func (s *Scheduler) TrySchedule() (scheduleAfter time.Duration) {
	var (
		jobPlan  *common.JobSchedulePlan
		now      time.Time
		nearTime *time.Time
	)

	if len(s.JobPlanTable) == 0 {
		scheduleAfter = 1 * time.Second
		return
	}

	now = time.Now()
	for _, jobPlan = range s.JobPlanTable {
		if jobPlan.NextTime.Before(now) || jobPlan.NextTime.Equal(now) {
			s.TryStartJob(jobPlan)
			jobPlan.NextTime = jobPlan.Schedule.Next(now)
		}

		if nearTime == nil || jobPlan.NextTime.Before(*nearTime) {
			nearTime = &jobPlan.NextTime
		}
	}
	scheduleAfter = (*nearTime).Sub(now)
	return
}

func (s *Scheduler) handleJobResult(result *common.JobExecuteResult) {
	delete(s.JobExecutingTable, result.ExecuteInfo.Job.Name)
	fmt.Println("执行任务", result.ExecuteInfo.Job.Name, string(result.Output), result.Err)
}

func (s *Scheduler) schedulerLoop() {
	var (
		jobEvent      *common.JobEvent
		scheduleAfter time.Duration
		scheduleTimer *time.Timer
		jobResult     *common.JobExecuteResult
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
		case jobResult = <-s.JobResultChan: //监听任务执行结果
			s.handleJobResult(jobResult)
		}
		scheduleAfter = s.TrySchedule()
		scheduleTimer.Reset(scheduleAfter)
	}
}

func (s *Scheduler) PushJobEvent(jobEvent *common.JobEvent) {
	s.JobEventChan <- jobEvent
}

func (s *Scheduler) PushJobResult(jobResult *common.JobExecuteResult) {
	s.JobResultChan <- jobResult
}

func NewScheduler() (err error) {
	GlobalScheduler = &Scheduler{
		JobEventChan:      make(chan *common.JobEvent, 10000),
		JobPlanTable:      make(map[string]*common.JobSchedulePlan),
		JobExecutingTable: make(map[string]*common.JobExecuteInfo),
		JobResultChan:     make(chan *common.JobExecuteResult, 1000),
	}
	go GlobalScheduler.schedulerLoop()
	return
}
