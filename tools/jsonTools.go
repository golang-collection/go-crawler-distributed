package tools

import (
	"GORM/model"
	"encoding/json"
)

/**
* @Author: super
* @Date: 2020-08-11 19:24
* @Description:
**/

func ActivityToJson(activity *model.Activity) (string, error) {
	str, err := json.Marshal(activity)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func JsonToActivity(str string) (*model.Activity, error) {
	activity := &model.Activity{}
	err := json.Unmarshal([]byte(str), activity)
	if err != nil {
		return nil, err
	}
	return activity, nil
}
