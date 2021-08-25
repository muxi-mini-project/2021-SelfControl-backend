package list

import (
	"SC/model"
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
	err := model.DB.Where("month < ? ", int(time.Now().Month())).First(&Rank).Error
	if err == nil {
		model.DB.Delete(Rank, "month < ? ", int(time.Now().Month()))
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
				StudentID: num.StudentId,
				Ranking:   r,
				Month:     int(time.Now().Month()),
				Number:    num.Number,
			}
			model.CreateMonthlist(&Rank)
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
		var rank model.UserRanking
		rank.Number = ran.Number
		rank.Ranking = ran.Ranking
		rank.StudentId = ran.StudentID
		rank.Name = u.Name
		rank.UserPicture = u.UserPicture
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
	if err := model.DB.Where("day <= ? ", time.Now().YearDay()-7).First(&Rank).Error; err == nil {
		model.DB.Delete(Rank, "day <= ? ", time.Now().YearDay()-7)
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
				StudentID: num.StudentId,
				Ranking:   r,
				Day:       time.Now().YearDay(),
				Number:    num.Number,
			}
			model.DB.Create(&Rank)
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
		var rank model.UserRanking
		rank.Number = ran.Number
		rank.Ranking = ran.Ranking
		rank.StudentId = ran.StudentID
		rank.Name = u.Name
		rank.UserPicture = u.UserPicture
		ranks = append(ranks, rank)
	}
	return ranks, ""
}

func getOrder(s []string) []model.UserAndNumber {
	var Numbers []model.UserAndNumber
	Number := model.UserAndNumber{StudentId: s[0], Number: 1}
	Numbers = append(Numbers, Number)
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(Numbers); j++ {
			if Numbers[j].StudentId == s[i] {
				Numbers[j].Number++
				break
			}
			if j == len(Numbers)-1 {
				Number := model.UserAndNumber{StudentId: s[i]}
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

	// 修改排行榜
	rank := model.WeekList{
		StudentID: id,
		Ranking:   former - ranking,
		Day:       time.Now().YearDay(),
		Number:    number,
	}
	err = ChangeWeekList(rank)
	return "", err
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
	// 修改排行榜
	rank := model.MonthList{
		StudentID: id,
		Ranking:   former - ranking,
		Month:     int(time.Now().Month()),
		Number:    number,
	}
	err = ChangeMonthList(rank)
	return "", err
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
	err := model.DB.Create(&rank).Error
	return err
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
	err := model.DB.Create(&rank).Error
	return err
}
