package Utils

import (
	"log"

	"github.com/nyancatda/ShionBot/src/Struct"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLLiteLink() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/database.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&Struct.UserInfo{})
	return db
}
