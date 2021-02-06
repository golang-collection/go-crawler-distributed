package mongoDB

import (
	"context"
	"go-crawler-distributed/pkg/setting"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-02 11:46
* @Description:
**/

func NewMongoDBEngine(mongoDbSetting *setting.MongoDBSettingS) (*mongo.Client, error) {
	var client *mongo.Client
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(mongoDbSetting.Timeout)*time.Second)
	opt := options.Client().ApplyURI(mongoDbSetting.Url)
	opt.SetMaxPoolSize(mongoDbSetting.MaxPoolSize)
	if client, err := mongo.Connect(ctx, opt); err != nil {
		return nil, err
	} else {
		ctx2, _ := context.WithTimeout(context.Background(), time.Duration(mongoDbSetting.Timeout)*time.Second)
		err := client.Ping(ctx2, readpref.Primary())
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}
