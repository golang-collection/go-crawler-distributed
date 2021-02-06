package job

import (
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
	if err != nil{
		global.Logger.Errorf(c, "app.EtcdSaveJob errs: %v", errs)
		response.ToErrorResponse(errcode.ErrorSaveFail)
		return
	}
	response.ToResponse(oldJob, "存储任务成功", http.StatusOK)
}

func DeleteJob(){

}