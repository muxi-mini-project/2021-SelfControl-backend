package handler

import (
	"SC/model"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary  排行榜数据
// @Description url最后面+week或month查询数据
// Accept application/json
// @Produce application/json
// @Param type path string true "type"
// @Success 200 {object} []model.UserAndNumber "获取前十用户"
// @Failure 204 "未检索到该时间段的打卡信息"
// Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /list/:type [get]
func List(c *gin.Context) {
	Type := c.Param("type")
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

// @Summary  兑换排名
// @Description 根据url末尾接收到的排名（第一名/第二名）
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param ranking body model.Ranking true "ranking"
// @Success 200 "兑换成功"
// @Failure 204 "金币不足"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /list [put]
func ChangeRanking(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	var ranking model.Ranking
	if err := c.BindJSON(&ranking); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if err, message := model.ChangeRanking(id, ranking.Ranking); message != "" {
		c.JSON(204, gin.H{"message": "金币不足"})
	} else if err != nil {
		c.JSON(400, gin.H{"message": "Fail."})
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"message": "兑换成功"})
}

// @Summary  排名兑换价格
// Tags user
// @Description 获取排名兑换价格
// @Accept application/json
// @Produce application/json
// Param token header string true "token"
// @Success 200 {object} []model.ListPrice "获取成功"
// Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /listprice [get]
func ListPrice(c *gin.Context) {
	prices := model.GetListPrice()
	c.JSON(200, prices)
}

// @Summary  背景价格
// Tags user
// @Description 获取背景价格
// Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Backdrop "获取成功"
// Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /backdrop [get]
func BackdropPrice(c *gin.Context) {
	prices := model.GetBackdropPrice()
	c.JSON(200, prices)
}

// @Summary 兑换背景
// Tags user
// @Description 根据背景id兑换背景
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param backdrop_id body model.BackdropID true "backdrop_id"
// @Success 200 "兑换成功"
// @Failure 204 "金币不足"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /backdrop [put]
func ChangeBackdrop(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	var b model.BackdropID
	c.BindJSON(&b)
	if err, message := model.ChangeBackdrop(id, b.BackdropID); message != "" {
		c.JSON(400, gin.H{"message": "金币不足"})
	} else if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}

	c.JSON(200, gin.H{"message": "兑换成功"})
}

// @Summary  我的背景
// @Tags user
// @Description 获取我的背景id
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Backdrop "获取成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /backdrops [get]
func MyBackdrops(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	backdrops := model.GetBackdrop(id)
	c.JSON(200, backdrops)
}
