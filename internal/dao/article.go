package dao

import (
	"github.com/jinzhu/gorm"
	"go-crawler-distributed/internal/model"
)

/**
* @Author: super
* @Date: 2021-01-05 15:55
* @Description:
**/

type IArticle interface {
	SaveArticle(article model.Article) (string, error)
}

type ArticleManager struct {
	table string
	conn  *gorm.DB
}

func NewArticleManager(table string, conn *gorm.DB) IArticle {
	return &ArticleManager{table: table, conn: conn}
}

func (m *ArticleManager) SaveArticle(article model.Article) (string, error) {
	return "", nil
}
