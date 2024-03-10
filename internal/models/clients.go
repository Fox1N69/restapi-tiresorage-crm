package models

type Clients struct {
	ID          uint   `json:"id"`
	Fio         string `json:"fio"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
