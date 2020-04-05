package worker

import (
	"go-crawler-distributed/engine"
	"go-crawler-distributed/mylog"
)

type CrawlService struct{}

func (CrawlService) Process(
	req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		mylog.LogError("worker.Process1", err)
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		mylog.LogError("worker.Process2", err)
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
