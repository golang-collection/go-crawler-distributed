package util

import (
	"encoding/json"
	"fmt"
)

/**
* @Author: super
* @Date: 2020-11-27 16:37
* @Description:
**/
type RGB struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

func (rgb *RGB) ToString() string {
	result := "#"
	if rgb.Red < 16 {
		result += fmt.Sprintf("0%x", rgb.Red)
	} else {
		result += fmt.Sprintf("%x", rgb.Red)
	}
	if rgb.Green < 16 {
		result += fmt.Sprintf("0%x", rgb.Green)
	} else {
		result += fmt.Sprintf("%x", rgb.Green)
	}
	if rgb.Blue < 16 {
		result += fmt.Sprintf("0%x", rgb.Blue)
	} else {
		result += fmt.Sprintf("%x", rgb.Blue)
	}
	return result
}

func RgbToHex(rgb string) (string, error) {
	rgbStruct := &RGB{}
	err := json.Unmarshal([]byte(rgb), rgbStruct)
	if err != nil {
		return "", err
	}
	return rgbStruct.ToString(), nil
}
