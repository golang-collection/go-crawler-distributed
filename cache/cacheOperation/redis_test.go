package cacheOperation

import (
	"fmt"
	"go-crawler-distributed/crawer/crawerConfig"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-15 08:31
* @Description:
**/
func TestGetAllElementFromSet(t *testing.T) {
	strings, err := GetAllElementFromSet(crawerConfig.BookDetailUrl)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(strings)
}

func TestDelAllElementFromSet(t *testing.T) {
	result, err := DelAllElementFromSet(crawerConfig.BookDetailUrl)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}
