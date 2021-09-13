package punch

import (
	"SC/model"
	"time"
)

func UpdatePunchHistoryEveryDay() {
	for { // 每日更新用户标签历史
		hour := time.Now().Hour()
		minute := time.Now().Minute()
		time.Sleep(1000000000 * 60)
		if hour == 23 && minute > 55 {
			model.UpdatePunchHistory(time.Now().Day())
		}
	}
}
