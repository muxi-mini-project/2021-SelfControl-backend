package user

import "SC/model"

func GetUserTitleByDay(id string, day int) ([]model.Title, error) {
	if _, err := model.GetUserInfo(id); err != nil {
		return nil, err
	}

	Titles := model.GetTitleHistory(id, day)

	var titles []model.Title
	for _, title := range Titles {
		titles = append(titles, model.Title{Title: title.Title})
	}

	return titles, nil
}
