package main

import (
	"SC/model"
	"SC/routers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := model.DB
	db, err := gorm.Open("mysql", "tao:12345678@(127.0.0.1:13306)/Self_Control?parseTime=True")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", db)
	//自动建表 迁移？
	db.AutoMigrate(&model.User{})
	r := gin.Default()
	routers.Router(r)
	r.Run(":2333")
	defer db.Close()
}
