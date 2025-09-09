package storage

import (
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupBookTestDB creates an in-memory SQLite database for testing
// and applies migrations for the Book model.
func setupBookTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&service.Book{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

// TestBookRepo_AddListDelete verifies the lifecycle of a Book in the repository:
// - Adding a new book
// - Listing books by wishlist ID
// - Deleting a book by ID
// - Ensuring the list is empty after deletion
func TestBookRepo_AddListDelete(t *testing.T) {
	db := setupBookTestDB(t)
	repo := NewBookRepo(db)

	// Add book
	b := &service.Book{WishlistID: 1, Title: "Go 101", Author: "Unknown"}
	err := repo.Add(b)
	assert.NoError(t, err)

	// List books
	books, err := repo.List(1)
	assert.NoError(t, err)
	assert.Len(t, books, 1)
	assert.Equal(t, "Go 101", books[0].Title)

	// Delete book
	err = repo.Delete(1, books[0].ID)
	assert.NoError(t, err)

	// Ensure empty list after deletion
	books, err = repo.List(1)
	assert.NoError(t, err)
	assert.Empty(t, books)
}
