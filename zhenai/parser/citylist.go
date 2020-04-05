package parser

import (
	"go-crawler-distributed/config"
	"go-crawler-distributed/engine"
	"regexp"
)

const CITY_LIST_RE = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(CITY_LIST_RE)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:    string(m[1]),
				Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
			})
	}
	return result
}
