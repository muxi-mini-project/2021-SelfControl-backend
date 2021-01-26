package model

type Punchs struct {
	Punchs []Punch `json:"punchs"`
}

type Punchs2 struct {
	Punchs []Punch2 `json:"punchs"`
}

type UserAndNumbers struct {
	UserAndNumber []UserAndNumber `json:"user_and_number"`
}

type ListPrices struct {
	ListPrice []ListPrice `json:"list_price"`
}

type Backdrops struct {
	Backdrop []Backdrop `json:"backdrop"`
}
