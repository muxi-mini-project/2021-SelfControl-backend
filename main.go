package main

import (
	"SC/model"
	"SC/routers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

// @title Self_Control API
// @version 1.0.0
// @description 自控力API
// @termsOfService http://swagger.io/terrms/
// @contact.name TAODEI
// @contact.email 864978550@qq.com
// @host 124.71.184.107
// @BasePath: /api/v1
// @Schemes http

func main() {
	//model.DB, err = gorm.Open("mysql", "tao:12345678@/Self_Control?parseTime=True")
	model.DB, err = gorm.Open("mysql", "root:1234@/Self_Control?parseTime=True")
	if err != nil {
		panic(err)
	}
	//自动建表 迁移？
	//model.DB.AutoMigrate(&model.User{})
	r := gin.Default()
	routers.Router(r)
	dbtest()
	r.Run(":2333")
	defer model.DB.Close()
}

func dbtest() {
	var (
		backdrop1 model.Backdrop
		backdrop2 model.Backdrop
		backdrop3 model.Backdrop
		backdrop4 model.Backdrop
		backdrop5 model.Backdrop

		punch1_1  = model.PunchContent{Type: "健康", Title: "吃水果"}
		punch1_2  = model.PunchContent{Type: "健康", Title: "吃早餐"}
		punch1_3  = model.PunchContent{Type: "健康", Title: "多喝水"}
		punch1_4  = model.PunchContent{Type: "健康", Title: "拒绝宵夜"}
		punch1_5  = model.PunchContent{Type: "健康", Title: "拒绝饮料"}
		punch1_6  = model.PunchContent{Type: "健康", Title: "拒绝久坐"}
		punch1_7  = model.PunchContent{Type: "健康", Title: "早起"}
		punch1_8  = model.PunchContent{Type: "健康", Title: "不要熬夜"}
		punch1_9  = model.PunchContent{Type: "健康", Title: "不翘二郎腿"}
		punch1_10 = model.PunchContent{Type: "健康", Title: "早起空腹喝水"}

		punch2_1 = model.PunchContent{Type: "运动", Title: "跑步"}
		punch2_2 = model.PunchContent{Type: "运动", Title: "俯卧撑"}
		punch2_3 = model.PunchContent{Type: "运动", Title: "跳绳"}
		punch2_4 = model.PunchContent{Type: "运动", Title: "仰卧起坐"}
		punch2_5 = model.PunchContent{Type: "运动", Title: "散步"}
		punch2_6 = model.PunchContent{Type: "运动", Title: "拉伸"}
		punch2_7 = model.PunchContent{Type: "运动", Title: "打篮球"}
		punch2_8 = model.PunchContent{Type: "运动", Title: "健身"}
		punch2_9 = model.PunchContent{Type: "运动", Title: "骑车"}

		punch3_1  = model.PunchContent{Type: "学习", Title: "自习"}
		punch3_2  = model.PunchContent{Type: "学习", Title: "阅读新闻"}
		punch3_3  = model.PunchContent{Type: "学习", Title: "练习乐器"}
		punch3_4  = model.PunchContent{Type: "学习", Title: "学习新语言"}
		punch3_5  = model.PunchContent{Type: "学习", Title: "背单词"}
		punch3_6  = model.PunchContent{Type: "学习", Title: "看纪录片"}
		punch3_7  = model.PunchContent{Type: "学习", Title: "做今日计划"}
		punch3_8  = model.PunchContent{Type: "学习", Title: "听力练习"}
		punch3_9  = model.PunchContent{Type: "学习", Title: "练字"}
		punch3_10 = model.PunchContent{Type: "学习", Title: "英语阅读训练"}

		lstao model.User
	)
	backdrop1.PictureUrl = "www.4399.com"
	backdrop1.Price = 50
	backdrop2.PictureUrl = "www.7k7k.com"
	backdrop2.Price = 50
	backdrop3.PictureUrl = "www.3839.com"
	backdrop3.Price = 100
	backdrop4.PictureUrl = "www.bilibili.com"
	backdrop4.Price = 100
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
