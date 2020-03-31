package main

import (
	_ "MoyrnBlog/conf"
	"MoyrnBlog/db"
	"MoyrnBlog/models"
	"MoyrnBlog/router"
	"MoyrnBlog/tools"
	"flag"
	"fmt"
	"os"
)

var syncDB = flag.Bool("syncdb", false, "sync database.")
var register = flag.Bool("register", false, "register new account.")

func init() {
	flag.Parse()
	db.ConnectDB()
	if *syncDB {
		db.SyncDB()
		os.Exit(0)
		return
	}
	if *register {
		var username, password string
		fmt.Println("Please input your username:")
		_, _ = fmt.Scanf("%s\n", &username)
		fmt.Println("Please input your password:")
		_, _ = fmt.Scanf("%s\n", &password)

		fmt.Println("Your username is :")
		fmt.Println(username)
		fmt.Println("Your password is :")
		fmt.Println(password)

		fmt.Println("Register this account? (yes/no)")
		var yes string
		_, _ = fmt.Scanf("%s\n", &yes)
		if yes == "yes" || yes == "y" {
			userKey := tools.CreateUserKey(username, password)
			user := models.User{
				Username: username,
				Password: tools.Sha256(password),
				UserKey:  userKey,
			}
			if err := db.DB.Debug().Create(&user).Error; err != nil {
				fmt.Println("create user error", err)
				os.Exit(0)
				return
			}

			fmt.Println("Your password seed is :")
			fmt.Println(userKey)
			fmt.Println("Please remember this seed.")
		}
		os.Exit(0)
	}
}

func main() {
	r := router.New()
	router.Run(r)
}
