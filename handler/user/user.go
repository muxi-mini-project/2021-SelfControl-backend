package user

import (
	"SC/handler"
	"SC/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary  用户信息
// @Tags user
// @Description 获取用户信息
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// Success 200 {object} model.User "获取成功"
// @Success 200 {object} model.User "{"msg":"获取成功"}"
// @Failure 203 {object} error.Error "{"error_code":"20001", "message":"Fail."}
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user [get]
func UserInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	u, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(203, gin.H{"message": "Fail."})
		return
	}
	handler.SendResponse(c, "获取成功", u)
}

// @Summary  修改用户信息
// @Tags user
// @Description 接收新的User结构体来修改用户信息
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param User body model.User true "需要修改的用户信息"
// Success 200 "修改成功"
// @Success 200 {object} handler.Response "{"msg":"修改成功"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user [put]
func ChangeUserInfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}

	if user.Privacy != 2 && user.Privacy != 1 && user.Privacy != 0 {
		c.JSON(400, gin.H{"message": "Privacy参数错误(2 = 公开， 1 = 不公开)"})
		return
	}
	user.StudentID = id
	if err := model.UpdateUserInfo(user); err != nil {
		c.JSON(400, gin.H{"message": "更新失败"})
		return
	}
	handler.SendResponse(c, "修改成功", nil)
}

// @Summary  金币历史
// @Tags user
// @Description 获取金币历史
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.GoldHistory "{"msg":"获取成功"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/goldhistory [get]
func GoldHistory(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	histories := model.GetGoldHistory(id)
	if len(histories) > 3 {
		histories = histories[len(histories)-3:]
	}

	handler.SendResponse(c, "获取成功", histories)
}

// @Summary  我的打卡数
// @Tags user
// @Description 获取我的打卡数
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 {object} []model.Punch "{"msg":"获取成功"}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/punch [get]
func PunchAndNumber(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "Token Invalid."})
		return
	}

	punchs := model.GetPunchAndNumber(id)

	handler.SendResponse(c, "获取成功", punchs)
}

// @Summary  隐私
// @Tags user
// @Description 判断该用户是否选择公开自己的打卡标签
// @Accept application/json
// @Produce application/json
// @Param id path int true "id"
// @Success 200 {object} handler.Response "{"msg":"默认为1 若要修改隐私 直接使用修改用户信息"}"
// @Failure 203 "未找到该用户"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录"
// Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/privacy/{id} [get]
func GetPrivacy(c *gin.Context) {
	id := c.Param("id")
	u, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(203, gin.H{"message": "未找到该用户"})
		return
	}

	handler.SendResponse(c, "默认为1 若要修改隐私 直接使用修改用户信息", u.Privacy)
}
