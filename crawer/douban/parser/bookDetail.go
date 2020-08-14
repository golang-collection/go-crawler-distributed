package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/model"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/tools"
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
var re = regexp.MustCompile(`<span class="pl"[^>]*>([^<]+)</span[^>]*>([^<]+)<`)
var re1 = regexp.MustCompile(`<span class="pl"[^>]*>([^<]+)</span>[^>]*>([^<]+)<`)

func ParseBookDetail(contents []byte, queueName string, url string) {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatal(err)
	}

	//初始化消息队列
	bookDetail := mqTools.NewRabbitMQSimple(queueName)
	book := &model.Book{}
	book.Url = url

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
	allSubmatch := re.FindAllSubmatch(contents, -1)
	for _, m := range allSubmatch {
		k := string(m[1])
		k = strings.TrimSpace(k)
		v := string(m[2])
		v = strings.TrimSpace(v)
		switch {
		case k == "ISBN:":
			book.ISBN = v
		case k == "出版年:":
			//TODO 有的包含数字
			if len(v) < 10 {
				v = "2006-01-02"
			}
			book.PublishYear = v
		case k == "副标题:":
			book.SubTitle = v
		case k == "原作名:":
			book.OriginalName = v
		case k == "定价:":
			//TODO 有的没有元
			v = v[:len(v)-3]
			p, _ := strconv.ParseFloat(v, 64)
			book.Price = p
		case k == "装帧:":
			book.Layout = v
		case k == "页数:":
			p, _ := strconv.Atoi(v)
			book.Pages = p
		case k == "出版社:":
			book.Publish = v
		}
	}
	allSubmatch = re1.FindAllSubmatch(contents, -1)
	for _, m := range allSubmatch {
		k := string(m[1])
		k = strings.TrimSpace(k)
		v := string(m[2])
		v = strings.TrimSpace(v)
		switch {
		case k == "丛书:":
			book.Series = v
		case k == "作者:":
			book.Author = v
		case k == "出品方:":
			book.Producer = v
		}
	}

	//评分
	result = dom.Find("strong")
	score, _ := strconv.ParseFloat(strings.TrimSpace(result.Text()), 64)
	book.Score = score
	//fmt.Println(score)

	//评价人数
	result = dom.Find("a[class=rating_people]")
	length := len(result.Text())
	if length <= 9 {
		book.Comments = 0
	} else {
		comment := result.Text()[:length-9]
		comments, _ := strconv.Atoi(comment)
		book.Comments = comments
	}
	//fmt.Println(comments)

	//短评
	result = dom.Find("div[class=indent]+p")
	commentUrl, _ := result.Children().Attr("href")
	book.CommentUrl = commentUrl
	//fmt.Println(commentUrl)

	//fmt.Println(book)

	//Book结构体转json
	bookJson, err := tools.BookToJson(book)
	if err != nil {
		log.Fatalln(err)
	} else {
		//将解析到的图书详细信息URL放到消息队列
		bookDetail.PublishSimple(bookJson)
		fmt.Println(bookJson)
	}
}
