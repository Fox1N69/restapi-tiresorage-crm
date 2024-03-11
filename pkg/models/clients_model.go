package models

type Client struct {
	ID          uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	FIO         string `json:"fio"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}
