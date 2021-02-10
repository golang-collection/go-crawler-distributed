package service

import (
	"context"
	"fmt"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/crontab/common"
	"go-crawler-distributed/pkg/app"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
* @Author: super
* @Date: 2021-02-06 19:18
* @Description:
**/

type SaveJobRequest struct {
	Name     string `json:"name" form:"name" binding:"required,min=2,max=4294967295"`
	Command  string `json:"command" form:"command" binding:"required,min=2,max=4294967295"`
	CronExpr string `json:"cronExpr" form:"cronExpr" binding:"required,min=2,max=4294967295"`
}

type DeleteJobRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=4294967295"`
}

type KillJobRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=4294967295"`
}

type JobLogRequest struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=4294967295"`
}

type IJobLogService interface {
	GetLogList(param *JobLogRequest, pager *app.Pager) (*common.JobLog, error)
}

func GetLogList(param *JobLogRequest, pager *app.Pager) ([]*common.JobLog, error) {
	filter := &common.JobLogFilter{
		JobName: param.Name,
	}
	fmt.Println(param.Name)
	logSort := &common.SortLogByStartTime{
		SortOrder: -1,
	}
	if pager.PageSize == 0 {
		pager.PageSize = 20
	}
	collection := global.MongoDBEngine.Database("cron").Collection("log")

	skip := int64(pager.Page)
	limit := int64(pager.PageSize)
	op := &options.FindOptions{
		Sort:  logSort,
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := collection.Find(context.TODO(), filter, op)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	result := make([]*common.JobLog, 0)
	for cursor.Next(context.TODO()) {
		jobLog := &common.JobLog{}
		if err := cursor.Decode(jobLog); err != nil {
			continue
		}
		result = append(result, jobLog)
	}
	return result, nil
}
