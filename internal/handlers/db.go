package handlers

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	var err error

	dsn := os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") +
		")/" + os.Getenv("DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
}