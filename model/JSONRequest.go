package model

type Title struct {
	Title string `json:"title"`
}

type Ranking struct {
	Ranking int `json:"ranking"`
}

type TitleAndGold struct {
	Gold  int    `json:"gold"`
	Title string `json:"title"`
}

type BackdropID struct {
	BackdropID int `json:"backdrop_id"`
}
