package main

import (
	"SC/model"
	"SC/routers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "tao:12345678@/users")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	r := gin.Default()
	routers.Router(r)
	r.Run(":2333")

}
