package database

import (
	"crud-crm/database/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "user=postgres password=8008 dbname=deplom-makar port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Connect database ", err)
	}

	DB.AutoMigrate(&models.Clients{})
	logrus.Println("Database migrate...")

	return DB
}
