package models

type Client struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Branch      string    `json:"branch"`
	FIO         string    `json:"fio"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	LastVisit   string    `json:"lastVisit"`
	CrequestID  uint      `json:"crequest_id"`
	Crequest    Crequests `gorm:"foreignKey:CrequestID"`
}
