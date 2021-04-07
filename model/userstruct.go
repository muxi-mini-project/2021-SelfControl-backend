package model

//User结构体在MysqlStruct.go

type UserHomePage struct {
	Name        string `json:"name"`
	UserPicture string `json:"user_picture"`
}

type Choice struct {
	Choice bool `json:"choice"`
}

type Punch struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
	ID     int    `json:"id"`
}

type Punch2 struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
}

type Privacy struct {
	Privacy int `json:"privacy"`
}

type UserAndNumber struct {
	StudentId string `json:"student_id"`
	Number    int    `json:"number"`
}

type BackdropRes struct {
	B1 int `json:"b_1"`
	B2 int `json:"b_2"`
	B3 int `json:"b_3"`
	B4 int `json:"b_4"`
	B5 int `json:"b_5"`
}

type UserRanking struct {
	StudentId string `json:"student_id"`
	Name      string `json:"name"`
	Number    int    `json:"number"`
	Ranking   int    `json:"ranking"`
}

type ListHistories struct {
	StudentID   string `json:"student_id"`
	WeekFormer  int    `json:"week_former"`
	WeekAfter   int    `json:"week_after"`
	MonthFormer int    `json:"month_former"`
	MonthAfter  int    `json:"month_after"`
}
