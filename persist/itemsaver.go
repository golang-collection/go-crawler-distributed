package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"go-crawler-distributed/engine"
	"go-crawler-distributed/mylog"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			str := fmt.Sprintf("%d %v", itemCount, item)
			mylog.LogInfo("ItemSaver", str)
			itemCount++

			err := Save(client, item, index)
			if err != nil {
				mylog.LogError("item saver error", err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
