package main

import (
	"go-crawler-distributed/crawer/crawerConfig"
	"go-crawler-distributed/crawer/douban/parser"
	"go-crawler-distributed/crawer/fetcher"
	"strconv"
	"sync"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-12 19:47
* @Description:
**/

func main() {
	var wg sync.WaitGroup
	wg.Add(6)

	for i := 0; i <= 100; i = i + 20 {
		go func() {
			contents, _ := fetcher.Fetch("https://book.douban.com/tag/%E5%B0%8F%E8%AF%B4?start=" + strconv.Itoa(i) + "&type=T")
			parser.ParseBookList(contents, crawerConfig.BOOK_DETAIL_URL)
			wg.Done()
		}()
		time.Sleep(5 * time.Second)
	}
	wg.Wait()
}
