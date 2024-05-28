package models

type Client struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FIO         string `json:"fio"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	LastVisit   string `json:"lastVisit"`
}
