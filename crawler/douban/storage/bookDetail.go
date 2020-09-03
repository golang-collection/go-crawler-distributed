package storage

import (
	"go-crawler-distributed/db/DBOperation"
	"go-crawler-distributed/model"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
)

/**
* @Author: super
* @Date: 2020-08-16 08:56
* @Description:
**/

func ParseAndStorage(contents []byte, _ string, _ string) {
	book := &model.Book{}
	err := book.UnmarshalJSON(contents)
	if err != nil {
		unifiedLog.GetLogger().Error("book unmarshal json error", zap.Error(err))
		return
	}

	err = DBOperation.InsertBook(book)
	if err != nil{
		unifiedLog.GetLogger().Error("book save info error", zap.Error(err))
	}
}
