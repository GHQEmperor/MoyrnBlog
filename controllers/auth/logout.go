package auth

import (
	"MoyrnBlog/ctx"
	"github.com/gin-gonic/contrib/sessions"
)

func Logout(c *ctx.Context) {
	session := sessions.Default(c.Context)
	session.Delete("user")
	if err := session.Save(); err != nil {
		c.Failed(10001, "未知错误")
		return
	}

	c.Success()
}
