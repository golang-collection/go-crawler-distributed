package common

/**
* @Author: super
* @Date: 2021-02-07 18:39
* @Description:
**/

// 任务执行日志
type JobLog struct {
	JobName      string `json:"jobName" bson:"jobName"`           // 任务名字
	Command      string `json:"command" bson:"command"`           // 脚本命令
	Err          string `json:"err" bson:"err"`                   // 错误原因
	Output       string `json:"output" bson:"output"`             // 脚本输出
	PlanTime     int64  `json:"planTime" bson:"planTime"`         // 计划开始时间
	ScheduleTime int64  `json:"scheduleTime" bson:"scheduleTime"` // 实际调度时间
	StartTime    int64  `json:"startTime" bson:"startTime"`       // 任务执行开始时间
	EndTime      int64  `json:"endTime" bson:"endTime"`           // 任务执行结束时间
}

// 日志批次,防止每条日志都单次插入数据库中
type LogBatch struct {
	Logs []interface{} // 多条日志
}

// 任务日志过滤条件
type JobLogFilter struct {
	JobName string `bson:"jobName"`
}

// 任务日志排序规则
type SortLogByStartTime struct {
	SortOrder int `bson:"startTime"` // {startTime: -1}
}
