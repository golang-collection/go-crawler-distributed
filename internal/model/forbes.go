package model

/**
* @Author: super
* @Date: 2020-12-30 11:18
* @Description: 福布斯排行榜
**/

type Forbes struct {
	NameEn         string `gorm:"column:name_en" json:"name_en"`
	Wealth         int    `gorm:"column:wealth" json:"wealth"`
	SourceOfWealth string `gorm:"column:source_of_wealth" json:"source_of_wealth"`
	Region         string `gorm:"column:region" json:"region"`
	ModifiedOn     string `gorm:"column:modified_on" json:"modified_on"`
	ID             string `gorm:"column:id;primary_key" json:"id"`
	Rank           int    `gorm:"column:rank" json:"rank"`
	Name           string `gorm:"column:name" json:"name"`
}

// TableName sets the insert table name for this struct type
func (f *Forbes) TableName() string {
	return "forbes_list"
}
