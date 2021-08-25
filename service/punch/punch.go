package punch

import (
	"SC/model"
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

func TodayPunch(StudentId string, TitleID int) bool {
	Punch := model.GetPunchContentById(TitleID)
	today := time.Now().YearDay()
	_, err := model.GetTodayPunchHistory(StudentId, Punch.Title, today)
	var choice bool
	if err != nil {
		choice = false
	} else {
		choice = true
	}
	return choice
}

func TodayPunches(id string) int {
	histories := model.GetUserPunchHistoriesByDay(id, time.Now().Day())
	Punchs := model.GetUserPunches(id)
	// 无打卡信息
	if len(Punchs) == 0 {
		return 0
	} // 未全部完成
	if len(Punchs) > len(histories) {
		return -1
	} // 返回全部完成的打卡数量
	return len(Punchs)
}

func GetDayPunches(StudentId string, day int) []model.Punch {
	histories := model.GetUserPunchHistoriesByDay(StudentId, day)
	var punchs2 []model.Punch
	var Punch model.Punch

	for _, history := range histories {
		p := model.GetPunchContentByTitle(history.Title)
		Punch.ID = p.ID
		Punch.Title = history.Title
		punchs2 = append(punchs2, Punch)

	}
	return punchs2
}

func GetWeekPunchs(id string, month int) []int {
	histories := model.GetPunchHistoriesByMonth(month)
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
	pun, err := model.GetUserPunchByTitle(id, title)
	if err != nil {
		return err
	}
	var punch model.PunchHistory
	punch.Title = title
	punch.StudentID = id
	punch.Time = time.Now().Format("2006-01-02 15:04:05")
	punch.Month = int(time.Now().Month())
	punch.Day = time.Now().YearDay()
	if err := model.CreatePunchHistory(&punch); err != nil {
		return err
	}
	pun.Number += 1
	if err := model.UpdateUserPunch(pun); err != nil {
		return err
	}

	punches := model.GetUserPunches(id)
	puns := GetDayPunches(id, time.Now().Day())
	if len(puns) == len(punches) {
		gold := 0
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
