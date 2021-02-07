package model

import (
	//"SC/util"
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	ErrorReasonServerBusy = "服务器繁忙"
	ErrorReasonReLogin    = "请重新登陆"
)

//var Secret = "vinegar" //加醋

//Jwt
type Jwt struct {
	StudentID string `json:"student_id"`
	jwt.StandardClaims
}

//user:
func VerifyToken(strToken string) (string, error) {
	//解析
	token, err := jwt.ParseWithClaims(strToken, &Jwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("vinegar"), nil
	})

	if err != nil {
		return "", errors.New(ErrorReasonServerBusy + ",或token解析失败")
	}
	claims, ok := token.Claims.(*Jwt)
	if !ok {
		return "", errors.New(ErrorReasonReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return "", errors.New(ErrorReasonReLogin)
	}
	return claims.StudentID, nil
}

func GetUserInfo(id string) User {
	var u User
	DB.Where("student_id = ?", id).First(&u)
	return u
}

func UpdateUserInfo(user User) error {
	var u User
	result := DB.Model(&u).Update(user)
	return result.Error
}

/*
func GetAchievement(id string) []string {
	var achievements []Achievement
	DB.Where("student_id = ?", id).Find(&achievements)
	var achs []string
	for i := 0; i < len(achievements); i++ {
		achs[i] = achievements[i].Achievement
	}
	return achs
}*/

//-----------------------------------------------
//punch:
//Punch 为 Title 与 Number 的结构体类型
func GetPunchAndNumber(id string) []Punch {
	var punchs []UsersPunch
	DB.Where("student_id = ?", id).Find(&punchs)
	var punchs2 []Punch
	var Punch Punch
	for i := 0; i < len(punchs); i++ {
		Punch.Title = punchs[i].Title
		Punch.Number = punchs[i].Number
		punchs2 = append(punchs2, Punch)
	}
	//Punchs := Punchs{Punchs: punchs2}
	return punchs2

}

func GetPunchs(TypeID string) []Punch2 {
	Type := Type(TypeID)
	var punchs []PunchContent
	DB.Where("type = ?", Type).Find(&punchs)
	var punchs2 []Punch2
	var Punch Punch2
	for i := 0; i < len(punchs); i++ {
		Punch.Title = punchs[i].Title
		Punch.ID = punchs[i].ID
		punchs2 = append(punchs2, Punch)
	}
	//Punchs := Punchs2{Punchs: punchs2}
	return punchs2
}

func GetMyPunch(id string) []Punch {
	var punchs []UsersPunch
	DB.Where("student_id = ?", id).Find(&punchs)
	var punchs2 []Punch
	var Punch Punch
	for i := 0; i < len(punchs); i++ {
		Punch.Title = punchs[i].Title
		Punch.Number = punchs[i].Number
		punchs2 = append(punchs2, Punch)
	}
	return punchs2
}

func TodayPunch(StudentId string, TitleID int) Choice {
	var Punch PunchContent
	DB.Where("id = ?", TitleID).First(&Punch)
	var punch PunchHistory
	today := time.Now().YearDay()
	result := DB.Where("student_id = ? AND title = ? AND day = ?", StudentId, Punch.Title, today).First(&punch)
	var choice bool
	if result.Error != nil {
		choice = false
	} else {
		choice = true
	}
	Choice := Choice{Choice: choice}
	return Choice
}

func CompletePunch(id string, title string, gold int) error {
	var punch PunchHistory
	//result:=DB.Model(&punch).Where("id = ? AND title = ?",id ,title).Update("choice", "hello")
	punch.Title = title
	punch.StudentID = id
	punch.Time = time.Now()
	punch.Month = int(time.Now().Month())
	//punch.Week=time.Now().ISOWeek()
	punch.Day = time.Now().YearDay()
	if result := DB.Create(&punch); result.Error != nil {
		return result.Error
	}
	//log.Printf("%v+++++++++++\n", punch)

	//修改用户金币
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	DB.Model(&user).Update("gold", gold+user.Gold)
	// var u User
	// u = user
	// u.Gold += gold
	// log.Println(user)
	// log.Println(u)
	// if result := DB.Model(&u).Update(user); result.Error != nil {
	// 	return result.Error
	// }
	//log.Println(u)
	s := strconv.Itoa(gold)
	//创建金币历史
	history := GoldHistory{
		StudentID:      id,
		Time:           time.Now(),
		ChangeNumber:   gold,
		ResidualNumber: user.Gold,
		Reason:         "完成打卡+" + s + "金币",
	}
	result := DB.Create(&history)
	return result.Error
}

func DeletePunch(id string, title string) (string, error) {
	var u UsersPunch
	if result := DB.Where("student_id = ? AND title = ? ", id, title).First(&u); result.Error != nil {
		return "用户未选择该标签", nil
	}
	result := DB.Delete(&u)
	return "", result.Error
}

