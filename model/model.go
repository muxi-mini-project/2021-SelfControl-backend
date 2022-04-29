package model

import (
	"errors"
	"github.com/ShiinaOrez/GoSecurity/security"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

func GetUserInfo(id string) (User, error) {
	var u User
	return u, DB.Where("student_id = ?", id).First(&u).Error
}

func UpdateUserInfo(user User) error {
	return DB.Model(&user).Where("student_id = ?", user.StudentID).Update(user).Error
}

// -----------------------------------------------
// punch:
// 今日该用户已选
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

func GetDayPunchHistory(StudentId string, title string, day int) (PunchHistory, error) {
	var punch PunchHistory
	return punch, DB.Where("student_id = ? AND title = ? AND day = ?", StudentId, title, day).First(&punch).Error
}

func GetUserPunchHistoriesByDay(id string, day int) []PunchHistory {
	var histories []PunchHistory
	DB.Where("student_id = ? AND day = ? ", id, day).Find(&histories)
	return histories
}

func GetPunchHistoriesByMonth(month int) []PunchHistory {
	var histories []PunchHistory
	DB.Where("month = ? ", month).Find(&histories)
	return histories
}

func GetUserPunchByTitle(id string, title string) (UsersPunch, error) {
	var pun UsersPunch
	return pun, DB.Where("student_id = ? AND title = ? ", id, title).First(&pun).Error
}

func GetUserPunchHistoriesByTitle(id string, title string) []PunchHistory {
	var histories []PunchHistory
	DB.Where("student_id = ? AND title = ? ", id, title).Find(&histories)
	return histories
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

	var userBackdrop UsersBackdrop
	DB.Where("backdrop_id = ? and student_id = ? ", BackdropID, id).First(&userBackdrop)
	if userBackdrop.BackdropID == BackdropID {
		return "", errors.New("已购买此背景")
	}
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
	_, err := GetUserPunchByTitle(id, title)
	if err == nil {
		return "该标签已选择", nil
	}

	Punch := UsersPunch{
		StudentID: id,
		Title:     title,
		Number:    len(GetUserPunchHistoriesByTitle(id, title)),
	}

	return "", DB.Create(&Punch).Error
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
	if Type == "week" {
		var u WeekList
		err := DB.Where("student_id = ? ", id).First(&u).Error
		if err != nil {
			return -1
		}
		return u.Ranking
	} else if Type == "month" {
		var u MonthList
		err := DB.Where("student_id = ? ", id).First(&u).Error
		if err != nil {
			return -1
		}
		return u.Ranking
	}
	return -1
}

func GetTitleHistory(id string, day int) []TitleHistory {
	var Titles []TitleHistory
	DB.Where("student_id = ? and day = ?", id, day).Find(&Titles)
	return Titles
}

func UpdatePunchHistory(day int) {
	var his TitleHistory
	if DB.Where("day = ? ", day).First(&his); his.ID != 0 { // 当日的已更新过
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
				Day:       day,
			}
			DB.Create(&history)
		}
	}
}

func GetChangeListRecords(day int) []ChangeListRecord {
	var records []ChangeListRecord
	DB.Where("day = ? ", day).Find(&records)
	return records
}

func GeneratePasswordHash(password string) string {
	return security.GeneratePasswordHash(password)
}

func CheckPassword(password, hashPwd string) bool {
	return security.CheckPasswordHash(password, hashPwd)
}
