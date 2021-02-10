package common

import "errors"

/**
* @Author: super
* @Date: 2021-02-06 19:56
* @Description:
**/

var ERR_LOCK_ALREDAY_REQUIRED = errors.New("锁被占用")
var ERR_NO_LOCAL_IP_FOUND = errors.New("没有找到网卡IP")

const (
	// 任务保存目录
	JOB_SAVE_DIR = "/cron/jobs/"

	// 任务强杀目录
	JOB_KILLER_DIR = "/cron/killer/"

	// 任务锁目录
	JOB_LOCK_DIR = "/cron/lock/"

	// 服务注册目录
	JOB_WORKER_DIR = "/cron/workers/"

	// 保存任务事件
	JOB_EVENT_SAVE = 1

	// 删除任务事件
	JOB_EVENT_DELETE = 2

	// 强杀任务事件
	JOB_EVENT_KILL = 3
)
