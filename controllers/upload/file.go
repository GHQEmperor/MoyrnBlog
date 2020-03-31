package upload

import (
	"MoyrnBlog/conf"
	"MoyrnBlog/ctx"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func File(c *ctx.Context) {
	_, err := c.UserFilter()
	if err != nil {
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.Failed(10001, "接收文件失败")
		return
	}

	splits := strings.Split(fileHeader.Filename, ".")
	var after string
	length := len(splits)
	if length > 1 {
		after = "." + splits[length-1]
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.Failed(10002, "接收文件失败")
		return
	}
	defer file.Close()

	obj := md5.New()
	_, err = io.Copy(obj, file)
	if err != nil {
		c.Failed(10003, "接收失败")
		return
	}
	obj.Write([]byte(strconv.Itoa(int(time.Now().UnixNano()))))
	filename := hex.EncodeToString(obj.Sum(nil))
	create, err := os.Create("./static/files/" + filename + after)
	if err != nil {
		c.Failed(10004, "创建文件失败")
		return
	}

	_, _ = file.Seek(0, 0)
	_, err = io.Copy(create, file)
	if err != nil {
		c.Failed(10005, "保存失败")
		return
	}

	c.Success([]interface{}{"url", conf.Conf.Get("serve_name") + "static/files/" + filename + after})
}
