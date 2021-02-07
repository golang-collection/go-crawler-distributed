package worker

import (
	"context"
	"github.com/coreos/etcd/clientv3"
)

/**
* @Author: super
* @Date: 2021-02-07 18:53
* @Description:
**/

type JobLocker struct {
	jobName string
	cancelFunc context.CancelFunc
	leaseId clientv3.LeaseID
	isLocked bool
}

func InitJobLocker(){

}