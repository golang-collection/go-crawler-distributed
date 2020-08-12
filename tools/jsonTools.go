package tools

import (
	"encoding/json"
	"go-crawler-distributed/model"
)

/**
* @Author: super
* @Date: 2020-08-11 19:24
* @Description:
**/

func ActivityToJson(story *model.Story) (string, error) {
	str, err := json.Marshal(story)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func JsonToActivity(str string) (*model.Story, error) {
	story := &model.Story{}
	err := json.Unmarshal([]byte(str), story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
