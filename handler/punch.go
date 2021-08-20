package handler

import (
	"SC/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// #@Summary  当前类型所有打卡
// #@Tags punch
// #@Description 在url末尾获取类型id（1：健康 2：运动 3：学习）
// Accept application/json
// #@Produce application/json
// #@Param type_id path int true "type_id"
// #@Success 200 {object} []model.Punch2 "获取成功"
// Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// #@Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// #@Router /punchs/{type_id} [get]
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
// @Success 200 {object} []model.Punch "{"msg":"获取成功"}"
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

	var Punchs []model.UsersPunch
	model.DB.Where("student_id = ?", id).Find(&Punchs)
	if len(Punchs) == 0 {
		c.JSON(200, Punchs)
		return
	}
	punchs := model.GetMyPunch(id)
	// if len(punchs) == 0 {
	// 	punchs = punchs[:0]
	// }
	SendResponse(c, "获取成功", punchs)
}

// @Summary  判断今天是否已打卡
// @Tags punch
// @Description 在url末尾获取打卡的id
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title_id path int true "title_id"
// @Success 200 {object} Response "{"msg":"获取成功"}"
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

	TitleID, _ := strconv.Atoi(c.Param("title_id"))
	choice := model.TodayPunch(id, TitleID)
	SendResponse(c, "获取成功", choice)
}

// @Summary  判断今天是否已完成全部打卡
// @Tags punch
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} Response "{"msg":"未完成"}/{"msg":"未选择打卡"}/{"msg":"已全部完成且数量为返回的值"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/todayall [get]
func TodayPunchs(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	num := model.TodayPunchs(id)
	switch num {
	case -1:
		SendResponse(c, "未完成", -1)
	case 0:
		SendResponse(c, "未选择打卡", 0)
	default:
		SendResponse(c, "已全部完成且数量为返回的值", num)
	}

}

// @Summary  完成打卡
// @Tags punch
// @Description 完成该用户今天的该打卡
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title body model.Title true "卡的Title"
// @Success 200 {object} Response "{"msg":"打卡成功"}"
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

	var title model.Title
	if err := c.BindJSON(&title); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if title.Title == "" {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	if err := model.CompletePunch(id, title.Title); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}
	SendResponse(c, "打卡成功", nil)
}

// @Summary 获取用户某天的打卡
// @Tags punch
// Description 获取我的打卡（标签）
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param day path int true "day"
// @Success 200 {object} []model.Punch "{"msg":"1"}/{"msg":"0"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/day/{day} [get]
func GetDayPunchs(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}
	a := c.Param("day")
	x, _ := strconv.Atoi(a)
	if x == 0 {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	punchs := model.GetDayPunchs(id, x)
	histories := model.GetGoldHistory(id)

	for _, history := range histories {
		Time := []byte(history.Time)
		Time = Time[8:10]
		Time2, _ := strconv.Atoi(string(Time))
		if Time2 == x {
			SendResponse(c, "1", punchs)
			return
		}
	}
	SendResponse(c, "0", punchs)
}

// @Summary 获取用户月报的周数据
// @Tags punch
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param month path int true "month"
// @Success 200 {object} []model.WeekPunch "{"msg":"打卡成功"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/week/{month} [get]
func GetWeekPunchs(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	a := c.Param("month")
	x, _ := strconv.Atoi(a)
	if x == 0 {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	nums := model.GetWeekPunchs(id, x)
	var weekPunch []model.WeekPunch
	for i, num := range nums {
		var WeekPunch model.WeekPunch
		WeekPunch.Week = i + 1
		WeekPunch.Number = num
		weekPunch = append(weekPunch, WeekPunch)
	}
	SendResponse(c, "获取成功", weekPunch)
}

// @Summary  增加标签
// @Tags punch
// @Description 该用户新增一个打卡任务
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title body model.Title true "title"
// @Success 200 {object} Response "{"msg":"新增标签成功"}"
// @Failure 203 "该标签已选择" or "今日已完成全部打卡，不能再新增标签"
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
	if title.Title == "" {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	num := model.TodayPunchs(id)

	if num > 0 {
		c.JSON(203, gin.H{"message": "今日已完成全部打卡，不能再新增标签"})
		return
	}
	if message, err := model.CreatePunch(id, title.Title); message != "" {
		c.JSON(203, gin.H{"message": message})
		return
	} else if err != nil {
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}
	SendResponse(c, "新增标签成功", nil)
}

// @Summary  删除标签
// @Tags punch
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param title body model.Title true "需要删除的打卡title"
// @Success 200 {object} Response "{"msg":"删除成功"}"
// @Failure 203 "删除失败,用户未选择该标签"
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

	var punch model.Title
	if err := c.BindJSON(&punch); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	if punch.Title == "" {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	if s, err := model.DeletePunch(id, punch.Title); s != "" {
		c.JSON(203, gin.H{"message": "删除失败,用户未选择该标签"})
		return
	} else if err != nil {
		c.JSON(400, gin.H{"message": "Fail."})
		return
	}
	SendResponse(c, "删除成功", nil)
}

// @Summary  获取某用户标签
// @Tags punch
// @Accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Success 200 {object} []model.Punch "{"msg":"获取成功"}"
// @Failure 203 "获取失败,用户未公开标签"
// @Failure 203 "未找到该用户"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/punch/{id} [get]
func GetPunchs(c *gin.Context) {
	id := c.Param("id")
	u, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(203, gin.H{"message": "未找到该用户"})
		return
	}

	if u.Privacy == 2 {
		c.JSON(203, gin.H{"message": "获取失败,用户未公开标签"})
		return
	}
	punchs := model.GetPunchAndNumber(id)
	SendResponse(c, "获取成功", punchs)
}

// @Summary  获取某用户月报
// @Tags punch
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Punch "{"msg":"获取成功"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /punch/month [get]
func Monthly(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	punch := model.GetMonthly(id)
	SendResponse(c, "获取成功", punch)

}
