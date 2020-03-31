package public

import (
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
)

func Tags(c *ctx.Context) {
	var tags []models.Tag
	db.DB.Limit(5).Order("id desc").Find(&tags)
	c.Success([]interface{}{"tags", tags})
}
