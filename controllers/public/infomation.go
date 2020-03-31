package public

import (
	"MoyrnBlog/conf"
	"MoyrnBlog/ctx"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Information(c *ctx.Context) {
	fmt.Println("information:", c.Request.Host)
	c.JSON(200, gin.H{
		"status":           10000,
		"message":          "success",
		"title":            conf.Conf.Get("title"),
		"record_number":    conf.Conf.Get("record_number"),
		"onetime_password": conf.Conf.Get("onetimePassword"),
	})
}
