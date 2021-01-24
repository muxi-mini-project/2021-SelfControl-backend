package user

import (
	"SC/model"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	name := model.GetName(id)
	PictureUrl := model.GetPicture(id)
	UserHomePage := model.User{Name: name, UserPicture: PictureUrl}
	c.JSON(200, UserHomePage)
}

func ChangeInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	user.StudentID = id
	if err := model.UpdateUserInfo(user); err != nil {
		c.JSON(400, gin.H{"message": "更新失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}

func Gold(c *gin.Context) {

}

func Achievement(c *gin.Context) {

}

func PunchNumber(c *gin.Context) {

}

func Privary(c *gin.Context) {

}
