package main

import (
	"fmt"
	"go-crawler-distributed/elastic/elasticOperation"
	"go-crawler-distributed/model"

	//"go-crawler-distributed/model"
)

/**
* @Author: super
* @Date: 2020-09-01 16:50
* @Description:
**/

func main() {
	//elasticOperation.IndexExist("test_els")

	article := model.Article{
		Title:"hello1",
		Url:"good1",
		Genres:[]string{"go", "elk"},
		Content:"this is a test1",
	}
	id, err := elasticOperation.SaveInfo("test_els", article)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(id)
	info, err := elasticOperation.GetInfo("test_els", id)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(info)
}