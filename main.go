package main

import (
	"fmt"
	"go-crawler-distributed/db"
)

/**
* @Author: super
* @Date: 2020-08-12 16:06
* @Description:
**/

func main() {
	story, err := db.SelectStoryRandom()
	if err != nil {
		return
	}
	fmt.Println(story)
}
