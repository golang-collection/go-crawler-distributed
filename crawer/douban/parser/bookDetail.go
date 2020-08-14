package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/model"
	"log"
	"regexp"
	"strconv"
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

	//初始化消息队列
	//bookDetail := mqTools.NewRabbitMQSimple(queueName)
	book := &model.Book{}

	//封面图片
	result := dom.Find("img[title]")
	img, _ := result.Attr("src")
	book.Img = img
	//fmt.Println(img)
	//书名
	title, _ := result.Attr("alt")
	book.Title = title
	//fmt.Println(title)

	//图书信息
	result = dom.Find("div[id=info]")
	fmt.Println(result.Text())
	carModelRe := regexp.MustCompile(`作者:`)

	//newResult := strings.TrimSpace(result.Text())
	//resultList := strings.Fields(newResult)
	//fmt.Println(resultList)
	//fmt.Println(len(resultList))
	//resultMap := make(map[string]interface{})
	//for i := 0; i < len(resultList); i = i + 2 {
	//	resultMap[resultList[i]] = resultList[i+1]
	//}
	//for k, v := range resultMap {
	//	switch {
	//	case k == "丛书:":
	//		book.Series = v.(string)
	//	case k == "ISBN:":
	//		book.ISBN = v.(string)
	//	case k == "作者:":
	//		book.Author = v.(string)
	//	case k == "出品方:":
	//		book.Producer = v.(string)
	//	case k == "出版年:":
	//		book.PublishYear = v.(string)
	//	case k == "出版社:":
	//		book.Publish = v.(string)
	//	case k == "副标题:":
	//		book.SubTitle = v.(string)
	//	case k == "原作名:":
	//		book.OriginalName = v.(string)
	//	case k == "定价:":
	//		p, _ := strconv.ParseFloat(strings.TrimSpace(v.(string)), 64)
	//		book.Price = p
	//	case k == "装帧:":
	//		book.Layout = v.(string)
	//	case k == "页数:":
	//		p, _ := strconv.Atoi(strings.TrimSpace(v.(string)))
	//		book.Pages = p
	//	}
	//}

	//评分
	result = dom.Find("strong")
	score, _ := strconv.ParseFloat(strings.TrimSpace(result.Text()), 64)
	book.Score = score
	//fmt.Println(score)

	//评价人数
	result = dom.Find("a[class=rating_people]")
	length := len(result.Text())
	comment := result.Text()[:length-9]
	comments, _ := strconv.Atoi(comment)
	book.Comments = comments
	//fmt.Println(comments)

	//短评
	result = dom.Find("div[class=indent]+p")
	commentUrl, _ := result.Children().Attr("href")
	book.CommentUrl = commentUrl
	//fmt.Println(commentUrl)

	fmt.Println(book)

	//Book结构体转json
	//bookJson, err := tools.BookToJson(book)
	//if err != nil {
	//	log.Fatalln(err)
	//} else {
	//	//将解析到的图书详细信息URL放到消息队列
	//	bookDetail.PublishSimple(bookJson)
	//}
}
