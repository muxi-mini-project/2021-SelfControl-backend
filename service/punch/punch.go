package punch

import (
	"SC/model"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Punch 为 Title 与 Number 的结构体类型
func GetPunchAndNumber(id string) []model.Punch {
	punches := model.GetUserPunches(id)
	var punchs2 []model.Punch
	var Punch model.Punch
	for _, punch := range punches {
		p := model.GetPunchContentByTitle(punch.Title)
		Punch.ID = p.ID
		Punch.Title = punch.Title
		Punch.Number = punch.Number
		punchs2 = append(punchs2, Punch)
	}
	return punchs2
}

type Punch3 struct {
	// model.Punch2
	Title string
	ID    int
	Ok    bool `json:"ok"`
}

func DayPunch(id string, day int) []Punch3 {
	var Punchs []Punch3
	if day == time.Now().YearDay() {
		userPunchs := model.GetUserPunches(id)
		for _, punch := range userPunchs {
			P := model.GetPunchContentByTitle(punch.Title)
			Punch := Punch3{
				punch.Title,
				P.ID,
				false,
			}
			history, _ := model.GetDayPunchHistory(id, punch.Title, day)
			if history.StudentID == "" {
				Punch.Ok = false
			} else {
				Punch.Ok = true
			}
			Punchs = append(Punchs, Punch)
		}
	} else {
		punchs := model.GetTitleHistory(id, day)

		for _, punch := range punchs {
			P := model.GetPunchContentByTitle(punch.Title)
			Punch := Punch3{
				punch.Title,
				P.ID,
				false,
			}
			history, _ := model.GetDayPunchHistory(id, punch.Title, day)
			fmt.Println(history.StudentID)
			if history.StudentID == "" {
				Punch.Ok = false
			} else {
				Punch.Ok = true
			}
			Punchs = append(Punchs, Punch)
		}
	}
	return Punchs
}

func DayPunches(id string, day int) int {
	histories := model.GetUserPunchHistoriesByDay(id, day)
	var Len int
	if day == time.Now().YearDay() {
		Len = len(model.GetUserPunches(id))
	} else {
		Len = len(model.GetTitleHistory(id, day))
	}

	// 无打卡信息
	if Len == 0 {
		return 0
	}
	// 未全部完成
	if Len > len(histories) {
		return -1
	}
	// 返回全部完成的打卡数量
	return Len
}

// 获取某日该用户已完成打卡
func GetDayPunches(id string, day int) []model.Punch {
	histories := model.GetUserPunchHistoriesByDay(id, day)
	var punchs2 []model.Punch

	for _, history := range histories {
		Punch := model.Punch{
			Title:  history.Title,
			ID:     model.GetPunchContentByTitle(history.Title).ID,
			Number: len(model.GetUserPunchHistoriesByTitle(id, history.Title)),
		}
		punchs2 = append(punchs2, Punch)
	}
	return punchs2
}

func GetWeekPunchs(id string, month int) []int {
	histories := model.GetUserPunchHistoriesByMonth(id, month)
	var days, Nums []int
	var nums [6]int
	tag := 0
	days = append(days, 0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365)
	// w := (20/4 - 2*20 + 21 + 21/4 + 13*(month+1)/5) % 7
	w := (-9 + 13*(month+1)/5) % 7
	day := days[month] - days[month-1]
	// 一个月6周 tag 就为1
	if day+w > 35 {
		tag = 1
	}
	for _, history := range histories {
		history.Day = history.Day - days[month-1]
		if history.Day <= 8-w {
			nums[0]++
			continue
		} else if history.Day <= 15-w {
			nums[1]++
			continue
		} else if history.Day <= 22-w {
			nums[2]++
			continue
		} else if history.Day <= 29-w {
			nums[3]++
			continue
		} else if history.Day <= 36-w {
			nums[4]++
			continue
		} else {
			nums[5]++
			continue
		}
	}
	if tag == 1 {
		for _, num := range nums {
			Nums = append(Nums, num)
		}
	} else {
		for i := 0; i < len(nums)-1; i++ {
			Nums = append(Nums, nums[i])
		}
	}
	return Nums
}

func CompletePunch(id string, title string) error {
	punches := model.GetUserPunches(id) // 获取该用户该title的打卡情况(包括总数量)

	// 判断是否已选
	Punch, err := model.GetUserPunchByTitle(id, title)
	if err != nil {
		return errors.New("未选择该标签")
	}

	// 判断是否已打卡
	Histories := model.GetUserPunchHistoriesByDay(id, time.Now().YearDay())
	for _, History := range Histories {
		if History.Title == title { // 已打卡
			return errors.New("今日已打该卡")
		}
	}

	// 允许打卡后创建打卡历史
	history := model.PunchHistory{
		Title:     title,
		StudentID: id,
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		Month:     int(time.Now().Month()),
		Day:       time.Now().YearDay(),
	}
	if err := model.CreatePunchHistory(&history); err != nil {
		return err
	}

	Punch.Number += 1

	if err := model.UpdateUserPunch(Punch); err != nil {
		return err
	}

	puns := GetDayPunches(id, time.Now().YearDay()) // 今日已完成
	if len(puns) == len(punches) {
		var gold int
		if len(puns) <= 5 {
			gold = len(puns) * 10
		} else {
			gold = ((len(puns)-5)*2 + 10) * len(puns)
		}
		return completeAllPunch(id, gold)
	}
	return nil
}

func completeAllPunch(id string, gold int) error {
	// 修改用户金币
	user, err := model.GetUserInfo(id)
	if err != nil {
		return err
	}
	user.Gold += gold
	if err := model.UpdateUserInfo(user); err != nil {
		return err
	}
	s := strconv.Itoa(gold)
	// 创建金币历史
	history := model.GoldHistory{
		StudentID:      id,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
		ChangeNumber:   gold,
		ResidualNumber: user.Gold,
		Reason:         "完成今日打卡+" + s + "金币",
	}
	err = model.CreateGoldHistory(&history)
	return err
}

func DeletePunch(id string, title string) (string, error) {
	u, err := model.GetUserPunchByTitle(id, title)
	if err != nil {
		return "用户未选择该标签", nil
	}
	return "", model.DeletePunch(&u)
}

func GetMonthly(id string) []model.Punch {
	punchs := model.GetUserPunchHistoriesByMonth(id, int(time.Now().Month()-1))
	var Punchs []model.Punch
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
			var Punch model.Punch
			Punch.Title = punch.Title
			Punch.Number++
			Punchs = append(Punchs, Punch)
		}
	}
	return Punchs
}
