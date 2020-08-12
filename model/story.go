package model

/**
* @Author: super
* @Date: 2020-08-12 17:16
* @Description:
**/

type Story struct {
	StoryID int    `gorm:"column:story_id" gorm:"PRIMARY_KEY" json:"story_id"`
	author  string `gorm:"column:author" json:"author"`
	story   string `gorm:"column:story" gorm:"type:text" json:"story"`
}
