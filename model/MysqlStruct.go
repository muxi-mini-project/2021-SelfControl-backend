package model

import "time"

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
}

type Achievement struct {
	StudentID   string `json:"student_id"`
	Achievement string `json:"achievement"`
}

type PunchContent struct {
	Type       string `json:"type"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	PictureUrl string `json:"picture_url"`
}

type UsersPunch struct {
	ID        int    `json:"id"`
	StudentID string `json:"student_id"`
	Title     string `json:"title"`
	Number    int    `json:"number"`
}
