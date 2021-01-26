package handler

import (
	"SC/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	Type := c.Param("type")
	numbers := model.List(Type)
	if len(numbers) > 10 {
		var numbers2 []model.UserAndNumber
		numbers2 = append(numbers2, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		numbers = numbers2
	}
	Numbers := model.UserAndNumbers{UserAndNumber: numbers}
	c.JSON(200, Numbers)
}

func ChangeRanking(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	title := c.Request.Header.Get("title")
	if err := model.ChangeRanking(id, title); err != nil {
		c.JSON(400, gin.H{"message": "兑换失败"})
		return
	}
	c.JSON(200, gin.H{"message": "兑换成功"})
}

func ListPrice(c *gin.Context) {
	prices := model.GetListPrice()
	c.JSON(200, prices)
}

func BackdropPrice(c *gin.Context) {
	prices := model.GetBackdropPrice()
	c.JSON(200, prices)
}

func ChangeBackdrop(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}
	BackdropId := c.Request.Header.Get("backdrop_id")
	backdropid, _ := strconv.Atoi(BackdropId)
	if err := model.ChangeBackdrop(id, backdropid); err != nil {
		c.JSON(400, gin.H{"message": "兑换失败"})
		return
	}
	c.JSON(200, gin.H{"message": "兑换成功"})
}

func MyBackdrops(c *gin.Context) {

}
