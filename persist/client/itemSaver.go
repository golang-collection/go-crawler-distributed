package client

import (
	"fmt"
	"go-crawler-distributed/engine"
	"go-crawler-distributed/mylog"
	"go-crawler-distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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

			//调用Rpc存储item
			result := ""
			err = client.Call("ItemSaverService.Save",
				item, &result)
			if err != nil {
				mylog.LogError("item saver error", err)
			}
		}
	}()

	return out, nil
}
