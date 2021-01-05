package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2020-09-23 21:01
* @Description: 用于处理响应超时，请求超过规定时间则停止执行
**/

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
