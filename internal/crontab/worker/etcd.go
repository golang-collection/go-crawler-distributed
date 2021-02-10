package worker

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/go-acme/lego/v3/log"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crontab/common"
	"net"
	"time"
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

// 获取本机网卡IP
func getLocalIP() (ipv4 string, err error) {
	var (
		addrs []net.Addr
		addr net.Addr
		ipNet *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String()	// 192.168.1.1
				return
			}
		}
	}
	err = common.ERR_NO_LOCAL_IP_FOUND
	return
}

func KeepOnline(){
	var (
		ip string
		regKey string
		leaseGrantResp *clientv3.LeaseGrantResponse
		err error
		keepAliveChan <- chan *clientv3.LeaseKeepAliveResponse
		keepAliveResp *clientv3.LeaseKeepAliveResponse
		cancelCtx context.Context
		cancelFunc context.CancelFunc
	)
	ip, err = getLocalIP()
	if err != nil{
		log.Println("ip获取失败", err)
		return
	}
	for {
		// 注册路径
		regKey = common.JOB_WORKER_DIR + ip

		cancelFunc = nil

		// 创建租约
		if leaseGrantResp, err = global.EtcdLease.Grant(context.TODO(), 10); err != nil {
			goto RETRY
		}

		// 自动续租
		if keepAliveChan, err = global.EtcdLease.KeepAlive(context.TODO(), leaseGrantResp.ID); err != nil {
			goto RETRY
		}

		cancelCtx, cancelFunc = context.WithCancel(context.TODO())

		// 注册到etcd
		if _, err = global.EtcdKV.Put(cancelCtx, regKey, "", clientv3.WithLease(leaseGrantResp.ID)); err != nil {
			goto RETRY
		}

		// 处理续租应答
		for {
			select {
			case keepAliveResp = <- keepAliveChan:
				if keepAliveResp == nil {	// 续租失败
					goto RETRY
				}
			}
		}

	RETRY:
		time.Sleep(1 * time.Second)
		if cancelFunc != nil {
			cancelFunc()
		}
	}
}

func CreateJobLocker(jobName string) (jobLocker *JobLocker) {
	jobLocker = NewJobLocker(jobName)
	return
}
