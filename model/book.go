package model

/**
* @Author: super
* @Date: 2020-08-14 15:29
* @Description:
**/
type Book struct {
	BookID       int     `gorm:"column:book_id" gorm:"PRIMARY_KEY" json:"book_id"`
	Title        string  `gorm:"column:title" json:"title"`
	SubTitle     string  `gorm:"column:sub_title" json:"sub_title"`
	Img          string  `gorm:"column:img" json:"img"`
	Author       string  `gorm:"column:author" json:"author"`
	Publish      string  `gorm:"column:publish" json:"publish"`
	Producer     string  `gorm:"column:producer" json:"producer"`
	PublishYear  string  `gorm:"column:publish_year" gorm:"type:date" json:"publish_year"`
	Pages        int     `gorm:"column:pages" json:"pages"`
	Price        float64 `gorm:"column:price" json:"price"`
	Layout       string  `gorm:"column:layout" json:"layout"`
	Series       string  `gorm:"column:series" json:"series"`
	ISBN         string  `gorm:"column:isbn" json:"isbn"`
	Score        float64 `gorm:"column:score" json:"score"`
	OriginalName string  `gorm:"column:original_name" json:"original_name"`
	Comments     int     `gorm:"column:comments" json:"comments"`
	CommentUrl   string  `gorm:"column:comment_url" json:"comment_url"`
}
