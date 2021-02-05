package handler

import (
	"SC/model"
	"log"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	Type := c.Param("type")
	//log.Printf("%v,,,,,,,,,,,,,,\n", Type)
	numbers, message := model.List(Type)
	if message != "" {
		c.JSON(400, gin.H{"message": message})
		return
	}
	if len(numbers) > 10 {
		var numbers2 []model.UserAndNumber
		numbers2 = append(numbers2, numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5], numbers[6], numbers[7], numbers[8], numbers[9])
		numbers = numbers2
	}
	c.JSON(200, numbers)
}

func ChangeRanking(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	var ranking model.Ranking
	if err := c.BindJSON(&ranking); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	if err, message := model.ChangeRanking(id, ranking.Ranking); message != "" {
		c.JSON(400, gin.H{"message": "金币不足"})
	} else if err != nil {
		c.JSON(400, gin.H{"message": "兑换失败"})
		log.Println(err)
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

	var b model.BackdropID
	c.BindJSON(&b)
	//BackdropId := c.Param("backdrop_id")
	//BackdropId := c.Request.Header.Get("backdrop_id")
	//backdropid, _ := strconv.Atoi(BackdropId)
	if err, message := model.ChangeBackdrop(id, b.BackdropID); message != "" {
		c.JSON(400, gin.H{"message": "金币不足"})
	} else if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "兑换失败"})
		return
	}

	c.JSON(200, gin.H{"message": "兑换成功"})
}

func MyBackdrops(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	backdrops := model.GetBackdrop(id)
	c.JSON(200, backdrops)
}
