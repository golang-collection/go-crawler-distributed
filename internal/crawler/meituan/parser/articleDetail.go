package parser

import (
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/internal/model"
	"go-crawler-distributed/pkg/util"
	"strings"
)

/**
* @Author: super
* @Date: 2021-01-05 15:12
* @Description:
**/

func ParseArticleDetail(contents []byte, url string) ([]string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		return []string{}, err
	}
	article := &model.Article{}

	result := dom.Find("a[rel=bookmark]")
	article.Url = url

	title := result.Text()
	article.Title = title

	s, err := util.ZipString(contents)
	if err != nil{
		return []string{}, err
	}
	article.Content = s

	result = dom.Find("a[rel=tag]")
	result.Each(func(i int, selection *goquery.Selection) {
		tag := selection.Text()
		article.Genres = append(article.Genres, tag)
	})

	articles := make([]string, 0)

	//Article结构体转json
	bytes, err := article.MarshalJSON()
	if err != nil {
		return []string{}, err
	} else {
		articleJson := string(bytes)
		articles = append(articles, articleJson)
	}
	return articles, nil
}