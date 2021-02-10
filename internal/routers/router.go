package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/middleware"
	"go-crawler-distributed/internal/routers/job"
	"go-crawler-distributed/internal/routers/sd"
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
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.Tracing())
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	jobGroup := r.Group("/job")
	{
		jobGroup.POST("/save", job.SaveJob)
		jobGroup.POST("/delete", job.DeleteJob)
		jobGroup.GET("/list", job.ListJobs)
		jobGroup.POST("/kill", job.KillJob)
		jobGroup.GET("/log", job.JobLog)
	}
	r.GET("/worker/list", job.WorkerList)

	return r
}
