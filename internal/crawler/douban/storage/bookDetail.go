package storage

import (
	"context"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/dao"
	"go-crawler-distributed/internal/model"
)

/**
* @Author: super
* @Date: 2020-08-16 08:56
* @Description:
**/

func ParseAndStorage(contents []byte, _ string, _ string) {
	book := model.Book{}
	err := book.UnmarshalJSON(contents)
	if err != nil {
		global.Logger.Error(context.Background(), err)
		return
	}

	bookManager := dao.NewBookManager("books", global.DBEngine)

	_, err = bookManager.SaveBook(book)
	if err != nil {
		global.Logger.Error(context.Background(), err)
	}
}
