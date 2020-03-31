package blog

import (
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
)

func SearchTags(c *ctx.Context) {
	_, err := c.UserFilter()
	if err != nil {
		return
	}

	blogID, err := c.PostFormInt("blog_id")
	if err != nil {
		c.Failed(10001, "请选择博客")
		return
	}

	var blogTags []models.BlogTag
	db.DB.Where("blog_id = ?", blogID).Preload("Tag").Find(&blogTags)

	tagsName := make([]string, 0, len(blogTags))
	for i := range blogTags {
		tagsName = append(tagsName, blogTags[i].Tag.Name)
	}

	c.Success([]interface{}{"tags", tagsName})
}
