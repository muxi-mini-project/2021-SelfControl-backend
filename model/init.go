package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB 全局变量
var DB *gorm.DB

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		fmt.Errorf("Open database failed, %s\n", err.Error())
		panic(err)
	}

	return db
}

func InitDB()  {
	DB = openDB(viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.addr"), viper.GetString("db.name"))
}