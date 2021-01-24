package routers

import (
	"SC/handler"
	"SC/handler/user"

	"github.com/gin-gonic/gin"
)

//Router 1
func Router(r *gin.Engine) {

	//user:
	g1 := r.Group("/api/v1/user")
	{
		//登陆
		g1.POST("/", user.Login)

		//用户主页
		g1.GET("/", user.Homepage)

		//修改用户信息
		g1.PUT("/", user.ChangeInfo)

		//金币数及历史
		g1.GET("/gold", user.Gold)

		//成就
		g1.GET("/achievement", user.Achievement)

		//用户打卡数
		g1.GET("/punch", user.PunchNumber)

		//打卡内容公开与否
		g1.POST("/privacy", user.Privary)
	}

	//显示当前类型所有打卡
	r.GET("/api/v1/punchs/{type}", handler.Punchs)

	g2 := r.Group("/api/v1/punch")
	{

		//我的打卡
		g2.GET("/", handler.MyPunch)

		//完成打卡
		g2.POST("/", handler.CompletePunch)

		//移除标签
		g2.DELETE("/", handler.DeletePunch)
	}
	//default:
	//排行榜
	r.GET("/api/v1/list", handler.List)

	//兑换排名
	r.POST("/api/v1/list", handler.ChangeRanking)

	//背景价格
	r.GET("/api/v1/backdrop", handler.BackdropPrice)

	//兑换背景
	r.POST("/api/v1/backdrop", handler.ChangeBackdrop)
}
