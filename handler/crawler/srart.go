package crawler

import (
	"github.com/gin-gonic/gin"
	"go-crawler-distributed/crawler"
	"go-crawler-distributed/crawler/crawerConfig"
	p1 "go-crawler-distributed/crawler/douban/parser"
	p2 "go-crawler-distributed/crawler/meituan/parser"
	"go-crawler-distributed/handler"
)

/**
* @Author: super
* @Date: 2020-09-03 20:54
* @Description:
**/

// @Summary start new crawler of douban
// @Description start new crawler of douban
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /start/douban [Post]
func StartDoubanCrawler(c *gin.Context){

	go crawler.Crawl("", crawerConfig.TagUrl, "tags", p1.ParseTagList)
	handler.SendResponse(c, nil, nil)
}

// @Summary start new crawler of meituan
// @Description start new crawler of meituan
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /start/meituan [Post]
func StartMeituanCrawler(c *gin.Context){

	go crawler.Crawl("", crawerConfig.ArticleList, "ArticleList", p2.ParseArticleList)
	handler.SendResponse(c, nil, nil)
}