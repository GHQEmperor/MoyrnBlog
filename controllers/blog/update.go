package blog

import (
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
)

func Update(c *ctx.Context) {
	_, err := c.UserFilter()
	if err != nil {
		return
	}

	id, err := c.PostFormInt("ID")
	if err != nil {
		c.Failed(10001, "参数不符合规范")
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")

	if result := db.DB.Model(&models.Blog{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"title":   title,
			"content": content,
		}); result.RowsAffected == 0 {
		c.Failed(10002, "修改失败")
		return
	}
	c.Success()
}
