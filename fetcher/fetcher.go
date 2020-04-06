package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"go-crawler-distributed/mylog"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter

	mylog.LogInfo("fetcher", fmt.Sprintf("fetching %s", url))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//logger.Error("mylog", zap.String("err", err.Error()))
		mylog.LogError("fetcher.statusOK", errors.New("statusok Error"))
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		//log.Printf("Fetch error: %v", err)
		mylog.LogError("Fetch error", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
