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
