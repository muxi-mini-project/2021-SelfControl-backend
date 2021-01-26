package user

import (
	"SC/model"
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
	_, err := model.GetUserInfoFormOne(p.StudentID, p.Password)
	if err != nil {
		//c.Abort()
		c.JSON(400, "登录失败")
		return
	}
	if ok := model.DB.NewRecord(&p); ok {
		p.Gold = 0
		p.Name = "小樨"
		p.Privacy = true
		p.UserPicture = "www.baidu.com"
		result := model.DB.Create(&p)
		if result.Error != nil {
			c.JSON(400, "登录失败")
			panic(result.Error)
		}
	}
	claims := &model.Jwt{StudentID: p.StudentID}

	claims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	var Secret = "vinegar" //加醋

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{"token": signedToken})
}
