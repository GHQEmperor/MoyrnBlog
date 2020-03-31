package blog

import (
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
	"encoding/json"
)

// todo: Write
func Write(c *ctx.Context) {
	user, err := c.UserFilter()
	if err != nil {
		return
	}

	title := c.PostForm("title")
	if title == "" {
		c.Failed(10001, "标题不能为空")
		return
	}
	tagsStr := c.PostForm("tags")
	if tagsStr == "" {
		c.Failed(10002, "请选择标签")
		return
	}
	var tags []string
	if err := json.Unmarshal([]byte(tagsStr), &tags); err != nil {
		c.Failed(10999, "参数不规范")
		return
	}

	content := c.PostForm("content")
	if content == "" {
		c.Failed(10003, "内容不能为空")
		return
	}

	tx := db.DB.Begin()

	blog := models.Blog{
		UserID:  user.ID,
		Title:   title,
		Content: content,
	}
	if err := tx.Create(&blog).Error; err != nil {
		tx.Rollback()
		c.Failed(10004, "发布失败")
		return
	}

	for i := range tags {
		newTag := models.Tag{
			Name: tags[i],
		}
		if err := tx.FirstOrCreate(&newTag, "name = ?", tags[i]).
			Error; err != nil {
			tx.Rollback()
			c.Failed(10005, "未知错误")
			return
		}
		if err := tx.Create(&models.BlogTag{
			BlogID: blog.ID,
			TagID:  newTag.ID,
		}).Error; err != nil {
			tx.Rollback()
			c.Failed(10006, "关联标签失败")
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.Failed(10007, "发布失败")
		return
	}

	c.Success()
}
