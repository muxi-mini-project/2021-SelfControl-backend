package punch

import (
	"SC/model"
	"time"
)

func UpdatePunchHistoryEveryDay() {
	for { // 每日更新用户标签历史
		day := time.Now().YearDay()
		// hour := time.Now().Hour()
		// minute := time.Now().Minute()
		time.Sleep(1000000000 * 60) // sleep 1 minute
		if day+1 == time.Now().YearDay() {
			model.UpdatePunchHistory(day)
		}
		// if hour == 23 && minute > 55 {
		// 	model.UpdatePunchHistory(time.Now().YearDay())
		// }
	}
}
