package main

import (
	"MoyrnBlog/UserKey/conf"
	_ "MoyrnBlog/UserKey/conf"
	"MoyrnBlog/tools"
	"os"

	"fmt"
	"time"
)

func main() {
	var count int
	fmt.Printf("\n\n")
	for {
		time.Sleep(time.Second)
		key, rest := tools.UserKey(conf.Conf.Get("user_key"))
		_, _ = fmt.Fprintf(os.Stdout, "\t\t动态密钥: %s \t 剩余时间：%d秒\r", key, rest)
		count++
	}
}
