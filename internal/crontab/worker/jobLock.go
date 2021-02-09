package worker

import (
	"context"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crontab/common"
)

/**
* @Author: super
* @Date: 2021-02-07 18:53
* @Description:
**/

type JobLocker struct {
	JobName    string
	CancelFunc context.CancelFunc
	LeaseId    clientv3.LeaseID
	IsLocked   bool
}

func (jobLocker *JobLocker) TryLock() (err error) {
	var (
		leaseGrantResp *clientv3.LeaseGrantResponse
		cancelCtx      context.Context
		cancelFunc     context.CancelFunc
		leaseId        clientv3.LeaseID
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		txn            clientv3.Txn
		lockKey        string
		txnResp        *clientv3.TxnResponse
	)
	if leaseGrantResp, err = global.EtcdLease.Grant(context.TODO(), 5); err != nil {
		return
	}
	cancelCtx, cancelFunc = context.WithCancel(context.TODO())
	leaseId = leaseGrantResp.ID

	if keepRespChan, err = global.EtcdLease.KeepAlive(cancelCtx, leaseId); err != nil {
		cancelFunc()
		global.EtcdLease.Revoke(context.TODO(), leaseId)
		return
	}

	go func() {
		var (
			keepResp *clientv3.LeaseKeepAliveResponse
		)
		for {
			select {
			case keepResp = <-keepRespChan: // 自动续租的应答
				if keepResp == nil {
					return
				}
			}
		}
	}()

	txn = global.EtcdKV.Txn(context.TODO())
	lockKey = common.JOB_LOCK_DIR + jobLocker.JobName

	txn.If(clientv3.Compare(clientv3.CreateRevision(lockKey), "=", 0)).
		Then(clientv3.OpPut(lockKey, "", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet(lockKey))

	if txnResp, err = txn.Commit(); err != nil {
		cancelFunc()
		global.EtcdLease.Revoke(context.TODO(), leaseId)
		return
	}

	if !txnResp.Succeeded {
		err = errors.New("锁被占用")
		cancelFunc()
		global.EtcdLease.Revoke(context.TODO(), leaseId)
		return
	}
	// 抢锁成功
	jobLocker.LeaseId = leaseId
	jobLocker.CancelFunc = cancelFunc
	jobLocker.IsLocked = true
	return
}

func (jobLocker *JobLocker) Unlock() {
	if jobLocker.IsLocked {
		jobLocker.CancelFunc()                                     // 取消我们程序自动续租的协程
		global.EtcdLease.Revoke(context.TODO(), jobLocker.LeaseId) // 释放租约
	}
}

func NewJobLocker(jobName string) *JobLocker {
	return &JobLocker{
		JobName: jobName,
	}
}
