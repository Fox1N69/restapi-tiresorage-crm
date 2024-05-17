package models

type Crequests struct {
	ID     uint   `json:"id"`
	FIO    string `json:"fio"`
	Data   string `json:"data"`
	Price  string `json:"price"`
	IsPaid string `json:"isPaid"`
	Status string `json:"status"`
}
