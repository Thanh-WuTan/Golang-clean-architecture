package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"onlyfounds/domain"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = connectDB()
	SyncDB()
	return DB
}

func connectDB() *gorm.DB {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: " + dsn + err.Error())
		return nil
	}
	return db
}

func SyncDB() {
	DB.AutoMigrate(&domain.User{})
}
