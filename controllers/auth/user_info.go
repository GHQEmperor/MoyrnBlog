package auth

import "MoyrnBlog/ctx"

// todo: UserInfo
func UserInfo(c *ctx.Context) {
	user, err := c.UserFilter()
	if err != nil {
		return
	}

	c.Success([]interface{}{"user", user})
}
