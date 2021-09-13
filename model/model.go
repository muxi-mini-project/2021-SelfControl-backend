package model

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
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

// -----------------------------------------------
// punch:
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
	return DB.Model(&pun).Where("id = ?", pun.ID).Update(pun).Error
}

func CreateGoldHistory(history *GoldHistory) error {
	return DB.Create(history).Error
}

func DeletePunch(u *UsersPunch) error {
	return DB.Delete(u).Error
}

func GetUserPunchHistoriesByMonth(id string, month int) []PunchHistory {
	var histories []PunchHistory
	DB.Where("student_id = ? AND month = ? ", id, month).Find(&histories)
	return histories
}

// 根据 类型 获取其全部打卡
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
// ------------------
// default:

func CreateMonthlist(Rank *MonthList) {
	DB.Create(Rank)
}

func GetMonthRanks(rank int) []MonthList {
	var Ranks []MonthList
	DB.Where("ranking = ? ", rank).Find(&Ranks)
	return Ranks
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

// 事务
func CreateGoldAndRankHistory(history *GoldHistory, History *ListHistory) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(history).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(History).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

func ChangeBackdrop(id string, BackdropID int) (string, error) {
	var backdrop Backdrop
	DB.Where("backdrop_id = ? ", BackdropID).First(&backdrop)
	var user User
	DB.Where("student_id = ? ", id).First(&user)
	if user.Gold < backdrop.Price {
		return "金币不足", nil
	}
	// 修改用户金币
	DB.Model(&user).Where("student_id = ? ", id).Update("gold", user.Gold-backdrop.Price)
	// 创建金币历史
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
	punch, err := GetUserPunchByTitle(id, title)
	if err == nil {
		return "该标签已选择", nil
	}
	punch.StudentID = id
	punch.Title = title
	return "", DB.Create(&punch).Error
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

func GetTitleHistory(id string, day int) []TitleHistory {
	var Titles []TitleHistory
	DB.Where("id = ? and day = ?", id, day).Find(&Titles)
	return Titles
}

func UpdatePunchHistory(day int) {
	var his TitleHistory
	DB.Where("").First(&his)
	if his.Day == time.Now().Day() {
		return
	}

	var users []User
	DB.Find(&users)
	for _, user := range users {
		var punches []UsersPunch
		DB.Where("student_id = ? ", user.StudentID).Find(&punches)
		for _, punch := range punches {
			history := TitleHistory{
				StudentID: user.StudentID,
				Title:     punch.Title,
				Day:       time.Now().YearDay(),
			}
			DB.Create(&history)
		}
	}
}
