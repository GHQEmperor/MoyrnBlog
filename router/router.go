package router

import (
	"MoyrnBlog/conf"
	"MoyrnBlog/ctx"
	"MoyrnBlog/models"
	"encoding/gob"
	"github.com/gin-gonic/contrib/sessions"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GET     = "GET"
	POST    = "POST"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	PUT     = "PUT"
	OPTIONS = "OPTIONS"
	HEAD    = "HEAD"
)

type HandleFunc func(ctx *ctx.Context)

type Item struct {
	Method   string
	Uri      string
	Function HandleFunc
}

func Handle(h HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := &ctx.Context{
			Context: c,
		}
		h(context)
	}
}

func New() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("./views/*")
	router.StaticFS("/static/", http.Dir("./static/"))

	gob.Register(models.User{})
	store, err := sessions.
		NewRedisStore(100, "tcp",
			conf.Conf.Get("redis_addr")+":"+conf.Conf.Get("redis_port"),
			conf.Conf.Get("redis_pass"),
			[]byte("secret"))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 60 * 60,
	})

	router.Use(sessions.Sessions("session", store))

	for _, v := range Routers {
		switch v.Method {
		case GET:
			router.GET(v.Uri, Handle(v.Function))
		case POST:
			router.POST(v.Uri, Handle(v.Function))
		case DELETE:
			router.DELETE(v.Uri, Handle(v.Function))
		case PATCH:
			router.PATCH(v.Uri, Handle(v.Function))
		case PUT:
			router.PUT(v.Uri, Handle(v.Function))
		case OPTIONS:
			router.OPTIONS(v.Uri, Handle(v.Function))
		case HEAD:
			router.HEAD(v.Uri, Handle(v.Function))
		default:
		}
	}
	return router
}

func Run(e *gin.Engine) {
	if err := e.Run(conf.Conf.Get("host")); err != nil {
		panic(err)
	}
}
