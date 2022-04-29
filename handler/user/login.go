package user

import (
	"SC/handler"
	"SC/model"
	"SC/service/user"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// @Summary  登录
// @Tags user
// @Description 学号密码登录
// @Accept application/json
// @Produce application/json
// @Param object body model.User true "登录的用户信息"
// Success 200 {object} Token "将student_id作为token保留"
// @Success 200 {object} handler.Response "{"msg":"将data保留，并作为token放入后续请求header"}"
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
	if p.StudentID == "" {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	pwd := p.Password
	// 首次登录 验证一站式
	if resu := model.DB.Where("student_id = ?", p.StudentID).First(&p); resu.Error != nil {
		_, err := model.GetUserInfoFormOne(p.StudentID, pwd)
		if err != nil {
			// c.Abort()
			c.JSON(401, "Password or account wrong.")
			return
		}
		p.CurrentBackdrop = 1
		p.Gold = 0
		p.Name = "小樨"
		p.Privacy = 1
		p.UserPicture = "www.baidu.com"

		p.Password = model.GeneratePasswordHash(pwd) // 加密后存数据库
		// p.Password = base64.StdEncoding.EncodeToString([]byte(p.Password)) // 加密后存数据库

		model.DB.Create(&p)
		// 增加拥有默认背景
		// var usersBackdrop model.UsersBackdrop
		// usersBackdrop.BackdropID = 0
		// usersBackdrop.StudentID = p.StudentID
		// model.DB.Create(&usersBackdrop)
	} else {
		// password, _ := base64.StdEncoding.DecodeString(p.Password) // 从数据库中解密后比较

		if !model.CheckPassword(pwd, p.Password) {
			c.JSON(401, "Password or account wrong.")
			return
		}
	}

	claims := &user.Jwt{StudentID: p.StudentID}

	claims.ExpiresAt = time.Now().Add(12 * 30 * 24 * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()

	var Secret = "vinegar" // 加醋

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Println(err)
	}

	handler.SendResponse(c, "将data保留，并作为token放入后续请求header", signedToken)
}
