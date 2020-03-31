package db

import (
	"MoyrnBlog/conf"
	"MoyrnBlog/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/url"
	"os"
	"time"
)

var DB *gorm.DB

func ConnectDB() {
	dbHost := conf.Conf.Get("db_host")
	dbPort := conf.Conf.Get("db_port")
	dbUser := conf.Conf.Get("db_user")
	dbPass := conf.Conf.Get("db_pass")
	dbName := conf.Conf.Get("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=%s&parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		url.QueryEscape("Asia/Shanghai"),
	)

	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("database connect error:", err)
		os.Exit(0)
	}

	DB.SingularTable(true)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetConnMaxLifetime(time.Minute)

	if conf.Conf.Get("debug") == "true" {
		DB = DB.Debug()
	}

	fmt.Println("database connect success")
}

func SyncDB() {
	ConnectDB()
	DB.Debug().Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.User{},
			&models.Blog{},
			&models.Tag{},
			&models.BlogTag{},
		)

	os.Exit(0)
}
