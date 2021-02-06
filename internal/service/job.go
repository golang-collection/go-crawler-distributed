package service

/**
* @Author: super
* @Date: 2021-02-06 19:18
* @Description:
**/

type SaveJobRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=2,max=4294967295"`
	Command  string `json:"command" form:"command" binding:"required,min=2,max=4294967295"`
	CronExpr string `json:"cron_expr" form:"cron_expr" binding:"required,min=2,max=4294967295"`
}

type DeleteJobRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=4294967295"`
}

type KillJobRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=4294967295"`
}
