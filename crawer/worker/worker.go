package worker

import "go-crawler-distributed/crawer/fetcher"

/**
* @Author: super
* @Date: 2020-08-16 07:55
* @Description:
**/
func Worker(r Request) {
	contents, _ := fetcher.Fetch(r.Url)
	r.Parser.Parse(contents, r.Url)
}
