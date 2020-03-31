package ctx

import "github.com/gin-gonic/gin"

func (c *Context) Success(other ...[]interface{}) {
	mp := make(map[string]interface{})
	mp["status"] = 10000
	mp["message"] = "success"

	for _, v := range other {
		mp[v[0].(string)] = v[1]
	}
	c.JSON(200, mp)
}

func (c *Context) Failed(status int, message string) {
	c.JSON(200, gin.H{
		"status":  status,
		"message": message,
	})
}
