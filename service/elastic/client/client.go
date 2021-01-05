package client

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"go-crawler-distributed/configs"
	"go-crawler-distributed/global"
	"go-crawler-distributed/internal/model"
	"go-crawler-distributed/pkg/util"
	"go-crawler-distributed/service/elastic/proto"
)

/**
* @Author: super
* @Date: 2021-01-05 19:34
* @Description:
**/


var elasticOP proto.ElasticOperationService

func init(){
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			configs.ConsulURL,
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.service.elastic.client"),
	)
	service.Init()
	elasticOP = proto.NewElasticOperationService("go.micro.service.elastic", service.Client())
}

func IndexExist(index string) (bool, error) {
	res, err := elasticOP.IndexExist(context.TODO(), &proto.IndexExistRequest{Index: index})
	if err != nil{
		global.Logger.Error(context.Background(), err)
		return false, err
	}
	return res.Exist, nil
}

func SaveInfo(table string, data *model.Article) (string, error){
	article := util.ArticleToProto(data)
	res, err := elasticOP.SaveInfo(context.TODO(), &proto.SaveInfoRequest{Table: table, Article: article})
	if err != nil{
		global.Logger.Error(context.Background(), err)
		return "", err
	}
	return res.Result, nil
}

func GetInfo(table string, id string) (*model.Article, error) {
	res, err := elasticOP.GetInfo(context.TODO(), &proto.GetInfoRequest{Table: table, Id: id})
	if err != nil{
		global.Logger.Error(context.Background(), err)
		return nil, err
	}
	article := util.ProtoToArticle(res.Article)
	return article, nil
}

func SearchInfo(table string, fieldName string, fieldValue string)([]*model.Article, error){
	res, err := elasticOP.SearchInfo(context.TODO(), &proto.SearchInfoRequest{Table: table, FieldName: fieldName, FieldValue: fieldValue})
	if err != nil{
		global.Logger.Error(context.Background(), err)
		return nil, err
	}
	l := len(res.Article)
	result := make([]*model.Article, l)

	for i:=0 ;i<l; i++{
		temp := util.ProtoToArticle(res.Article[i])
		result = append(result, temp)
	}
	return result, nil
}