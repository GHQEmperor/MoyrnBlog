package public

import (
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

// todo: Blogs
func Blogs(c *ctx.Context) {
	var err error
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	var offset, limit int
	if offset, err = strconv.Atoi(offsetStr); err != nil {
		offset = 0
	}
	if limit, err = strconv.Atoi(limitStr); err != nil {
		limit = 10
	}

	tagID, err := strconv.Atoi(c.Query("tag_id"))
	if err != nil || tagID == 0 {
		var total int
		db.DB.Model(&models.Blog{}).Count(&total)

		if offset >= total {
			c.Failed(10002, "已经是最后一页了")
			return
		}
		var blogs []models.Blog

		db.DB.Limit(limit).Offset(offset).Order("id desc").Preload("Tags").Find(&blogs)

		c.Success([]interface{}{"blogs", blogs},
			[]interface{}{"count", total})
		return
	}

	var blogTags []models.BlogTag
	qs := db.DB.Model(&models.BlogTag{}).Where("tag_id = ?", tagID)
	var total int
	qs.Count(&total)
	if offset >= total {
		c.Failed(10002, "已经是最后一页了")
		return
	}

	qs.Preload("Blog", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Tags")
	}).Offset(offset).Limit(limit).Order("blog_id desc").Find(&blogTags)

	blogs := make([]models.Blog, 0, len(blogTags))
	for i := range blogTags {
		if blogTags[i].Blog != nil {
			blogs = append(blogs, *blogTags[i].Blog)
		}
	}

	c.Success([]interface{}{"blogs", blogs},
		[]interface{}{"count", total})
	return
}
