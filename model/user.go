package model

//User 结构体
type User struct {
	StudentID   string `json:"student_id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	UserPicture string `json:"user_picture"`
	Gold        int    `json:"gold"`
	Privacy     bool   `json:"privacy"`
}

type UserHomePage struct {
	Name        string `json:"student_id"`
	UserPicture string `json:"name"`
}

type Gold struct {
	Gold int `json:"gold"`
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
	Url   string `json:"url"`
}
type Punchs struct {
	Punchs []Punch `json:"punchs"`
}
type Privacy struct {
	Privacy bool `json:"privacy"`
}
