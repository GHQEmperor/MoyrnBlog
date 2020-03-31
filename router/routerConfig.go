package router

import (
	"MoyrnBlog/controllers/auth"
	"MoyrnBlog/controllers/blog"
	"MoyrnBlog/controllers/public"
	"MoyrnBlog/controllers/self"
	"MoyrnBlog/controllers/upload"
)

var Routers = []Item{
	{Method: GET, Uri: "/", Function: public.Index},
	{Method: GET, Uri: "/information", Function: public.Information},
	{Method: GET, Uri: "/blogs", Function: public.Blogs},
	{Method: GET, Uri: "/tags", Function: public.Tags},
	{Method: POST, Uri: "/admin/login", Function: auth.Login},
	{Method: POST, Uri: "/admin/logout", Function: auth.Logout},
	{Method: POST, Uri: "/admin/user/info", Function: auth.UserInfo},
	{Method: POST, Uri: "/admin/upload/file", Function: upload.File},
	{Method: POST, Uri: "/admin/blog/write", Function: blog.Write},
	{Method: POST, Uri: "/admin/blog/tags", Function: blog.SearchTags},
	{Method: POST, Uri: "/admin/blog/update", Function: blog.Update},
	{Method: POST, Uri: "/admin/blog/delete", Function: blog.Delete},
	{Method: POST, Uri: "/admin/user/setAvatar", Function: self.SetAvatar},
}
