package user

import (
	"SC/model"
	"SC/token"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var p model.User
	if err := c.BindJSON(&p); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	//modle.GetUserInfoFormOne("2020213675","lst...")
	stu, err := model.GetUserInfoFormOne(p.StudentID, p.Password)
	if err != nil {
		//c.Abort()
		c.JSON(400, "登录失败")
		return
	}
	if ok := main.db.NewRecord(&p); ok {
		c.JSON(400, "该用户已注册")
		return
	}
	data := token.Jwt{
		ID:   p.StudentID,
		Name: stu.User.Name,
	}

	data.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	data.IssuedAt = time.Now().Unix()

	var Secret = "vinegar" //加醋

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{"token": signedToken})
}
