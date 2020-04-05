package persist

import (
	"fmt"
	"github.com/olivere/elastic"
	"go-crawler-distributed/engine"
	"go-crawler-distributed/mylog"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := Save(s.Client, item, s.Index)
	if err == nil {
		str := fmt.Sprintf("%v", item)
		mylog.LogInfo("rpc.Item %v saved", str)
		*result = "ok"
	} else {
		mylog.LogError("rpc.Save", err)
	}
	return err
}
