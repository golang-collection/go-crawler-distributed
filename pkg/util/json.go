package util

import "encoding/json"

/**
* @Author: super
* @Date: 2020-12-06 17:43
* @Description:
**/

func EncodeToJson(object interface{}) (string, error) {
	encodeBytes, err := json.Marshal(object)
	if err != nil {
		return "", err
	}
	return string(encodeBytes), nil
}

func DecodeToStruct(input string) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
