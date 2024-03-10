package models

type Test struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Email  string `json:"email"`
	Fio    string `json:"fio"`
}
