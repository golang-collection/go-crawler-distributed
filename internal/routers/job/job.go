package job

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crontab/common"
	"go-crawler-distributed/internal/crontab/master"
	"go-crawler-distributed/internal/service"
	"go-crawler-distributed/pkg/app"
	"go-crawler-distributed/pkg/errcode"
	"net/http"
)

/**
* @Author: super
* @Date: 2021-02-06 16:44
* @Description:
**/

// 将任务保存到etcd中
func SaveJob(c *gin.Context) {
	name, _ := c.GetPostForm("name")
	fmt.Println("1", c.Param("name"))
	fmt.Println("2", name)
	fmt.Println("3", c.PostForm("name"))
	fmt.Println()
	param := service.SaveJobRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	job := &common.Job{
		Name:     param.Name,
		Command:  param.Command,
		CronExpr: param.CronExpr,
	}

	oldJob, err := master.EtcdSaveJob(c, job)
	if err != nil {
		global.Logger.Errorf(c, "app.EtcdSaveJob err: %v", err)
		response.ToErrorResponse(errcode.ErrorSaveFail)
		return
	}
	response.ToResponse(oldJob, "存储任务成功", http.StatusOK)
}

func DeleteJob(c *gin.Context) {
	param := service.DeleteJobRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	oldJob, err := master.EtcdDeleteJob(c, param.Name)
	if err != nil {
		global.Logger.Errorf(c, "app.EtcdDeleteJob err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteFail)
		return
	}
	response.ToResponse(oldJob, "删除任务成功", http.StatusOK)
}

func ListJobs(c *gin.Context) {
	response := app.NewResponse(c)
	jobs, err := master.EtcdListJobs(c)
	if err != nil {
		global.Logger.Errorf(c, "app.EtcdListJobs err: %v", err)
		response.ToErrorResponse(errcode.ErrorListFail)
		return
	}
	response.ToResponse(jobs, "获取任务列表成功", http.StatusOK)
}

func KillJob(c *gin.Context) {
	param := service.KillJobRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	err := master.EtcdKillJob(c, param.Name)
	if err != nil {
		global.Logger.Errorf(c, "app.EtcdKillJob err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteFail)
		return
	}
	response.ToResponse(gin.H{}, "杀死任务成功", http.StatusOK)
}

func JobLog(c *gin.Context) {
	param := service.JobLogRequest{}
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	result, err := service.GetLogList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c, "service.GetLogList err: %v", err)
		response.ToErrorResponse(errcode.ErrorLogListFail)
		return
	}
	response.ToResponse(result, "获取日志列表成功", http.StatusOK)
}

func WorkerList(c *gin.Context) {
	response := app.NewResponse(c)
	workers, err := master.ListWorkers()
	if err != nil {
		global.Logger.Errorf(c, "appWorkerList err: %v", err)
		response.ToErrorResponse(errcode.ErrorWorkerListFail)
		return
	}
	response.ToResponse(workers, "获取worker列表成功", http.StatusOK)
}