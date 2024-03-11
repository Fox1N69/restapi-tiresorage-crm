package models

import "crud-crm/database"


type Clients struct {
	ID          uint   `json:"id"`
	FIO         string `json:"fio"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

func AutoMigrate() {
	database.DB.AutoMigrate(&Clients{})
}
