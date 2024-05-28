package models

import "time"

type Crequests struct {
	ID           uint      `json:"id"`
	FIO          string    `json:"fio"`
	StorageCell  int       `json:"storage_cell"`
	DiskSize     int       `json:"disk_size"`
	ArrivalDate  time.Time `json:"arrivalDate" gorm:"type:date"`
	DeliveryDate time.Time `json:"deliveryDate" gorm:"type:date"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Price        string    `json:"price"`
	IsPaid       string    `json:"isPaid"`
	Status       string    `json:"status"`
}
