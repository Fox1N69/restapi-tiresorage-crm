package models

import (
	"time"

	"gorm.io/gorm"
)

type Crequests struct {
	ID           uint      `json:"id"`
	Branch       string    `json:"branch"`
	FIO          string    `json:"fio"`
	StorageCell  string    `json:"storageCell"`
	DiskSize     string    `json:"diskSize"`
	ArrivalDate  time.Time `json:"arrivalDate" gorm:"type:date"`
	DeliveryDate time.Time `json:"deliveryDate" gorm:"type:date"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Price        string    `json:"price"`
	IsPaid       string    `json:"isPaid"`
	Status       string    `json:"status"`
}

func (c *Crequests) AfterCreate(tx *gorm.DB) (err error) {
	client := Client{
		FIO:        c.FIO,
		Branch:     c.Branch,
		LastVisit:  time.Now().Format("2006-01-02"),
		CrequestID: c.ID,
	}
	if err := tx.Create(&client).Error; err != nil {
		return err
	}
	return nil
}

func (c *Crequests) BeforeDelete(tx *gorm.DB) (err error) {
	if err := tx.Where("crequest_id = ?", c.ID).Delete(&Client{}).Error; err != nil {
		return err
	}
	return nil
}
