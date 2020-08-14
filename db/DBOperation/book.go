package DBOperation

/**
* @Author: super
* @Date: 2020-08-14 16:28
* @Description:
**/

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-crawler-distributed/db/dbTools"
	"go-crawler-distributed/model"
)

func SelectBookById(id int) (*model.Book, error) {
	db := dbtools.GetDB()

	book := &model.Book{}

	result := db.Where("book_id = ?", id).Find(book)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}

	return book, nil
}

func SelectBookRandom() (*model.Book, error) {
	db := dbtools.GetDB()

	book := &model.Book{}

	// 原生 SQL
	//TODO 完善全语句
	rows, err := db.Raw("select book_id, author from stories ORDER BY RAND() LIMIT 1").Rows()
	if err != nil {
		return nil, errors.New("random select fail")
	}
	defer rows.Close()

	for rows.Next() {
		//TODO 完善全字段
		err = rows.Scan(&book.BookID, &book.Author)
		if err != nil {
			return nil, errors.New("random select fail")
		}
	}

	return book, err
}

func InsertBook(book *model.Book) error {
	db := dbtools.GetDB()

	result := db.Create(book)

	if result.RowsAffected == int64(0) {
		return errors.New("insert error")
	}

	return nil
}

func UpdateBook(book *model.Book) error {
	db := dbtools.GetDB()

	result := db.Model(book).Where("book_id = ?", book.BookID).Updates(book)

	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}

	return nil
}

func DeleteBook(id int) error {
	db := dbtools.GetDB()

	result := db.Where("book_id = ?", id).Delete(model.Book{})

	if result.RowsAffected == int64(0) {
		return errors.New("delete error")
	}

	return nil
}

func CountBook() (int, error) {
	db := dbtools.GetDB()

	var count int

	result := db.Model(&model.Book{}).Count(&count)

	if result.RecordNotFound() {
		return -1, errors.New("count error")
	}

	return count, nil
}
