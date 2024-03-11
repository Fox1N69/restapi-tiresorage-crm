package database

import (
	"crud-crm/database/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Taipei"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}

	DB.AutoMigrate(&models.Clients{})

	return DB
}
