package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-crawler-distributed/global"
)

/**
* @Author: super
* @Date: 2021-02-06 16:34
* @Description:
**/

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		//r.Use(middleware.AccessLog())
		//r.Use(middleware.Recovery())
	}

	r.POST("/job/save", )

	return r
}