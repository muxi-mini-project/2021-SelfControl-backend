package main

import (
	"SC/config"
	"SC/model"
	"SC/router"
	"SC/service/punch"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// @title Self_Control API
// @version 1.0.0
// @description 自控力API
// @termsOfService http://swagger.io/terrms/
// @contact.name TAODEI
// @contact.email tao_dei@qq.com
// @host self-control.muxixyz.com:2333
// @BasePath /api/v1
// @Schemes http

func main() {
	err := config.Init("./conf/config.yaml", "")
	if err != nil {
		panic(err)
	}

	model.InitDB()

	go concurrent() // 并发

	r := gin.Default()
	router.Router(r)
	init := os.Getenv("INIT")
	if init == "yes" {
		dbInitTest()
	}
	dbInitTest()
	port := viper.GetString("port")
	r.Run(port)
	defer model.DB.Close()
}

// 并发
func concurrent() {
	punch.UpdatePunchHistoryEveryDay() // 每日更新用户标签
}

func dbInitTest() {
	var (
		backdrop1 model.Backdrop
		backdrop2 model.Backdrop
		backdrop3 model.Backdrop
		backdrop4 model.Backdrop
		backdrop5 model.Backdrop

		punch1_1  = model.PunchContent{Type: "健康", Title: "吃水果", ID: 1500045}
		punch1_2  = model.PunchContent{Type: "健康", Title: "吃早餐", ID: 1500057}
		punch1_3  = model.PunchContent{Type: "健康", Title: "多喝水", ID: 1500007}
		punch1_4  = model.PunchContent{Type: "健康", Title: "拒绝夜宵", ID: 1500068}
		punch1_5  = model.PunchContent{Type: "健康", Title: "拒绝饮料", ID: 1500043}
		punch1_6  = model.PunchContent{Type: "健康", Title: "拒绝久坐", ID: 1500083}
		punch1_7  = model.PunchContent{Type: "健康", Title: "早起", ID: 1500024}
		punch1_8  = model.PunchContent{Type: "健康", Title: "早睡", ID: 1500046}
		punch1_9  = model.PunchContent{Type: "健康", Title: "不翘二郎腿", ID: 1500031}
		punch1_10 = model.PunchContent{Type: "健康", Title: "早起空腹喝水", ID: 1500041}

		punch2_1 = model.PunchContent{Type: "运动", Title: "跑步", ID: 1500026}
		punch2_2 = model.PunchContent{Type: "运动", Title: "俯卧撑", ID: 1500008}
		punch2_3 = model.PunchContent{Type: "运动", Title: "跳绳", ID: 1500074}
		punch2_4 = model.PunchContent{Type: "运动", Title: "仰卧起坐", ID: 1500066}
		punch2_5 = model.PunchContent{Type: "运动", Title: "散步", ID: 1500033}
		punch2_6 = model.PunchContent{Type: "运动", Title: "拉伸", ID: 1500073}
		punch2_7 = model.PunchContent{Type: "运动", Title: "打篮球", ID: 1500071}
		punch2_8 = model.PunchContent{Type: "运动", Title: "健身", ID: 1500020}
		punch2_9 = model.PunchContent{Type: "运动", Title: "骑车", ID: 1500065}

		punch3_1  = model.PunchContent{Type: "学习", Title: "自习", ID: 1500060}
		punch3_2  = model.PunchContent{Type: "学习", Title: "阅读新闻", ID: 1500025}
		punch3_3  = model.PunchContent{Type: "学习", Title: "练习乐器", ID: 1500069}
		punch3_4  = model.PunchContent{Type: "学习", Title: "学习新语言", ID: 1500044}
		punch3_5  = model.PunchContent{Type: "学习", Title: "背单词", ID: 1500010}
		punch3_6  = model.PunchContent{Type: "学习", Title: "看纪录片", ID: 1500028}
		punch3_7  = model.PunchContent{Type: "学习", Title: "做今日计划", ID: 1500067}
		punch3_8  = model.PunchContent{Type: "学习", Title: "听力训练", ID: 1500011}
		punch3_9  = model.PunchContent{Type: "学习", Title: "练字", ID: 1500059}
		punch3_10 = model.PunchContent{Type: "学习", Title: "英语阅读训练", ID: 1500085}

		lstao model.User
	)
	backdrop1.BackdropID = 1
	backdrop1.PictureUrl = "www.4399.com"
	backdrop1.Price = 50
	backdrop2.BackdropID = 2
	backdrop2.PictureUrl = "www.7k7k.com"
	backdrop2.Price = 50
	backdrop3.BackdropID = 3
	backdrop3.PictureUrl = "www.3839.com"
	backdrop3.Price = 100
	backdrop4.BackdropID = 4
	backdrop4.PictureUrl = "www.bilibili.com"
	backdrop4.Price = 100
	backdrop5.BackdropID = 5
	backdrop5.PictureUrl = "www.cf.qq.com"
	backdrop5.Price = 100

	// punch1.Content = "听分享时睡觉"
	// punch1.Title = "睡觉"
	// punch1.Type = "健康"
	// punch2.Content = "吃两碗饭"
	// punch2.Title = "吃饭"
	// punch2.Type = "运动"
	// punch3.Content = "上课滑水"
	// punch3.Title = "滑水"
	// punch3.Type = "学习"

	lstao.StudentID = "2020213675"
	lstao.Name = "TAODEI"
	lstao.Gold = 520
	lstao.Password = "2333333"
	lstao.UserPicture = "www.muxi.com"
	lstao.Privacy = 1

	model.DB.Create(&backdrop1)
	model.DB.Create(&backdrop2)
	model.DB.Create(&backdrop3)
	model.DB.Create(&backdrop4)
	model.DB.Create(&backdrop5)

	model.DB.Create(&punch1_1)
	model.DB.Create(&punch1_2)
	model.DB.Create(&punch1_3)
	model.DB.Create(&punch1_4)
	model.DB.Create(&punch1_5)
	model.DB.Create(&punch1_6)
	model.DB.Create(&punch1_7)
	model.DB.Create(&punch1_8)
	model.DB.Create(&punch1_9)
	model.DB.Create(&punch1_10)
	model.DB.Create(&punch2_1)
	model.DB.Create(&punch2_2)
	model.DB.Create(&punch2_3)
	model.DB.Create(&punch2_4)
	model.DB.Create(&punch2_5)
	model.DB.Create(&punch2_6)
	model.DB.Create(&punch2_7)
	model.DB.Create(&punch2_8)
	model.DB.Create(&punch2_9)
	model.DB.Create(&punch3_1)
	model.DB.Create(&punch3_2)
	model.DB.Create(&punch3_3)
	model.DB.Create(&punch3_4)
	model.DB.Create(&punch3_5)
	model.DB.Create(&punch3_6)
	model.DB.Create(&punch3_7)
	model.DB.Create(&punch3_8)
	model.DB.Create(&punch3_9)
	model.DB.Create(&punch3_10)

	model.DB.Create(&lstao)
}
