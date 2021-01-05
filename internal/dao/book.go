package dao

import (
	"github.com/jinzhu/gorm"
	"go-crawler-distributed/internal/model"
)

/**
* @Author: super
* @Date: 2021-01-05 18:55
* @Description:
**/

type IBook interface {
	SaveBook(book model.Book) (string, error)
}

type BookManager struct {
	table string
	conn  *gorm.DB
}

func NewBookManager(table string, conn *gorm.DB) IBook {
	return &BookManager{table: table, conn: conn}
}

func (m *BookManager) SaveBook(book model.Book) (string, error) {
	return "", nil
}
