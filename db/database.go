package db

import (
	"time"

	"github.com/Fox1N69/rest-tsc/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "user=postgres password=8008 dbname=deplom-makar port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Connect DB ", err)
	}

	DB = db

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Clients{})
	DB.AutoMigrate(&models.Test{})

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = InitDB()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			logrus.Println("Database is unavaibl. Wait for %d sec.\n")
			time.Sleep(sleep * time.Second)
			DB = InitDB()
		}
	}

	return DB
}
