package conn

/**
* @Author: super
* @Date: 2020-08-11 16:09
* @Description: gorm自动映射操作activities表
**/

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	tools "go-crawler-distributed/dbTools"
	"go-crawler-distributed/model"
)

func SelectStoryById(id int) (*model.Story, error) {
	db := tools.GetDB()

	activity := &model.Story{}

	result := db.Where("activity_id = ?", id).Find(activity)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}

	return activity, nil
}

func SelectStoryRandom() (*model.Story, error) {
	db := tools.GetDB()

	activity := &model.Story{}

	result := db.Take(activity)

	if result.RecordNotFound() {
		return nil, errors.New("empty table")
	}

	return activity, nil
}

func InsertStory(activity *model.Story) error {
	db := tools.GetDB()

	result := db.Create(activity)

	if result.RowsAffected == int64(0) {
		return errors.New("insert error")
	}

	return nil
}

func UpdateStory(activity *model.Story) error {
	db := tools.GetDB()

	result := db.Model(activity).Where("activity_id = ?", activity.Story).Updates(activity)

	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}

	return nil
}

func DeleteStory(id int) error {
	db := tools.GetDB()

	result := db.Where("activity_id = ?", id).Delete(model.Story{})

	if result.RowsAffected == int64(0) {
		return errors.New("delete error")
	}

	return nil
}

func CountStory() (int, error) {
	db := tools.GetDB()

	var count int

	result := db.Model(&model.Story{}).Count(&count)

	if result.RecordNotFound() {
		return -1, errors.New("count error")
	}

	return count, nil
}
