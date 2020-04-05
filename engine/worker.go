package engine

import "go-crawler-distributed/fetcher"

func worker(r Request) (ParseResult, error) {
	//str := fmt.Sprintf("%s", r.Url)
	//mylog.LogInfo("fetching", str)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body, r.Url)
	return parseResult, nil
}
