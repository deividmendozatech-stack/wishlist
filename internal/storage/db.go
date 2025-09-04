package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("wishlist.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error conectando a SQLite:", err)
	}
}
