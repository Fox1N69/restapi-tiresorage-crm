package models

type Client struct {
	ID          uint   `json:"id"`
	FIO         string `json:"fio"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
