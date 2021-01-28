package user

import (
	"SC/model"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Login 登录
func Login(c *gin.Context) {
	var p model.User
	if err := c.BindJSON(&p); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	_, err := model.GetUserInfoFormOne(p.StudentID, p.Password)
	if err != nil {
		//c.Abort()
		c.JSON(400, "用户名或密码错误")
		return
	}
	if resu := model.DB.Where("student_id = ?", p.StudentID).First(&p); resu.Error != nil {
		//log.Printf(",,,,,,,\n")
		p.Gold = 0
		p.Name = "小樨"
		p.Privacy = true
		p.UserPicture = "www.baidu.com"
		model.DB.Create(&p)
	}
	//log.Printf(",%v======\n", p)
	//增加拥有默认背景
	var usersBackdrop model.UsersBackdrop
	usersBackdrop.BackdropID = 1
	usersBackdrop.StudentID = p.StudentID
	model.DB.Create(&usersBackdrop)

	claims := &model.Jwt{StudentID: p.StudentID}

	claims.ExpiresAt = time.Now().Add(200 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	var Secret = "vinegar" //加醋

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{"token": signedToken})
}
