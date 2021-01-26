package model

import "time"

type User struct {
	StudentID   string `json:"student_id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	UserPicture string `json:"user_picture"`
	Gold        int    `json:"gold"`
	Privacy     bool   `json:"privacy"`
}

type GoldHistory struct {
	StudentID      string    `json:"student_id"`
	Time           time.Time `json:"time"`
	ChangeNumber   int       `json:"change_number"`
	ResidualNumber int       `json:"residual_number"`
	Reason         string    `json:"reason"`
}

type Backdrop struct {
	BackdropID int    `json:"backdrop_id"`
	PictureUrl string `json:"picture_url"`
	Price      int    `json:"price"`
}

type PunchHistory struct {
	ID        int       `json:"id"`
	StudentID string    `json:"student_id"`
	Title     string    `json:"title"`
	Time      time.Time `json:"time"`
	Day       int       `json:"day"`
	Month     int       `json:"month"`
}

type Achievement struct {
	StudentID   string `json:"student_id"`
	Achievement string `json:"achievement"`
}

type PunchContent struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UsersPunch struct {
	ID        int    `json:"id"`
	StudentID string `json:"student_id"`
	Title     string `json:"title"`
	Number    int    `json:"number"`
}

type ListPrice struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

type UsersBackdrop struct {
	ID         int    `json:"id"`
	StudentID  string `json:"student_id"`
	BackdropID int    `json:"backdrop_id"`
}
