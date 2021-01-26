package model

import (
	//"SC/util"
	"errors"
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
		//log.Println("errs++++++++++")
		//log.Println(err.Error())
		return "", errors.New(ErrorReasonReLogin + "666")
	}
	if err := token.Claims.Valid(); err != nil {
		//fmt.Println(err.Error())
		return "", errors.New(ErrorReasonReLogin)
	}

	return claims.StudentID, nil //token.Jwt 结构体类型
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

func GetAchievement(id string) []string {
	var achievements []Achievement
	DB.Where("student_id = ?", id).Find(&achievements)
	var achs []string
	for i := 0; i < len(achievements); i++ {
		achs[i] = achievements[i].Achievement
	}
	return achs
}

//-----------------------------------------------
//punch:
//Punchs 是数组结构体 punchs 是 []Punch   Punch 为 Title 与 Number 的结构体类型
func GetPunchAndNumber(id string) Punchs {
	var punchs []UsersPunch
	DB.Where("student_id = ?", id).Find(&punchs)
	var punchs2 []Punch
	for i := 0; i < len(punchs); i++ {
		punchs2[i].Title = punchs[i].Title
		punchs2[i].Number = punchs[i].Number
	}
	Punchs := Punchs{Punchs: punchs2}
	return Punchs

}

func GetPunchs(Type string) Punchs2 {
	var punchs []PunchContent
	DB.Where("type = ?", Type).Find(&punchs)
	var punchs2 []Punch2
	for i := 0; i < len(punchs); i++ {
		punchs2[i].Title = punchs[i].Title
	}
	Punchs := Punchs2{Punchs: punchs2}
	return Punchs
}

func GetMyPunch(id string) Punchs {
	var punchs []UsersPunch
	DB.Where("student_id = ?", id).Find(&punchs)
	var punchs2 []Punch
	for i := 0; i < len(punchs); i++ {
		punchs2[i].Title = punchs[i].Title
		punchs2[i].Number = punchs[i].Number
	}
	Punchs := Punchs{Punchs: punchs2}
	return Punchs
}

func TodayPunch(StudentId string, title string) Choice {
	var punch PunchHistory
	today := time.Now().YearDay()
	result := DB.Where("student_id = ? AND title = ? AND day = ?", StudentId, title, today).First(&punch)

	var choice bool
	if result.Error != nil {
		choice = false
	}
	choice = true
	Choice := Choice{Choice: choice}
	return Choice
}

func CompletePunch(id string, title string) error {
	var punch PunchHistory
	//result:=DB.Model(&punch).Where("id = ? AND title = ?",id ,title).Update("choice", "hello")
	punch.Title = title
	punch.StudentID = id
	punch.Time = time.Now()
	punch.Month = int(time.Now().Month())
	//punch.Week=time.Now().ISOWeek()
	punch.Day = time.Now().YearDay()
	result := DB.Create(&punch)
	return result.Error
}

func DeletePunch(id string, title string) error {
	var u UsersPunch
	DB.Where("student_id = ? AND title = ? ", id, title).First(&u)
	result := DB.Delete(&u)
	return result.Error
}

//-----------------------------------------------
//default:

func List(Type string) []UserAndNumber {
	type Result struct {
		StudentID string
	}
	var results []Result
	if Type == "month" {
		DB.Table("punch_histories").Select("student_id").Where("month = ?", int(time.Now().Month())).Scan(&results)
	}
	if Type == "week" {
		DB.Table("punch_histories").Select("student_id").Where("day >= ?", int(time.Now().YearDay())-7).Scan(&results)
	}
	var s []string
	for i := 0; i < len(results); i++ {
		s[i] = results[i].StudentID
	}
	return GetOrder(s)
	//var punch PunchHistory
	//DB.Where("student_id = ? AND month = ? ", int(time.Now().Month())).Find(&punch)

}
func GetOrder(s []string) []UserAndNumber {
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
	return Numbers
}

func GetListPrice() ListPrices {
	var prices []ListPrice
	DB.Find(&prices)
	Prices := ListPrices{ListPrice: prices}
	return Prices
}

func GetBackdropPrice() Backdrops {
	var backdrop []Backdrop
	DB.Find(&backdrop)
	Backdrops := Backdrops{Backdrop: backdrop}
	return Backdrops
}

func ChangeRanking(id string, title string) error {
	var listPrice ListPrice
	DB.Where("title = ? ", title).First(&listPrice)
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < listPrice.Price {
		return errors.New("金币不足")
	}
	var u User
	u = user
	u.Gold -= listPrice.Price
	result := DB.Model(&u).Update(user)
	return result.Error
}

func ChangeBackdrop(id string, BackdropId int) error {
	var backdrop Backdrop
	DB.Where("backdrop_id = ? ", BackdropId).First(&backdrop)
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < backdrop.Price {
		return errors.New("金币不足")
	}
	var u User
	u = user
	u.Gold -= backdrop.Price
	result := DB.Model(&u).Update(user)
	if result.Error != nil {
		return result.Error
	}
	var usersBackdrop UsersBackdrop
	usersBackdrop.BackdropID = BackdropId
	usersBackdrop.StudentID = id
	result = DB.Create(&usersBackdrop)
	return result.Error
}

func GetBackdrop(id string) Backdrops {
	var backdrops []Backdrop
	DB.Table("users_backdrops").Where("student_id = ? ", id).First(&backdrops)
	Backdrops := Backdrops{Backdrop: backdrops}
	return Backdrops
}
