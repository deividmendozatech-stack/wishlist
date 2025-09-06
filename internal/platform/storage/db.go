package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection(path string) (*gorm.DB, error) {
	if path == "" {
		path = "wishlist.db"
	}
	return gorm.Open(sqlite.Open(path), &gorm.Config{})
}
