package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"go-crawler-distributed/pkg/setting"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-06 18:22
* @Description:
**/

func NewEtcdEngine(etcdSetting *setting.EtcdSettingS) (client *clientv3.Client, kv clientv3.KV, lease clientv3.Lease, err error) {
	config := clientv3.Config{
		Endpoints:   []string{etcdSetting.Endpoint},
		DialTimeout: time.Duration(etcdSetting.DialTimeout) * time.Millisecond,
	}
	if client, err = clientv3.New(config); err != nil {
		return
	}
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	return
}

