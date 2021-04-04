package handler

import (
	"SC/model"
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary  排行榜数据
// @Description url最后面+week或month查询数据
// @Tags List
// Accept application/json
// @Produce application/json
// @Param type path string true "type"
// @Success 200 {object} []model.UserRanking "获取前五用户"
// @Failure 203 "未检索到该时间段的打卡信息"
// Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /lists/{type} [get]
func List(c *gin.Context) {
	Type := c.Param("type")
	var numbers []model.UserRanking
	var message string
	if Type == "month" {
		numbers, message = model.GetMonthList()
	} else if Type == "week" {
		numbers, message = model.GetWeekList()
	} else {
		c.JSON(400, gin.H{"message": "url最后面+ week或month查询数据"})
		return
	}
	if message != "" {
		c.JSON(400, gin.H{"message": message})
		return
	}
	if len(numbers) > 5 {
		numbers = numbers[5:]
	}
	c.JSON(200, numbers)
}

// @Summary  兑换排名
// @Description 需要前进的排名
// @Tags List
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param type path string true "type"
// @Param ranking body model.Ranking true "ranking"
// @Success 200 "兑换成功"
// @Failure 203 "金币不足"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /list/{type} [put]
func ChangeRanking(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	Type := c.Param("type")
	var ranking model.Ranking
	if err := c.BindJSON(&ranking); err != nil || ranking.Ranking < 1 || ranking.Ranking > 10 {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if ranking.Ranking == 0 {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if Type == "week" {
		if err, message := model.ChangeWeekRanking(id, ranking.Ranking); message != "" {
			c.JSON(203, gin.H{"message": "金币不足"})
			return
		} else if err != nil {
			c.JSON(400, gin.H{"message": "Fail."})
			log.Println(err)
			return
		}
	} else if Type == "month" {
		if err, message := model.ChangeMonthRanking(id, ranking.Ranking); message != "" {
			c.JSON(203, gin.H{"message": "金币不足"})
			return
		} else if err != nil {
			c.JSON(400, gin.H{"message": "Fail."})
			log.Println(err)
			return
		}
	} else {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	c.JSON(200, gin.H{"message": "兑换成功"})
}

// @Summary  获取兑换排名历史
// @Tags List
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} model.ListHistories
// @Failure 203 "金币不足"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /list/history [get]
func ListHistory(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	history := model.GetListHistory(id)
	c.JSON(200, history)
}

// @Summary  背景价格
// @Tags Backdrop
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
// @Tags Backdrop
// @Description 根据背景id兑换背景
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param backdrop_id body model.BackdropID true "backdrop_id"
// @Success 200 "兑换成功"
// @Failure 203 "金币不足"
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
	if err := c.BindJSON(&b); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	if b.BackdropID == 0 {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if err, message := model.ChangeBackdrop(id, b.BackdropID); message != "" {
		c.JSON(203, gin.H{"message": "金币不足"})
		return
	} else if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}

	c.JSON(200, gin.H{"message": "兑换成功"})
}

// @Summary  我的背景
// @Tags Backdrop
// @Description 获取我的背景id
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} model.BackdropRes "获取成功"
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
	var Backdrops model.BackdropRes
	for _, value := range backdrops {
		switch value.BackdropID {
		case 1:
			Backdrops.B1 = 1
		case 2:
			Backdrops.B2 = 1
		case 3:
			Backdrops.B3 = 1
		case 4:
			Backdrops.B4 = 1
		case 5:
			Backdrops.B5 = 1
		}
	}
	c.JSON(200, Backdrops)
}

// @Summary  用户排名
// @Tags List
// @Description 根据 type 和 id 获取用户排名
// @Accept application/json
// @Produce application/json
// @Param type path string true "type"
// @Param id path string true "id"
// @Success 200 {object} model.Ranking "获取成功"
// Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /list/user/{id}/{type} [get]
func UserRanking(c *gin.Context) {
	id := c.Param("id")
	Type := c.Param("type")
	rank := model.GetUserRanking(id, Type)
	r := model.Ranking{
		Ranking: rank,
	}
	c.JSON(200, r)
}
