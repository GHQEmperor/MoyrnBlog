package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	User   *User  `json:"user"`
	Title  string `json:"title"`
	//Cover string `gorm:"type:text" json:"cover"`

	Content string `gorm:"type:text" json:"content"`
	Tags    []Tag  `gorm:"many2many:blog_tag" json:"tags"`
}
