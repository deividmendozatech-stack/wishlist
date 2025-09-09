package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// newConnection (privada)
func newConnection(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitDB (pública) → expuesta para main.go
func InitDB(path string) (*gorm.DB, error) {
	return newConnection(path)
}
