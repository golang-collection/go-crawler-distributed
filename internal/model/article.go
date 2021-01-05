package model

import "fmt"

/**
* @Author: super
* @Date: 2020-09-01 17:23
* @Description:
**/

type Article struct {
	Title   string   `json:"title"`
	Url     string   `json:"url"`
	Genres  []string `json:"genres"`
	Content string   `json:"content"`
}

func (article Article) String() string {
	return fmt.Sprintf("title: %s\n"+
		"url: %s\n"+
		"geners: %v\n"+
		"content: %s\n",
		article.Title,
		article.Url,
		article.Genres,
		article.Content)
}
