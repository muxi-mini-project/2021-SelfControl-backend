package list

import (
	"SC/model"
	"fmt"
	"strconv"
	"time"
)

func GetMonthList() ([]model.UserRanking, string) {
	var (
		ranks []model.UserRanking
		Ranks []model.MonthList
		s     []string
		Rank  model.MonthList
	)
	model.DB.Delete(Rank) // 删除月排行

	PunchHistory := model.GetPunchHistoriesByMonth(int(time.Now().Month()))
	for _, ph := range PunchHistory {
		s = append(s, ph.StudentID)
	}
	UserNumbers := getOrder(s)
	r := 1
	for i, num := range UserNumbers {
		if i > 0 && num.Number < UserNumbers[i-1].Number {
			r++
		}
		Rank = model.MonthList{
			StudentID: num.StudentID,
			Ranking:   r,
			Month:     int(time.Now().Month()),
			Number:    num.Number,
		}
		model.CreateMonthlist(&Rank)
	}
	// 将今日的兑换排名修改进list
	records := model.GetChangeListRecords(time.Now().YearDay())
	for _, record := range records {
		// 修改排行榜
		if record.Type == 1 {
			continue
		}
		var rank model.MonthList
		model.DB.Where("student_id = ? ", record.StudentID).First(&rank)
		Rank := model.MonthList{
			StudentID: record.StudentID,
			Ranking:   rank.Ranking - record.Ranking,
			Month:     int(time.Now().Month()),
			Number:    rank.Number,
		}

		err := ChangeMonthList(Rank)
		if err != nil {
			fmt.Println(err)
		}
	}

	// 把排名前10的加进来
	for i := 1; i <= 10; i++ {
		Rank := model.GetMonthRanks(i)
		Ranks = append(Ranks, Rank...)
	}
	for _, ran := range Ranks {
		u, err := model.GetUserInfo(ran.StudentID)
		if err != nil {
			return nil, "获取用户信息错误"
		}
		rank := model.UserRanking{
			StudentID:   ran.StudentID,
			Name:        u.Name,
			Number:      ran.Number,
			Ranking:     ran.Ranking,
			UserPicture: u.UserPicture,
		}
		ranks = append(ranks, rank)
	}
	return ranks, ""
}

