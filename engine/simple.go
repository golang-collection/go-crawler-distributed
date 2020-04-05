package engine

import (
	"fmt"
	"go-crawler-distributed/mylog"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			str := fmt.Sprintf("%v", item)
			mylog.LogInfo("parseResult.Item", str)
		}
	}
}
