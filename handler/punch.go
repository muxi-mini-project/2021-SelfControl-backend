package handler

import (
	"SC/model"

	"github.com/gin-gonic/gin"
)

func Punchs(c *gin.Context) {
	var punch model.Punch
	if err := c.BindJSON(&punch); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}
func MyPunch(c *gin.Context) {

}
func DeletePunch(c *gin.Context) {

}
func CompletePunch(c *gin.Context) {

}
