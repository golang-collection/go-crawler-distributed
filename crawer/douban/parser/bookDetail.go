package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-14 14:22
* @Description:
**/

func ParseBookDetail(contents []byte, queueName string) {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatal(err)
	}

	//封面图片
	result := dom.Find("img[title]")
	fmt.Println(result.Attr("src"))
	//书名
	fmt.Println(result.Attr("alt"))

	//图书信息
	result = dom.Find("div[id=info]")
	newResult := strings.TrimSpace(result.Text())
	resultList := strings.Fields(newResult)
	for _, s := range resultList {
		fmt.Println(strings.TrimSpace(s))
	}

	//评分
	result = dom.Find("strong")
	fmt.Println(result.Text())

	//评价人数
	result = dom.Find("a[class=rating_people]")
	fmt.Println(result.Text())

	//短评
	result = dom.Find("div[class=indent]+p")
	fmt.Println(result.Children().Attr("href"))

}
