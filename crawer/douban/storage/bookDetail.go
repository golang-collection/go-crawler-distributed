package storage

import (
	"go-crawler-distributed/db/DBOperation"
	"go-crawler-distributed/tools"
)

/**
* @Author: super
* @Date: 2020-08-16 08:56
* @Description:
**/

func ParseAndStorage(contents []byte) error {

	book, err := tools.JsonToBook(string(contents))
	if err != nil {
		return err
	}

	err = DBOperation.InsertBook(book)
	return err
}
