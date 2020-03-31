package ctx

import "strconv"

func (c *Context) PostFormInt(key string) (int, error) {
	return strconv.Atoi(c.PostForm(key))
}
