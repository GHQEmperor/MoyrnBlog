package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:128;unique" json:"username"`
	Password string `json:"-"`

	// 令牌
	UserKey string `gorm:"type:text" json:"-"`

	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`

	// 个人简介
	Remark string `gorm:"type:text" json:"remark"`
}
