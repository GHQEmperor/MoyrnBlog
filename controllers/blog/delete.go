package blog

import (
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
)

func Delete(c *ctx.Context) {
	_, err := c.UserFilter()
	if err != nil {
		return
	}

	blogID, err := c.PostFormInt("blog_id")
	if err != nil {
		c.Failed(10001, "参数输入错误")
		return
	}

	if result := db.DB.Delete(&models.Blog{}, "id = ?", blogID); result.Error != nil {
		c.Failed(10002, "无此博客")
		return
	}

	c.Success()
}
