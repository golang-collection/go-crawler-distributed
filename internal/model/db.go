package model

/**
* @Author: super
* @Date: 2020-11-18 15:07
* @Description:
**/

type Model struct {
	CreatedOn  string `gorm:"column:created_on" json:"created_on"`
	CreatedBy  string `gorm:"column:created_by" json:"created_by"`
	DeletedOn  string `gorm:"column:deleted_on" json:"deleted_on"`
	ModifiedBy string `gorm:"column:modified_by" json:"modified_by"`
	ModifiedOn string `gorm:"column:modified_on" json:"modified_on"`
	ID         string `gorm:"column:id;primary_key" json:"id"`
	IsDel      int    `gorm:"column:is_del" json:"is_del"`
}
