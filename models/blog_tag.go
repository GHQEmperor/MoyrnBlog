package models

type BlogTag struct {
	BlogID uint  `gorm:"primary_key;AUTO_INCREMENT:false" json:"blog_id"`
	Blog   *Blog `json:"blog"`

	TagID uint `gorm:"primary_key;AUTO_INCREMENT:false" json:"tag_id"`
	Tag   *Tag `json:"tag"`
}
