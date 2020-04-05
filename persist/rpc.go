package persist

import (
	"github.com/olivere/elastic"
	"go-crawler-distributed/engine"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error{
	err := Save(s.Client, item, s.Index)
	if err == nil{
		*result = "ok"
	}
	return err
}