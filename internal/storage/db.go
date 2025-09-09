package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewConnection opens a new SQLite database connection using GORM.
//
// path: the file path of the SQLite database.
//
//	Use ":memory:" to create an in-memory database (useful for testing).
//
// Returns:
//   - *gorm.DB: the GORM database instance
//   - error: an error if the connection fails
func NewConnection(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
