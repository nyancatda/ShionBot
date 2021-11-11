package utils

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"xyz.nyan/MediaWiki-Bot/src/Struct"
)

func SQLLiteLink() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&Struct.UserInfo{})
	return db
}