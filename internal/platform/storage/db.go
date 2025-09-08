package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewConnection opens a SQLite DB with GORM.
// Uses "wishlist.db" if path is empty.
func NewConnection(path string) (*gorm.DB, error) {
	if path == "" {
		path = "wishlist.db"
	}
	return gorm.Open(sqlite.Open(path), &gorm.Config{})
}
