package main

import (
	"fmt"
	"go-crawler-distributed/elastic/elasticOperation"
	"go-crawler-distributed/service/watchConfig"
)

/**
* @Author: super
* @Date: 2020-09-01 21:12
* @Description:
**/

func main() {
	index, _ := watchConfig.GetElasticIndex()
	articles, _ := elasticOperation.SearchInfo(index, "genres","后台")
	fmt.Println(articles)
}