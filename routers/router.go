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

		//用户信息
		g1.GET("/", user.UserInfo)

		//修改用户信息
		g1.PUT("/", user.ChangeUserInfo)

		//金币历史
		g1.GET("/goldhistory", user.GoldHistory)

		//用户打卡数
		g1.GET("/punch", user.PunchAndNumber)

		//获取某用户隐私
		g1.GET("/privacy/:id", user.GetPrivacy)

	}

	//显示当前类型所有打卡
	r.GET("/api/v1/punchs/:type_id", handler.Punchs)

	g2 := r.Group("/api/v1/punch")
	{
		//我的打卡
		g2.GET("/", handler.MyPunch)

		//判断今天是否已打卡
		g2.GET("/today/:title_id", handler.TodayPunch)

		//完成打卡
		g2.POST("/", handler.CompletePunch)

		//添加标签
		g2.POST("/create", handler.CreatePunch)

		//移除标签
		g2.DELETE("/", handler.DeletePunch)

		//获取某用户标签
		g1.GET("/punch/:id", handler.GetPunchs)

		//月报
		g2.GET("/month", handler.Monthly)

		//获取用户某天的打卡
		g2.GET("/day/:day", handler.GetDayPunchs)

	}
	//default:
	//排行榜
	r.GET("/api/v1/lists/:type", handler.List)

	//获取兑换排行历史
	r.GET("/api/v1/list/history", handler.ListHistory)

	//用户排名
	r.GET("/api/v1/list/user/:id/:type", handler.UserRanking)

	//兑换排名
	r.PUT("/api/v1/list/:type", handler.ChangeRanking)

	//背景价格
	r.GET("/api/v1/backdrop", handler.BackdropPrice)

	//兑换背景
	r.PUT("/api/v1/backdrop", handler.ChangeBackdrop)

	//我的背景
	r.GET("/api/v1/backdrops", handler.MyBackdrops)
}
