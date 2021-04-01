package handler

import (
	"SC/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary  当前类型所有打卡
// @Tags punch
// @Description 在url末尾获取类型id（1：健康 2：运动 3：学习）
// Accept application/json
// @Produce application/json
// @Param type_id path int true "type_id"
// @Success 200 {object} []model.Punch2 "获取成功"
// Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punchs/{type_id} [get]
func Punchs(c *gin.Context) {
	TypeID := c.Param("type_id")
	punchs := model.GetPunchs(TypeID)
	c.JSON(200, punchs)
}

// @Summary  我的打卡
// @Tags punch
// @Description 获取我的打卡（标签）
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Punch "获取成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch [get]
func MyPunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	punchs := model.GetMyPunch(id)
	c.JSON(200, punchs)
}

// @Summary  判断今天是否已打卡
// @Tags punch
// @Description 在url末尾获取打卡的id
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title_id path int true "title_id"
// @Success 200 {object} model.Choice "获取成功（未打卡为false）"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/today/{title_id} [get]
func TodayPunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	TitleID, err := strconv.Atoi(c.Param("title_id"))
	choice := model.TodayPunch(id, TitleID)
	c.JSON(200, choice)
}

// @Summary  完成打卡
// @Tags punch
// @Description 完成该用户今天的该打卡
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title_and_gold body model.TitleAndGold true "卡的Title和完成本次打卡得到的金币数"
// @Success 200 "打卡成功"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch [post]
func CompletePunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	var a model.TitleAndGold
	//title := c.Param("title")
	//Gold := c.Param("gold")
	//Gold := c.Request.Header.Get("gold")
	//gold, _ := strconv.Atoi(Gold)
	//title := c.Request.Header.Get("title")
	if err := c.BindJSON(&a); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if err := model.CompletePunch(id, a.Title, a.Gold); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}
	c.JSON(200, gin.H{"message": "打卡成功"})
}

// @Summary  增加标签
// @Tags punch
// @Description 该用户新增一个打卡任务
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title body model.Title true "title"
// @Success 200 "新增标签成功"
// @Failure 204 "该标签已选择"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/create [post]
func CreatePunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	var title model.Title
	if err := c.BindJSON(&title); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	//title := c.Param("title")
	//title := c.Request.Header.Get("title")
	if err, message := model.CreatePunch(id, title.Title); message != "" {
		c.JSON(204, gin.H{"message": message})
		return
	} else if err != nil {
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}
	c.JSON(200, gin.H{"message": "新增标签成功"})
}

// @Summary  金币历史
// @Tags punch
// @Description 获取金币历史
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title body model.Punch2 true "需要删除的打卡title"
// @Success 200 "删除成功"
// @Failure 204 "删除失败,用户未选择该标签"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch [delete]
func DeletePunch(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	var punch model.Punch2
	c.BindJSON(&punch)
	if s, err := model.DeletePunch(id, punch.Title); s != "" {
		c.JSON(204, gin.H{"message": "删除失败,用户未选择该标签"})
		return
	} else if err != nil {
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

// @Summary  获取某用户标签
// @Tags punch
// @Accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Success 200 {object} []model.Punch "获取成功"
// @Failure 204 "获取失败,用户未公开标签"
// @Failure 203 "未找到该用户"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/punch/{id} [get]
func GetPunchs(c *gin.Context) {
	id := c.Param("id")
	err, u := model.GetUserInfo(id)
	if err != nil {
		c.JSON(203, gin.H{"message": "未找到该用户"})
		return
	}

	if u.Privacy == 0 {
		c.JSON(204, gin.H{"message": "获取失败,用户未公开标签"})
		return
	}
	punchs := model.GetPunchAndNumber(id)
	c.JSON(200, punchs)
}
