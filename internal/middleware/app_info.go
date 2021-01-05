package middleware

import "github.com/gin-gonic/gin"

/**
* @Author: super
* @Date: 2020-09-23 20:49
* @Description:
**/

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "superTools")
		c.Set("app_version", "0.0.3")
		c.Next()
	}
}
