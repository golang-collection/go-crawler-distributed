package crawer

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-crawler-distributed/db"
	"go-crawler-distributed/model"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

/**
* @Author: super
* @Date: 2020-08-12 19:47
* @Description:
**/

func main() {
	resp, err := http.Get("https://s.weibo.com/weibo/%2523%25E8%25AE%25B2%25E7%25BB%2599%25E5%25A5%25B3%25E6%259C%258B%25E5%258F%258B%25E7%259A%2584%25E7%259D%25A1%25E5%2589%258D%25E5%25B0%258F%25E6%2595%2585%25E4%25BA%258B%2523?topnav=1&wvr=6&b=1&page=2")
	if err != nil {
		panic("error")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	ParseStory(all)
}

//自动判断编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func ParseStory(contents []byte) {
	re := regexp.MustCompile(`<p class="txt" node-type="feed_list_content_full" nick-name=[^"]*"([^"]+)" style="display: none">`)
	matches := re.FindAllSubmatch(contents, -1)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		log.Fatal(err)
	}

	result := dom.Find("p[style]")
	result.Each(func(i int, selection *goquery.Selection) {
		i = i - 1
		if i < 0 {
			i = 0
		}
		fmt.Println(string(matches[i][1]) + ":" + selection.Text())
		story := &model.Story{}
		story.Author = string(matches[i][1])
		story.Story = selection.Text()

		err = db.InsertStory(story)
		if err != nil {
			panic(err)
		}
	})
}
