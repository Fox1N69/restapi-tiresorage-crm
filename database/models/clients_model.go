package models


type Clients struct {
	ID          uint   `json:"id"`
	FIO         string `json:"fio"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
