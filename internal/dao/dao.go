package dao

import "github.com/jinzhu/gorm"

/**
* @Author: super
* @Date: 2020-09-22 09:35
* @Description: 用于统一配置DB引擎
**/

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
