package worker

import "go-crawler-distributed/internal/crawler/fetcher"

/**
* @Author: super
* @Date: 2021-01-05 15:03
* @Description:
**/

func Worker(r Request) ([]string, error){
	contents, err := fetcher.Fetch(r.Url)
	if err != nil{
		return []string{}, err
	}
	return r.Parser.Parse(contents, r.Url)
}