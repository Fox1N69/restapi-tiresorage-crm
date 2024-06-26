package database

import (
	"crud-crm/pkg/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGormDB() *gorm.DB {
	dsn := "postgres://root:LLBHjXo4ieF2rDjyycbLn9TL2Ywl4esq@dpg-cp3k278l6cac73f78clg-a.frankfurt-postgres.render.com/deplom"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Connect database ", err)
	}
	logrus.Infoln("Database connect...")

	DB.AutoMigrate(&models.Client{}, &models.User{}, &models.Crequests{})
	DB.Exec("ALTER SEQUENCE clients_id_seq RESTART WITH 1")

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = InitGormDB()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			logrus.Infoln("Database is unavaibl. Wait for sec")
			time.Sleep(sleep * time.Second)
			DB = InitGormDB()
		}
	}

	return DB
}
