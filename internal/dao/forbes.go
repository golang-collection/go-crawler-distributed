package dao

import (
	"errors"
	"github.com/jinzhu/gorm"

	"go-crawler-distributed/internal/model"
	"go-crawler-distributed/pkg/app"
)

/**
* @Author: super
* @Date: 2020-12-30 11:21
* @Description:
**/

type Forbes struct {
	NameEn         string `json:"name_en"`
	Wealth         int    `json:"wealth"`
	SourceOfWealth string `json:"source_of_wealth"`
	Region         string `json:"region"`
	ModifiedOn     string `json:"modified_on"`
	ID             string `json:"id"`
	Rank           int    `json:"rank"`
	Name           string `json:"name"`
}

type IForbes interface {
	SelectAll() ([]*Forbes, error)
	SelectList(page, pageSize int) ([]*Forbes, error)
}

type ForbesManager struct {
	table string
	conn  *gorm.DB
}

func NewForbesManager(table string, conn *gorm.DB) IForbes {
	return &ForbesManager{table: table, conn: conn}
}

func (m *ForbesManager) SelectAll() ([]*Forbes, error) {
	var f []*model.Forbes
	if err := m.conn.Find(&f).Error; err != nil {
		return nil, errors.New("select all forbes error")
	}
	forbess := make([]*Forbes, 0)
	for _, forbes := range f {
		temp := &Forbes{
			ID:             forbes.ID,
			Rank:           forbes.Rank,
			Name:           forbes.Name,
			NameEn:         forbes.NameEn,
			Wealth:         forbes.Wealth,
			SourceOfWealth: forbes.SourceOfWealth,
			Region:         forbes.Region,
			ModifiedOn:     forbes.ModifiedOn,
		}
		forbess = append(forbess, temp)
	}
	return forbess, nil
}

func (m *ForbesManager) SelectList(page, pageSize int) ([]*Forbes, error) {
	pageOffset := app.GetPageOffset(page, pageSize)
	if pageOffset < 0 && pageSize < 0 {
		pageOffset = 0
		pageSize = 5
	}
	fields := []string{"id", "rank", "name", "name_en", "wealth", "source_of_wealth", "region", "modified_on"}
	rows, err := m.conn.Offset(pageOffset).Limit(pageSize).Select(fields).Table(m.table).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forbess []*Forbes
	for rows.Next() {
		forbes := &Forbes{}
		if err := rows.Scan(&forbes.ID,
			&forbes.Rank,
			&forbes.Name,
			&forbes.NameEn,
			&forbes.Wealth,
			&forbes.SourceOfWealth,
			&forbes.Region,
			&forbes.ModifiedOn); err != nil {
			return nil, err
		}
		forbess = append(forbess, forbes)
	}
	return forbess, nil
}