//-----------------------------------------------
//default:

func List(Type string) ([]UserAndNumber, string) {
	var PunchHistory []PunchHistory
	if Type == "month" {
		DB.Where("month = ?", int(time.Now().Month())).Find(&PunchHistory)
		//DB.Table("punch_histories").Select("student_id").Where("month = ?", int(time.Now().Month())).Scan(&PunchHistory)
	}
	if Type == "week" {
		//result := DB.Where("day >= ?", int(time.Now().YearDay())-7).Find(&PunchHistory)
		DB.Table("punch_histories").Select("student_id").Where("day >= ?", int(time.Now().YearDay())-7).Scan(&PunchHistory)
	}
	var s []string
	for i := 0; i < len(PunchHistory); i++ {
		s = append(s, PunchHistory[i].StudentID)
	}
	if len(s) == 0 {
		var a []UserAndNumber
		return a, "未检索到该时间段的打卡信息"
	}
	return GetOrder(s)
	//var punch PunchHistory
	//DB.Where("student_id = ? AND month = ? ", int(time.Now().Month())).Find(&punch)

}
func GetOrder(s []string) ([]UserAndNumber, string) {
	var Numbers []UserAndNumber
	Number := UserAndNumber{StudentId: s[0], Number: 1}
	Numbers = append(Numbers, Number)
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(Numbers); j++ {
			if Numbers[j].StudentId == s[i] {
				Numbers[j].Number++
				break
			}
			if j == len(Numbers)-1 {
				Number := UserAndNumber{StudentId: s[i]}
				Numbers = append(Numbers, Number)
			}
		}
	}
	n := len(Numbers)
	for i := 0; i < n-1; i++ {
		max := i
		for j := i + 1; j < n; j++ {
			if Numbers[j].Number > Numbers[max].Number {
				max = j
			}
		}
		Numbers[i], Numbers[max] = Numbers[max], Numbers[i]
	}
	return Numbers, ""
}

func GetListPrice() []ListPrice {
	var prices []ListPrice
	DB.Find(&prices)
	return prices
}

func GetBackdropPrice() []Backdrop {
	var backdrop []Backdrop
	DB.Find(&backdrop)
	return backdrop
}

func ChangeRanking(id string, ranking string) (error, string) {
	var listPrice ListPrice
	DB.Where("ranking = ? ", ranking).First(&listPrice)
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < listPrice.Price {
		return nil, "金币不足"
	}

	//修改用户金币
	DB.Model(&user).Update("gold", user.Gold-listPrice.Price)

	//创建金币历史
	price := listPrice.Price
	history := GoldHistory{
		StudentID:      id,
		Time:           time.Now(),
		ChangeNumber:   -price,
		ResidualNumber: user.Gold,
		Reason:         "兑换排名：" + ranking,
	}
	result := DB.Create(&history)
	return result.Error, ""
}

func ChangeBackdrop(id string, BackdropID int) (error, string) {
	var backdrop Backdrop
	DB.Where("backdrop_id = ? ", BackdropID).First(&backdrop)
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < backdrop.Price {
		return nil, "金币不足"
	}
	//修改用户金币
	DB.Model(&user).Update("gold", user.Gold-backdrop.Price)
	//创建金币历史
	s := strconv.Itoa(backdrop.BackdropID)
	history := GoldHistory{
		StudentID:      id,
		Time:           time.Now(),
		ChangeNumber:   -backdrop.Price,
		ResidualNumber: user.Gold,
		Reason:         "兑换背景 " + s,
	}
	//log.Printf("%v+++++++++\n---%v\n", BackdropID, backdrop.Price)
	DB.Create(&history)
	var usersBackdrop UsersBackdrop
	usersBackdrop.BackdropID = BackdropID
	usersBackdrop.StudentID = id
	result := DB.Create(&usersBackdrop)
	return result.Error, ""
}

func GetBackdrop(id string) []Backdrop {
	var backdrops []Backdrop
	DB.Table("users_backdrops").Where("student_id = ? ", id).First(&backdrops)
	return backdrops
}

func GetGoldHistory(id string) []GoldHistory {
	var histories []GoldHistory
	DB.Where("student_id = ? ", id).Find(&histories)
	return histories
}

func CreatePunch(id string, title string) (error, string) {
	var punch UsersPunch
	punch.StudentID = id
	punch.Title = title
	if result := DB.Where("student_id = ? AND title = ？", id, title).First(&punch); result.Error == nil {
		return nil, "该标签已选择"
	}
	result := DB.Create(&punch)
	return result.Error, ""
}

func Type(id string) string {
	if id == "1" {
		return "健康"
	} else if id == "2" {
		return "运动"
	} else {
		return "学习"
	}
}
