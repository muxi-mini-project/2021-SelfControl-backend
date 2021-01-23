package user

import (
	"SC/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var p model.User
	c.BindJSON(&p)
	//modle.GetUserInfoFormOne("2020213675","lst...")
	user, err := model.GetUserInfoFormOne(p.StudentID, p.Password)
	if err != nil {
		c.Abort()
		c.JSON(400, "登录失败")
		return
	}
	fmt.Println(user)
}
