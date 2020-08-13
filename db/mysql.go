package db

/**
* @Author: super
* @Date: 2020-08-11 16:09
* @Description: gorm自动映射操作stories表
**/

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-crawler-distributed/db/dbTools"
	"go-crawler-distributed/model"
)

func SelectStoryById(id int) (*model.Story, error) {
	db := dbtools.GetDB()

	story := &model.Story{}

	result := db.Where("story_id = ?", id).Find(story)
	if result.RecordNotFound() {
		return nil, errors.New("wrong id")
	}

	return story, nil
}

func SelectStoryRandom() (*model.Story, error) {
	db := dbtools.GetDB()

	story := &model.Story{}

	// 原生 SQL
	rows, err := db.Raw("select story_id, author, story from stories ORDER BY RAND() LIMIT 1").Rows()
	if err != nil {
		return nil, errors.New("random select fail")
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&story.StoryID, &story.Author, &story.Story)
		if err != nil {
			return nil, errors.New("random select fail")
		}
	}

	return story, err
}

func InsertStory(story *model.Story) error {
	db := dbtools.GetDB()

	result := db.Create(story)

	if result.RowsAffected == int64(0) {
		return errors.New("insert error")
	}

	return nil
}

func UpdateStory(story *model.Story) error {
	db := dbtools.GetDB()

	result := db.Model(story).Where("story_id = ?", story.StoryID).Updates(story)

	if result.RowsAffected == int64(0) {
		return errors.New("update error")
	}

	return nil
}

func DeleteStory(id int) error {
	db := dbtools.GetDB()

	result := db.Where("story_id = ?", id).Delete(model.Story{})

	if result.RowsAffected == int64(0) {
		return errors.New("delete error")
	}

	return nil
}

func CountStory() (int, error) {
	db := dbtools.GetDB()

	var count int

	result := db.Model(&model.Story{}).Count(&count)

	if result.RecordNotFound() {
		return -1, errors.New("count error")
	}

	return count, nil
}
