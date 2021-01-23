package model

//User 结构体
type User struct {
	StudentID   string `json:"student_id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	UserPicture string `json:"user_picture"`
	Gold        int    `json:"gold"`
}
