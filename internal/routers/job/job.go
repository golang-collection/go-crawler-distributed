package job

import (
	"github.com/gin-gonic/gin"
	"go-crawler-distributed/pkg/app"
)

/**
* @Author: super
* @Date: 2021-02-06 16:44
* @Description:
**/

// 将任务保存到etcd中
func SaveJob(c *gin.Context){
	response := app.NewResponse(c)

}