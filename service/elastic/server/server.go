package server

import (
	"context"
	"go-crawler-distributed/pkg/elastic"
	"go-crawler-distributed/pkg/util"
	"go-crawler-distributed/service/elastic/proto"
)

/**
* @Author: super
* @Date: 2020-09-01 21:33
* @Description:
**/

type Elastic struct {
}

func (e *Elastic) IndexExist(ctx context.Context, req *proto.IndexExistRequest, res *proto.IndexExistResponse) error {
	exist, err := elastic.IndexExist(req.Index)
	if err != nil {
		return err
	}
	res.Exist = exist
	return nil
}

func (e *Elastic) SaveInfo(ctx context.Context, req *proto.SaveInfoRequest, res *proto.SaveInfoResponse) error {

	article := util.ProtoToArticle(req.Article)

	id, err := elastic.SaveInfo(req.Table, article)
	if err != nil {
		return err
	}
	res.Result = id
	return nil
}

func (e *Elastic) GetInfo(ctx context.Context, req *proto.GetInfoRequest, res *proto.GetInfoResponse) error {
	article, err := elastic.GetInfo(req.Table, req.Id)
	if err != nil {
		return err
	}
	result := util.ArticleToProto(article)
	res.Article = result
	return nil
}

func (e *Elastic) SearchInfo(ctx context.Context, req *proto.SearchInfoRequest, res *proto.SearchInfoResponse) error {
	articles, err := elastic.SearchInfo(req.Table, req.FieldName, req.FieldValue)
	if err != nil {
		return err
	}
	l := len(articles)
	result := make([]*proto.Article, l)

	for i := 0; i < l; i++ {
		temp := util.ArticleToProto(articles[i])
		result = append(result, temp)
	}
	res.Article = result
	return nil
}
