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

func ParseAndStorage(contents []byte) error {
	book := &model.Book{}
	err := book.UnmarshalJSON(contents)
	if err != nil {
		return err
	}

	err = DBOperation.InsertBook(book)
	return err
}
