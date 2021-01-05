package parser

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

/**
* @Author: super
* @Date: 2021-01-05 15:05
* @Description:
**/

func ParseArticleUrl(contents []byte, url string) ([]string ,error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		return []string{}, err
	}
	urls := make([]string, 0)

	result := dom.Find("a[rel=bookmark]")
	result.Each(func(i int, selection *goquery.Selection) {
		href, exist := selection.Attr("href")
		if exist{
			urls = append(urls, href)
		}
	})
	return urls, nil
}