package fetcher

import (
	"bufio"
	"context"
	"fmt"
	"go-crawler-distributed/global"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

/**
* @Author: super
* @Date: 2021-01-05 14:58
* @Description:
**/

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		global.Logger.Error(context.Background(), "new request error: ", err)
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode)
	}

	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//自动判断编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		global.Logger.Error(context.Background(), "determine coder error", err)
		//默认UTF8编码
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
