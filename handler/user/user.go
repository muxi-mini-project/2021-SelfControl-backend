package user

import (
	"SC/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Homepage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	u := model.GetUserInfo(id)
	UserHomePage := model.UserHomePage{Name: u.Name, UserPicture: u.UserPicture}
	c.JSON(200, UserHomePage)
}

func ChangeUserInfo(c *gin.Context) {
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
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	u := model.GetUserInfo(id)
	Gold := model.Gold{Gold: u.Gold}
	c.JSON(200, Gold)
}

func GoldHistory(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	histories := model.GetGoldHistory(id)
	c.JSON(200, histories)
}

/*
func Achievement(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	achievements := model.GetAchievement(id)
	Achievement := model.Achievements{Achievements: achievements}
	c.JSON(200, Achievement)
}*/

func PunchAndNumber(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	punchs := model.GetPunchAndNumber(id)
	c.JSON(200, punchs)
}

func GetPrivacy(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	u := model.GetUserInfo(id)
	privacy := model.Privacy{Privacy: u.Privacy}
	c.JSON(200, privacy)
}

/*
func ChangePrivacy(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	var privacy bool
	if err := c.BindJSON(&privacy); err != nil {
		c.JSON(400, gin.H{"message": "修改失败"})
		return
	}
	u:= model.GetUserInfo(id)
	u.Privacy=privacy
	user:=model.UpdateUserInfo(u)
	Privacy:=model.Privacy{Privacy: user.Privacy}
	c.JSON(200,Privacy)
}*/
