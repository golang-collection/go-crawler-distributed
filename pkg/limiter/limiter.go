package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"

	"time"
)

/**
* @Author: super
* @Date: 2020-09-23 20:51
* @Description: 限流器接口
**/

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
