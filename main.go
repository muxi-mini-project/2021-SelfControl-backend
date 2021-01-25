package main

import (
	"SC/model"
	"SC/routers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

func main() {
	model.DB, err = gorm.Open("mysql", "tao:12345678@(127.0.0.1:3306)/Self_Control?parseTime=True")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%T", model.DB)
	//自动建表 迁移？
	//model.DB.AutoMigrate(&model.User{})
	r := gin.Default()
	routers.Router(r)
	r.Run(":2333")
	defer model.DB.Close()
}
