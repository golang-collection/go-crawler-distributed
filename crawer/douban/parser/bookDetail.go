package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/model"
	"go-crawler-distributed/mq/mqTools"
	"go-crawler-distributed/tools"
	"go-crawler-distributed/unifiedLog"
	"go.uber.org/zap"
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
var DateRe = regexp.MustCompile(`([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8])))`)
var priceRe = regexp.MustCompile(`[0-9]+[.]?[0-9]*`)

var logger = unifiedLog.GetLogger()

func ParseBookDetail(contents []byte, queueName string, url string) {

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		logger.Error("new doc reader error", zap.Error(err))
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
			dateMatch := DateRe.FindAllSubmatch([]byte(v), -1)
			if len(dateMatch) == 0 {
				v = "2006-01-02"
			}
			if v == "" {
				v = "2006-01-02"
			}
			book.PublishYear = v
		case k == "副标题:":
			book.SubTitle = v
		case k == "原作名:":
			book.OriginalName = v
		case k == "定价:":
			priceMatch := priceRe.Find([]byte(v))
			if len(priceMatch) == 0 {
				v = "0"
			} else {
				v = string(priceMatch)
			}
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

	//Book结构体转json
	bookJson, err := tools.BookToJson(book)
	if err != nil {
		logger.Error("book to json error", zap.Error(err))
	} else {
		//将解析到的图书详细信息URL放到消息队列
		bookDetail.PublishSimple(bookJson)
	}
}
