package model

//User结构体在MysqlStruct.go

type UserHomePage struct {
	Name        string `json:"student_id"`
	UserPicture string `json:"name"`
}

type Gold struct {
	Gold int `json:"gold"`
}

type Choice struct {
	Choice bool `json:"choice"`
}
type Achievements struct {
	Achievements []string `json:"achievement"`
}

type Punch struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

type Punch2 struct {
	Title string `json:"title"`
}

type Privacy struct {
	Privacy bool `json:"privacy"`
}

type UserAndNumber struct {
	StudentId string `json:"student_id"`
	Number    int    `json:"number"`
}
