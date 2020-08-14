package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/fetcher"
)

/**
* @Author: super
* @Date: 2020-08-14 20:46
* @Description:
**/
func main() {
	//var wg sync.WaitGroup
	//wg.Add(200)
	//
	//for i := 0; i <= 980; i = i + 20 {
	//	go func() {
	//		//`https://book.douban.com/tag/`
	//		contents, _ := fetcher.Fetch("https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4?start=" + strconv.Itoa(i) + "&type=T")
	//		parser.ParseTagList(contents, crawerConfig.TAG_URL)
	//		wg.Done()
	//	}()
	//	time.Sleep(5 * time.Second)
	//}
	//wg.Wait()
	url := "https://book.douban.com/tag/"
	contents, _ := fetcher.Fetch(url)
	parser.ParseTagList(contents, crawerConfig.TagUrl, url)
}
