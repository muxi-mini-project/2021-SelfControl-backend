package handler

import (
	"SC/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Punchs(c *gin.Context) {
	TypeID := c.Param("type_id")
	//TypeID, _ := strconv.Atoi(TypeId)
	punchs := model.GetPunchs(TypeID)
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

	TitleID := c.Param("title_id")
	//title := c.Request.Header.Get("title")
	TitleId, _ := strconv.Atoi(TitleID)
	choice := model.TodayPunch(id, TitleId)
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
	if s, err := model.DeletePunch(id, punch.Title); s != "" {
		c.JSON(400, gin.H{"message": "删除失败,用户未选择该标签"})
		return
	} else if err != nil {
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

	var a model.TitleAndGold
	//title := c.Param("title")
	//Gold := c.Param("gold")
	//Gold := c.Request.Header.Get("gold")
	//gold, _ := strconv.Atoi(Gold)
	//title := c.Request.Header.Get("title")
	if err := c.BindJSON(&a); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	if err := model.CompletePunch(id, a.Title, a.Gold); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "打卡失败"})
		return
	}
	c.JSON(200, gin.H{"message": "打卡成功"})
}

func CreatePunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户信息，请先登录"})
		return
	}

	var title model.Title
	if err := c.BindJSON(&title); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	//title := c.Param("title")
	//title := c.Request.Header.Get("title")
	if err, message := model.CreatePunch(id, title.Title); message != "" {
		c.JSON(400, gin.H{"message": message})
		return
	} else if err != nil {
		c.JSON(400, gin.H{"message": "新增标签失败"})
		return
	}
	c.JSON(200, gin.H{"message": "新增标签成功"})
}
