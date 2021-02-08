package worker

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crontab/common"
)

/**
* @Author: super
* @Date: 2021-02-07 17:14
* @Description:
**/

func WatchJobs(ctx context.Context) (err error) {
	getResp, err := global.EtcdKV.Get(ctx, common.JOB_SAVE_DIR, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for i := 0; i < len(getResp.Kvs); i++ {
		job := &common.Job{}
		err := job.UnmarshalJSON(getResp.Kvs[i].Value)
		if err == nil {
			jobEvent := common.BuildJobEvent(common.JOB_EVENT_SAVE, job)
			//添加到任务调度器
			GlobalScheduler.PushJobEvent(jobEvent)
		}
	}

	revision := getResp.Header.Revision

	go func(watchStartRevision int64) {
		watchChan := global.EtcdWatcher.Watch(ctx, common.JOB_SAVE_DIR,
			clientv3.WithRev(watchStartRevision),
			clientv3.WithPrefix())
		for watchResp := range watchChan {
			for _, watchEvent := range watchResp.Events {
				var jobEvent *common.JobEvent
				switch watchEvent.Type {
				case mvccpb.PUT:
					job := &common.Job{}
					err := job.UnmarshalJSON(watchEvent.Kv.Value)
					if err != nil {
						continue
					}
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, job)
				case mvccpb.DELETE:
					jobName := common.ExtractJobName(string(watchEvent.Kv.Key))
					job := &common.Job{
						Name: jobName,
					}
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_DELETE, job)
				}
				//将变化情况推送给调度器
				GlobalScheduler.PushJobEvent(jobEvent)
			}
		}
	}(revision + 1)
	return
}

func WatchKiller(ctx context.Context) {
	go func() {
		// 监听/cron/killer/目录的变化
		watchChan := global.EtcdWatcher.Watch(ctx, common.JOB_KILLER_DIR, clientv3.WithPrefix())
		// 处理监听事件
		for watchResp := range watchChan {
			for _, watchEvent := range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 杀死任务事件
					jobName := common.ExtractKillerName(string(watchEvent.Kv.Key))
					job := &common.Job{Name: jobName}
					jobEvent := common.BuildJobEvent(common.JOB_EVENT_KILL, job)
					// 事件推给scheduler
					GlobalScheduler.PushJobEvent(jobEvent)
				case mvccpb.DELETE: // killer标记过期, 被自动删除
				}
			}
		}
	}()
}
