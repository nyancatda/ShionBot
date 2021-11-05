package utils

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLLiteLink() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}