package self

import (
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
)

// todo: SetAvatar
func SetAvatar(c *ctx.Context) {
	user, err := c.UserFilter()
	if err != nil {
		return
	}

	url := c.PostForm("url")

	if db.DB.Model(&user).Where("id = ?", user.ID).
		Update("avatar", url).RowsAffected == 0 {
		c.Failed(10001, "更新头像失败")
		return
	}

	c.Success()
}
