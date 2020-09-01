package storage

import (
	"go-crawler-distributed/db/DBOperation"
	"go-crawler-distributed/model"
)

/**
* @Author: super
* @Date: 2020-08-16 08:56
* @Description:
**/

func ParseAndStorage(data interface{}) error {

	contents := data.([]byte)

	book := &model.Book{}
	err := book.UnmarshalJSON(contents)
	if err != nil {
		return err
	}

	err = DBOperation.InsertBook(book)
	return err
}
