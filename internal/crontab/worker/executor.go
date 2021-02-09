package worker

import (
	"go-crawler-distributed/internal/crontab/common"
	"os/exec"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-09 10:45
* @Description:
**/

type Executor struct {
}

var (
	GlobalExecutor *Executor
)

func (e *Executor) ExecuteJob(info *common.JobExecuteInfo) {
	go func() {
		var (
			cmd       *exec.Cmd
			err       error
			output    []byte
			result    *common.JobExecuteResult
			jobLocker *JobLocker
		)
		// 任务结果
		result = &common.JobExecuteResult{
			ExecuteInfo: info,
			Output:      make([]byte, 0),
		}

		//初始化分布式锁
		jobLocker = CreateJobLocker(info.Job.Name)

		result.StartTime = time.Now()

		err = jobLocker.TryLock()
		defer jobLocker.Unlock()

		if err != nil {
			result.Err = err
			result.EndTime = time.Now()
		} else {
			result.StartTime = time.Now()
			// 执行shell命令
			cmd = exec.CommandContext(info.CancelCtx, "/bin/bash", "-c", info.Job.Command)

			// 执行并捕获输出
			output, err = cmd.CombinedOutput()

			// 记录任务结束时间
			result.EndTime = time.Now()
			result.Output = output
			result.Err = err
		}

		GlobalScheduler.PushJobResult(result)
	}()
}

func NewExecutor() (err error) {
	GlobalExecutor = &Executor{}
	return
}
