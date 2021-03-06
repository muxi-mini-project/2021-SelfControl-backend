package user

import (
	"SC/model"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Token struct {
	Token string `json:"token"`
}

// @Summary  登录
// @Tags user
// @Description 学号密码登录
// @Accept application/json
// @Produce application/json
// @Param object body model.User true "登录的用户信息"
// @Success 200 {object} Token "将student_id作为token保留"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Password or account wrong."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user [post]
func Login(c *gin.Context) {
	var p model.User
	if err := c.BindJSON(&p); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	pwd := p.Password
	//首次登录 验证一站式
	if resu := model.DB.Where("student_id = ?", p.StudentID).First(&p); resu.Error != nil {
		_, err := model.GetUserInfoFormOne(p.StudentID, pwd)
		if err != nil {
			//c.Abort()
			c.JSON(401, "Password or account wrong.")
			return
		}
		p.Gold = 0
		p.Name = "小樨"
		p.Privacy = 1
		p.UserPicture = "www.baidu.com"
		model.DB.Create(&p)
	} else {
		if p.Password != pwd {
			c.JSON(401, "Password or account wrong.")
			return
		}
	}

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
	var Token Token
	Token.Token = signedToken
	c.JSON(200, Token)
}
