package models


type Crequests struct {
	ID           uint   `json:"id"`
	FIO          string `json:"fio"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Status       bool   `json:"status"`
	StorageData  string `json:"storageData"`
	StoragePrice string `json:"storagePrice"`
}
