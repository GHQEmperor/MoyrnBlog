package auth

import (
	"MoyrnBlog/conf"
	"MoyrnBlog/ctx"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
	"MoyrnBlog/tools"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
)

func Login(c *ctx.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	onetimePassword := c.PostForm("onetime_password")
	if username == "" || password == "" {
		c.Failed(10001, "输入不能为空")
		return
	}

	var user models.User
	res := db.DB.Where("username = ? and password = ?", username, tools.Sha256(password)).
		First(&user)
	if res.Error != nil {
		c.Failed(10001, "登录失败")
		return
	}

	serverKey, _ := tools.UserKey(user.UserKey)
	if conf.Conf.Get("onetimePassword") == "true" && onetimePassword != serverKey {
		fmt.Println(onetimePassword, serverKey)
		c.Failed(10001, "登录失败")
		return
	}

	session := sessions.Default(c.Context)
	session.Set("user", user)
	if err := session.Save(); err != nil {
		fmt.Println("session save error:", err)
		c.Failed(10002, "登录错误")
		return
	}

	c.Success()
}