func GetWeekList() ([]model.UserRanking, string) {
	var (
		PunchHistory []model.PunchHistory
		ranks        []model.UserRanking
		Ranks        []model.WeekList
		s            []string
		Rank         model.WeekList
	)
	model.DB.Delete(Rank) // 把排名删除
	model.DB.Table("punch_histories").Select("student_id").Where("day >= ?", int(time.Now().YearDay())-7).Scan(&PunchHistory)
	for _, ph := range PunchHistory {
		s = append(s, ph.StudentID)
	}
	UserNumbers := getOrder(s)
	r := 1
	for i, num := range UserNumbers {
		if i > 0 && num.Number < UserNumbers[i-1].Number {
			r++
		}
		Rank := model.WeekList{
			StudentID: num.StudentID,
			Ranking:   r,
			Day:       time.Now().YearDay(),
			Number:    num.Number,
		}
		model.DB.Create(&Rank)
	}
	// 将今日的兑换排名修改进list
	records := model.GetChangeListRecords(time.Now().YearDay())
	for _, record := range records {
		// 修改排行榜
		if record.Type == 2 {
			continue
		}
		var rank model.WeekList
		model.DB.Where("student_id = ? ", record.StudentID).First(&rank)
		Rank := model.WeekList{
			StudentID: record.StudentID,
			Ranking:   rank.Ranking - record.Ranking,
			Day:       time.Now().YearDay(),
			Number:    rank.Number,
		}

		if Rank.Ranking < 1 { // 防止在兑换后排名提升导致超出排行榜
			Rank.Ranking = 1
		}
		err := ChangeWeekList(Rank)
		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 1; i <= 10; i++ {
		var Rank []model.WeekList
		model.DB.Where("ranking = ? ", i).Find(&Rank)
		Ranks = append(Ranks, Rank...)
	}
	for _, ran := range Ranks {
		u, err := model.GetUserInfo(ran.StudentID)
		if err != nil {
			return nil, "获取用户信息错误"
		}
		rank := model.UserRanking{
			StudentID:   ran.StudentID,
			Name:        u.Name,
			Number:      ran.Number,
			Ranking:     ran.Ranking,
			UserPicture: u.UserPicture,
		}
		ranks = append(ranks, rank)
	}
	return ranks, ""
}

func getOrder(s []string) []model.UserAndNumber {
	var Numbers []model.UserAndNumber
	Number := model.UserAndNumber{StudentID: s[0], Number: 1}
	Numbers = append(Numbers, Number)
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(Numbers); j++ {
			if Numbers[j].StudentID == s[i] {
				Numbers[j].Number++
				break
			}
			if j == len(Numbers)-1 {
				Number := model.UserAndNumber{StudentID: s[i]}
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

func ChangeWeekRanking(id string, ranking int) (string, error) {
	gold := 48 + ranking*2
	user, err := model.GetUserInfo(id)
	if err != nil {
		return "获取用户信息错误", err
	}
	if user.Gold < gold {
		return "金币不足", nil
	}

	// 创建修改历史
	History := model.ListHistory{
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
		if UandN.StudentID == id {
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
	// 修改用户金币
	user.Gold -= gold
	model.UpdateUserInfo(user)

	// 创建金币历史
	price := gold
	history := model.GoldHistory{
		StudentID:      id,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
		ChangeNumber:   -price,
		ResidualNumber: user.Gold,
		Reason:         "兑换周排名:前进" + strconv.Itoa(ranking) + "名",
	}
	if err := model.CreateGoldAndRankHistory(&history, &History); err != nil {
		return "", err
	}

	// 插入到兑换记录表
	record := model.ChangeListRecord{
		StudentID: id,
		Type:      1, // week
		Ranking:   ranking,
		Day:       time.Now().YearDay(),
	}
	return "", model.DB.Create(&record).Error
}

func ChangeMonthRanking(id string, ranking int) (string, error) {
	gold := 48 + ranking*2
	user, err := model.GetUserInfo(id)
	if err != nil {
		return "获取用户信息错误", err
	}
	if user.Gold < gold {
		return "金币不足", nil
	}

	// 创建修改历史
	History := model.ListHistory{
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
		if UandN.StudentID == id {
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
	// 修改用户金币
	user.Gold -= gold
	model.UpdateUserInfo(user)
	// 创建金币历史
	price := gold
	history := model.GoldHistory{
		StudentID:      id,
		Time:           time.Now().Format("2006-01-02 15:04:05"),
		ChangeNumber:   -price,
		ResidualNumber: user.Gold,
		Reason:         "兑换月排名:前进" + strconv.Itoa(ranking) + "名",
	}
	if err := model.CreateGoldAndRankHistory(&history, &History); err != nil {
		return "", err
	}

	// 插入到兑换记录表
	record := model.ChangeListRecord{
		StudentID: id,
		Type:      2, // week
		Ranking:   ranking,
		Day:       time.Now().YearDay(),
	}
	return "", model.DB.Create(&record).Error
}

func ChangeWeekList(rank model.WeekList) error {
	var r model.WeekList
	model.DB.Where("day <= ? ", rank.Day-7).Delete(&r)
	var ranks []model.WeekList
	if err := model.DB.Where("student_id = ? ", rank.StudentID).First(&r).Error; err != nil {
		model.DB.Model(r).Where("ranking < ? AND ranking >= ? ", rank.Ranking, r.Ranking).Find(&ranks)
	} else {
		model.DB.Model(r).Where("ranking < ? ", rank.Ranking).Find(&ranks)
	}
	// 后面的排名++
	for _, Rank := range ranks {
		Rank.Ranking++
		model.DB.Model(r).Where("student_id = ? ", Rank.StudentID).Update("ranking", Rank.Ranking)
	}

	model.DB.Where("student_id = ? ", rank.StudentID).Delete(&rank)
	return model.DB.Create(&rank).Error
}

func ChangeMonthList(rank model.MonthList) error {
	var r model.MonthList
	model.DB.Where("month != ? ", rank.Month).Delete(&r)
	var ranks []model.MonthList
	if err := model.DB.Where("student_id = ? ", rank.StudentID).First(&r).Error; err != nil {
		model.DB.Model(r).Where("ranking < ? AND ranking >= ? ", rank.Ranking, r.Ranking).Find(&ranks)
	} else {
		model.DB.Model(r).Where("ranking < ? ", rank.Ranking).Find(&ranks)
	}
	// 后面的排名++
	for _, Rank := range ranks {
		Rank.Ranking++
		model.DB.Model(r).Where("student_id = ? ", Rank.StudentID).Update("ranking", Rank.Ranking)
	}
	model.DB.Where("student_id = ? ", rank.StudentID).Delete(&rank)
	return model.DB.Create(&rank).Error
}
