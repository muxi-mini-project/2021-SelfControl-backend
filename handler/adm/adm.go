package adm

import (
	"SC/handler"
	"SC/model"

	"github.com/gin-gonic/gin"
)

// @Summary  (管理员)新增金币历史
// @Description 管理员给某用户新增金币历史
// @Tags administrator
// @Accept application/json
// @Produce application/json
// @Param goldhistory body model.GoldHistory true "需要新增的金币历史:time和student_id重要,其他参数非必须 time格式 2021-08-06 19:04:05"
// @Success 200 {object} handler.Response "{"msg":"新增成功"}"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /adm/goldhistory [post]
func GoldHistory(c *gin.Context) {
	var goldhistory model.GoldHistory
	if err := c.BindJSON(&goldhistory); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	result := model.DB.Create(&goldhistory)
	if result.Error != nil {
		c.JSON(400, "Fail.")
	}

	handler.SendResponse(c, "新增成功", nil)
}

// @Summary  (管理员)新增打卡记录
// @Description 管理员给某用户新增打卡记录
// @Tags administrator
// @Accept application/json
// @Produce application/json
// @Param punch body model.PunchHistory true "需要新增的打卡记录:student_id month day为重要参数其他参数非必须 day是指这天为今年的第几天 time格式 2021-08-06 19:04:05"
// @Success 200 {object} handler.Response "{"msg":"新增成功"}"
// @Failure 400 {object} error.Error "{"error_code":"20001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /adm/punch [post]
func PunchRecord(c *gin.Context) {
	var punch model.PunchHistory
	if err := c.BindJSON(&punch); err != nil {
		c.JSON(400, gin.H{"message": "Lack Param Or Param Not Satisfiable."})
		return
	}
	result := model.DB.Create(&punch)
	if result.Error != nil {
		c.JSON(400, "Fail.")
	}

	handler.SendResponse(c, "新增成功", nil)
}

// @Summary  (管理员)清除用户背景
// @Description 管理员直接删除用户拥有的背景
// @Tags administrator
// @Accept application/json
// @Produce application/json
// @Param student_id path int true "student_id"
// @Success 200 {object} handler.Response "{"msg":"删除成功"}"
// @Router /adm/del_backdrop/{student_id} [get]
func DeleteBackdrop(c *gin.Context) {
	sid := c.Param("student_id")

	model.DB.Delete(model.UsersBackdrop{}, "student_id = ?", sid)
	// model.DB.Exec("detele from users_backdrops WHERE student_id = ?", sid)
	handler.SendResponse(c, "删除成功", nil)
}
