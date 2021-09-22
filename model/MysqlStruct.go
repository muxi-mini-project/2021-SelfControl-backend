package model

type User struct {
	StudentID       string `json:"student_id"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	UserPicture     string `json:"user_picture"`
	Gold            int    `json:"gold"`
	Privacy         int    `json:"privacy"`
	CurrentBackdrop int    `json:"current_backdrop"`
}

type GoldHistory struct {
	StudentID      string `json:"student_id"`
	Time           string `json:"time"`
	ChangeNumber   int    `json:"change_number"`
	ResidualNumber int    `json:"residual_number"`
	Reason         string `json:"reason"`
}

type Backdrop struct {
	BackdropID int    `json:"backdrop_id"`
	PictureUrl string `json:"picture_url"`
	Price      int    `json:"price"`
}

type PunchHistory struct {
	ID        int    `json:"id"`
	StudentID string `json:"student_id"`
	Title     string `json:"title"`
	Time      string `json:"time"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
}

type PunchContent struct {
	ID      int    `json:"id"`
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

type UsersBackdrop struct {
	ID         int    `json:"id"`
	StudentID  string `json:"student_id"`
	BackdropID int    `json:"backdrop_id"`
}

type WeekList struct {
	ID        int    `json:"id"`
	Ranking   int    `json:"ranking"`
	StudentID string `json:"student_id"`
	Number    int    `json:"number"`
	Day       int    `json:"day"`
}

type MonthList struct {
	ID        int    `json:"id"`
	Ranking   int    `json:"ranking"`
	StudentID string `json:"student_id"`
	Number    int    `json:"number"`
	Month     int    `json:"month"`
}

// Type 1为周 2为月
type ListHistory struct {
	ID        int    `json:"id"`
	StudentID string `json:"student_id"`
	Type      int    `json:"type"`
	Former    int    `json:"former"`
	After     int    `json:"after"`
}

type TitleHistory struct {
	ID        int    `json:"id"`
	StudentID string `json:"student_id"`
	Title     string `json:"title"`
	Day       int    `json:"day"`
}

type ChangeListRecord struct {
	ID        int    `json:"id"`
	Type      int    `json:"type"`
	StudentID string `json:"student_id"`
	Day       int    `json:"day"`
	Ranking   int    `json:"ranking"`
}
