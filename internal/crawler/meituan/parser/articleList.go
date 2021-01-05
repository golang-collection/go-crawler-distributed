package parser

import (
	"strconv"
)

/**
* @Author: super
* @Date: 2021-01-05 15:10
* @Description:
**/


func ParseArticleList(contents []byte, url string) ([]string, error){
	result := make([]string, 0)
	for i := 2; i<22;i++{
		url := "https://tech.meituan.com//page/"+ strconv.Itoa(i) +".html"
		result = append(result, url)
	}
	return result, nil
}