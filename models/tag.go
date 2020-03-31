package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name  string `gorm:"unique;size:128" json:"name"`
	Blogs []Blog `gorm:"many2many:blog_tag" json:"blogs"`
}
