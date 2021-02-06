package common

/**
* @Author: super
* @Date: 2021-02-06 18:44
* @Description:
**/

type Job struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	CronExpr string `json:"cron_expr"`
}
