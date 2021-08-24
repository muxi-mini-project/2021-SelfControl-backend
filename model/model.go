package model

import (
	"strconv"
	"time"
)

func GetUserInfo(id string) (User, error) {
	var u User
	result := DB.Where("student_id = ?", id).First(&u)
	return u, result.Error
}

func UpdateUserInfo(user User) error {
	result := DB.Model(&user).Where("student_id = ?", user.StudentID).Update(user)
	return result.Error
}

//-----------------------------------------------
//punch:
func GetUserPunches(id string) []UsersPunch {
	var punchs []UsersPunch
	DB.Where("student_id = ?", id).Find(&punchs)
	return punchs
}

func GetPunchContentByTitle(title string) PunchContent {
	var p PunchContent
	DB.Where("title = ? ", title).First(&p)
	return p
}
func GetPunchContentById(TitleID int) PunchContent {
	var p PunchContent
	DB.Where("id = ? ", TitleID).First(&p)
	return p
}

func GetTodayPunchHistory(StudentId string, title string, today int) (PunchHistory, error) {
	var punch PunchHistory
	result := DB.Where("student_id = ? AND title = ? AND day = ?", StudentId, title, today).First(&punch)
	return punch, result.Error
}

func GetUserPunchHistoriesByDay(id string, day int) []PunchHistory {
	var histories []PunchHistory
	DB.Where("student_id = ? AND day = ? ", id, time.Now().Day()).Find(&histories)
	return histories
}

func GetPunchHistoriesByMonth(month int) []PunchHistory {
	var histories []PunchHistory
	DB.Where("month = ? ", month).Find(&histories)
	return histories
}

func GetUserPunchByTitle(id string, title string) (UsersPunch, error) {
	var pun UsersPunch
	err := DB.Where("student_id = ? AND title = ? ", id, title).First(&pun).Error
	return pun, err
}

func CreatePunchHistory(punch *PunchHistory) error {
	return DB.Create(punch).Error
}

func UpdateUserPunch(pun UsersPunch) error {

	result := DB.Model(&pun).Where("id = ?", pun.ID).Update(pun)
	return result.Error
}

func CreateGoldHistory(history *GoldHistory) error {
	return DB.Create(history).Error
}

// //根据 类型 获取其全部打卡
// func GetPunchs(TypeID string) []Punch2 {
// 	Type := Type(TypeID)
// 	var punchs []PunchContent
// 	DB.Where("type = ?", Type).Find(&punchs)
// 	var punchs2 []Punch2
// 	var Punch Punch2
// 	for i := 0; i < len(punchs); i++ {
// 		Punch.Title = punchs[i].Title
// 		Punch.ID = punchs[i].ID
// 		punchs2 = append(punchs2, Punch)
// 	}
// 	return punchs2
// }

func DeletePunch(u *UsersPunch) error {
	return DB.Delete(u).Error
}

func GetMonthly(id string) []Punch {
	var punchs []PunchHistory
	DB.Where("student_id = ? AND month = ? ", id, time.Now().Month()-1).Find(&punchs)
	var Punchs []Punch
	for _, punch := range punchs {
		tag := 0
		for i, punch2 := range Punchs {
			if punch2.Title == punch.Title {
				Punchs[i].Number++
				tag = 1
				break
			}
		}
		if tag == 0 {
			var Punch Punch
			Punch.Title = punch.Title
			Punch.Number++
			Punchs = append(Punchs, Punch)
		}
	}
	return Punchs
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
	// 把排名前10的加进来
	for i := 1; i <= 10; i++ {
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
		rank.UserPicture = u.UserPicture
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
	for i := 1; i <= 10; i++ {
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
		rank.UserPicture = u.UserPicture
		ranks = append(ranks, rank)
	}
	return ranks, ""
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

func ChangeWeekRanking(id string, ranking int) (string, error) {
	gold := 48 + ranking*2
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < gold {
		return "金币不足", nil
	}

	//创建修改历史
	History := ListHistory{
		StudentID: id,
		Type:      1,
	}
	UserNumber, str := GetWeekList()
	if str != "" {
		return str, nil
	}
	former := 0
	number := 0
	for _, UandN := range UserNumber {
		if UandN.StudentId == id {
			former = UandN.Ranking
			if former <= ranking {
				return "超出可兑换限制", nil
			}
			History.Former = former
			History.After = former - ranking
			number = UandN.Number
			break
		}
	}
	if former == 0 || number == 0 {
		return "错误:该用户兑换排名前没有该排名", nil
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
		return "", result.Error
	}

	if err := CreateRankingHistory(History); err != nil {
		return "", err
	}

	//修改排行榜
	rank := WeekList{
		StudentID: id,
		Ranking:   former - ranking,
		Day:       time.Now().YearDay(),
		Number:    number,
	}
	err := ChangeWeekList(rank)
	return "", err
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

	DB.Where("student_id = ? ", rank.StudentID).Delete(&rank)
	err := DB.Create(&rank).Error
	return err
}

func ChangeMonthRanking(id string, ranking int) (string, error) {
	gold := 48 + ranking*2
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < gold {
		return "金币不足", nil
	}

	//创建修改历史
	History := ListHistory{
		StudentID: id,
		Type:      2,
	}
	UserNumber, str := GetMonthList()
	if str != "" {
		return str, nil
	}
	former := 0
	number := 0
	for _, UandN := range UserNumber {
		if UandN.StudentId == id {
			former = UandN.Ranking
			if former <= ranking {
				return "超出可兑换限制", nil
			}
			History.Former = former
			History.After = former - ranking
			number = UandN.Number
			break
		}
	}
	if former == 0 || number == 0 {
		return "错误:该用户兑换排名前没有该排名", nil
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
		return "", result.Error
	}

	if err := CreateRankingHistory(History); err != nil {
		return "", err
	}
	//修改排行榜
	rank := MonthList{
		StudentID: id,
		Ranking:   former - ranking,
		Month:     int(time.Now().Month()),
		Number:    number,
	}
	err := ChangeMonthList(rank)
	return "", err
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
	DB.Where("student_id = ? ", rank.StudentID).Delete(&rank)
	err := DB.Create(&rank).Error
	return err
}

func ChangeBackdrop(id string, BackdropID int) (string, error) {
	var backdrop Backdrop
	DB.Where("backdrop_id = ? ", BackdropID).First(&backdrop)
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < backdrop.Price {
		return "金币不足", nil
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
	return "", result.Error
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

func CreatePunch(id string, title string) (string, error) {
	// var u UsersPunch
	// if result := DB.Where("student_id = ? AND title = ? ", id, title).First(&u); result.Error == nil {
	// 	return nil, "用户已选择该标签"
	// }
	var punch UsersPunch
	punch.StudentID = id
	punch.Title = title
	if result := DB.Where("student_id = ? AND title = ？", id, title).First(&punch); result.Error == nil {
		return "该标签已选择", nil
	}
	result := DB.Create(&punch)
	return "", result.Error
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
