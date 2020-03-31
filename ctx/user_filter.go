package ctx

import (
	"MoyrnBlog/db"
	"MoyrnBlog/models"
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"log"
)

var (
	ErrNotLogin = errors.New("not login")
)

func (c *Context) UserFilter() (*models.User, error) {
	session := sessions.Default(c.Context)
	userInterface := session.Get("user")
	user, ok := userInterface.(models.User)
	if !ok {
		c.AbortWithStatus(403)
		log.Printf("UserFilter error:%s\n", ErrNotLogin)
		return nil, ErrNotLogin
	}
	var userNew models.User
	db.DB.Where("id = ?", user.ID).First(&userNew)
	sessionNew := sessions.Default(c.Context)
	sessionNew.Set("user", userNew)
	if err := sessionNew.Save(); err != nil {
		c.Failed(10001, "未知错误")
		return nil, err
	}
	return &userNew, nil
}
