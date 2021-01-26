package handler

import (
	"SC/model"

	"github.com/gin-gonic/gin"
)

func Punchs(c *gin.Context) {
	Type := c.Param("type")
	punchs := model.GetPunchs(Type)
	c.JSON(200, punchs)
}

func MyPunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	punchs := model.GetMyPunch(id)
	c.JSON(200, punchs)
}

func TodayPunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	title := c.Request.Header.Get("title")
	choice := model.TodayPunch(id, title)
	c.JSON(200, choice)
}

func DeletePunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	var punch model.Punch2
	c.BindJSON(&punch)
	if err := model.DeletePunch(id, punch.Title); err != nil {
		c.JSON(400, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

func CompletePunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	title := c.Request.Header.Get("title")

	if err := model.CompletePunch(id, title); err != nil {
		c.JSON(400, gin.H{"message": "打卡失败"})
		return
	}
	c.JSON(200, gin.H{"message": "打卡成功"})
}
