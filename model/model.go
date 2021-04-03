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

func GetUserInfo(id string) (error, User) {
	var u User
	result := DB.Where("student_id = ?", id).First(&u)
	return result.Error, u
}

func UpdateUserInfo(user User) error {
	var u User
	result := DB.Model(&u).Where("student_id = ?", user.StudentID).Update(user)
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
	punch.Time = time.Now().Format("2006-01-02 15:04:05")
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
	DB.Model(&user).Where("student_id = ? ", id).Update("gold", gold+user.Gold)
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
		Time:           time.Now().Format("2006-01-02 15:04:05"),
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

//default:

func GetMonthList() ([]UserRanking, string) {
	var (
		PunchHistory []PunchHistory
		ranks        []UserRanking
		Ranks        []MonthList
		s            []string
		Rank         MonthList
	)
	if err := DB.Where("month < ? ", int(time.Now().Month())).First(&Rank).Error; err == nil {
		DB.Delete(Rank, "month < ? ", int(time.Now().Month()))
		DB.Where("month = ?", int(time.Now().Month())).Find(&PunchHistory)
		for _, ph := range PunchHistory {
			s = append(s, ph.StudentID)
		}
		UserNumbers := GetOrder(s)
		r := 1
		for i, num := range UserNumbers {
			if i > 0 && num.Number < UserNumbers[i-1].Number {
				r++
			}
			Rank = MonthList{
				StudentID: num.StudentId,
				Ranking:   r,
				Month:     int(time.Now().Month()),
				Number:    num.Number,
			}
			DB.Create(&Rank)
		}
	}
	for i := 1; i <= 5; i++ {
		var Rank []MonthList
		DB.Where("ranking = ? ", i).Find(&Rank)
		Ranks = append(Ranks, Rank...)
	}
	for _, ran := range Ranks {
		var u User
		DB.Where("student_id = ? ", ran.StudentID).First(&u)
		var rank UserRanking
		rank.Number = ran.Number
		rank.Ranking = ran.Ranking
		rank.StudentId = ran.StudentID
		rank.Name = u.Name
		ranks = append(ranks, rank)
	}
	return ranks, ""
}
func GetWeekList() ([]UserRanking, string) {
	var (
		PunchHistory []PunchHistory
		ranks        []UserRanking
		Ranks        []WeekList
		s            []string
		Rank         WeekList
	)
	if err := DB.Where("day <= ? ", time.Now().YearDay()-7).First(&Rank).Error; err == nil {
		DB.Delete(Rank, "day <= ? ", time.Now().YearDay()-7)
		DB.Table("punch_histories").Select("student_id").Where("day >= ?", int(time.Now().YearDay())-7).Scan(&PunchHistory)
		for _, ph := range PunchHistory {
			s = append(s, ph.StudentID)
		}
		UserNumbers := GetOrder(s)
		r := 1
		for i, num := range UserNumbers {
			if i > 0 && num.Number < UserNumbers[i-1].Number {
				r++
			}
			Rank := WeekList{
				StudentID: num.StudentId,
				Ranking:   r,
				Day:       time.Now().YearDay(),
				Number:    num.Number,
			}
			DB.Create(&Rank)
		}
	}
	for i := 1; i <= 5; i++ {
		var Rank []WeekList
		DB.Where("ranking = ? ", i).Find(&Rank)
		Ranks = append(Ranks, Rank...)
	}
	for _, ran := range Ranks {
		var u User
		DB.Where("student_id = ? ", ran.StudentID).First(&u)
		var rank UserRanking
		rank.Number = ran.Number
		rank.Ranking = ran.Ranking
		rank.StudentId = ran.StudentID
		rank.Name = u.Name
		ranks = append(ranks, rank)
	}

	// if len(PunchHistory) == 0 {
	// 	return nil, "未检索到该时间段的打卡信息"
	// }
	return ranks, ""
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

func GetListHistory(id string) ListHistories {
	var (
		history1 ListHistory
		history2 ListHistory
	)
	DB.Where("student_id = ? AND type = 1 ", id).First(&history1)
	DB.Where("student_id = ? AND type = 2 ", id).First(&history2)
	histories := ListHistories{
		StudentID:   id,
		WeekFormer:  history1.Former,
		WeekAfter:   history1.After,
		MonthFormer: history2.Former,
		MonthAfter:  history2.After,
	}
	return histories
}

func GetBackdropPrice() []Backdrop {
	var backdrop []Backdrop
	DB.Find(&backdrop)
	return backdrop
}

func ChangeWeekRanking(id string, ranking int) (error, string) {
	gold := 48 + ranking*2
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < gold {
		return nil, "金币不足"
	}

	//修改用户金币
	DB.Model(&user).Where("student_id = ? ", id).Update("gold", user.Gold-gold)

	//创建金币历史
	price := gold
	history := GoldHistory{
		StudentID:      id,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
		ChangeNumber:   -price,
		ResidualNumber: user.Gold,
		Reason:         "兑换周排名:前进" + strconv.Itoa(ranking) + "名",
	}
	if result := DB.Create(&history); result.Error != nil {
		return result.Error, ""
	}

	//创建修改历史
	History := ListHistory{
		StudentID: id,
		Type:      1,
		Former:    0,
		After:     ranking,
	}

	//修改排行榜
	rank := WeekList{
		StudentID: id,
		Ranking:   ranking,
		Day:       time.Now().YearDay(),
		Number:    0,
	}
	UserNumber, str := GetWeekList()
	if str != "" {
		return nil, str
	}
	for _, UandN := range UserNumber {
		if UandN.StudentId == id {
			History.Former = rank.Ranking
			rank.Number = UandN.Number
			break
		}
	}
	if err := CreateRankingHistory(History); err != nil {
		return err, ""
	}
	err := ChangeWeekList(rank)
	return err, ""
}

func CreateRankingHistory(history ListHistory) error {
	DB.Where("type = ? AND student_id = ? ", history.Type, history.StudentID).Delete(&history)
	err := DB.Create(&history).Error
	if err != nil {
		return err
	}
	return nil
}

func ChangeWeekList(rank WeekList) error {
	var r WeekList
	DB.Where("day <= ? ", rank.Day-7).Delete(&r)
	var ranks []WeekList
	if err := DB.Where("student_id = ? ", rank.StudentID).First(&r).Error; err != nil {
		DB.Model(r).Where("ranking < ? AND ranking >= ? ", rank.Ranking, r.Ranking).Find(&ranks)
	} else {
		DB.Model(r).Where("ranking < ? ", rank.Ranking).Find(&ranks)
	}
	//后面的排名++
	for _, Rank := range ranks {
		Rank.Ranking++
		DB.Model(r).Where("student_id = ? ", Rank.StudentID).Update("ranking", Rank.Ranking)
	}

	err := DB.Create(&rank).Error
	return err
}

func ChangeMonthRanking(id string, ranking int) (error, string) {
	gold := 48 + ranking*2
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < gold {
		return nil, "金币不足"
	}

	//修改用户金币
	DB.Model(&user).Where("student_id = ? ", id).Update("gold", user.Gold-gold)

	//创建金币历史
	price := gold
	history := GoldHistory{
		StudentID:      id,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
		ChangeNumber:   -price,
		ResidualNumber: user.Gold,
		Reason:         "兑换月排名:前进" + strconv.Itoa(ranking) + "名",
	}
	if result := DB.Create(&history); result.Error != nil {
		return result.Error, ""
	}

	//创建修改历史
	History := ListHistory{
		StudentID: id,
		Type:      2,
		Former:    0,
		After:     ranking,
	}

	//修改排行榜
	rank := MonthList{
		StudentID: id,
		Ranking:   ranking,
		Month:     int(time.Now().Month()),
		Number:    0,
	}
	UserNumber, str := GetMonthList()
	if str != "" {
		return nil, str
	}
	for _, UandN := range UserNumber {
		if UandN.StudentId == id {
			History.Former = rank.Ranking
			rank.Number = UandN.Number
			break
		}
	}
	if err := CreateRankingHistory(History); err != nil {
		return err, ""
	}
	err := ChangeMonthList(rank)
	return err, ""
}

func ChangeMonthList(rank MonthList) error {
	var r MonthList
	DB.Where("month != ? ", rank.Month).Delete(&r)
	var ranks []MonthList
	if err := DB.Where("student_id = ? ", rank.StudentID).First(&r).Error; err != nil {
		DB.Model(r).Where("ranking < ? AND ranking >= ? ", rank.Ranking, r.Ranking).Find(&ranks)
	} else {
		DB.Model(r).Where("ranking < ? ", rank.Ranking).Find(&ranks)
	}
	//后面的排名++
	for _, Rank := range ranks {
		Rank.Ranking++
		DB.Model(r).Where("student_id = ? ", Rank.StudentID).Update("ranking", Rank.Ranking)
	}
	err := DB.Create(&rank).Error
	return err
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
	DB.Model(&user).Where("student_id = ? ", id).Update("gold", user.Gold-backdrop.Price)
	//创建金币历史
	s := strconv.Itoa(BackdropID)
	history := GoldHistory{
		StudentID:      id,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
		ChangeNumber:   -backdrop.Price,
		ResidualNumber: user.Gold,
		Reason:         "兑换背景 " + s,
	}
	DB.Create(&history)
	var usersBackdrop UsersBackdrop
	usersBackdrop.BackdropID = BackdropID
	usersBackdrop.StudentID = id
	result := DB.Create(&usersBackdrop)
	return result.Error, ""
}

func GetBackdrop(id string) []Backdrop {
	var backdrops []Backdrop
	DB.Table("users_backdrops").Where("student_id = ? ", id).Find(&backdrops)
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
	switch id {
	case "1":
		return "健康"
	case "2":
		return "运动"
	case "3":
		return "学习"
	}
	return "typeID错误"
}

func GetUserRanking(id string, Type string) int {
	var u WeekList
	err := DB.Where("student_id = ? ", id).First(&u).Error
	if err != nil {
		return u.Ranking
	}
	return -1
}
