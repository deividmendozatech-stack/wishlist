package storage

import (
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupWishlistTestDB creates an in-memory SQLite database
// and runs migrations for the Wishlist model.
//
// Params:
//   - t: testing context
//
// Returns:
//   - *gorm.DB: a GORM database connection
func setupWishlistTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}
	if err := db.AutoMigrate(&service.Wishlist{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

// TestWishlistRepo_AddListDelete verifies that the WishlistRepo correctly
// performs Add, List, and Delete operations against the database.
func TestWishlistRepo_AddListDelete(t *testing.T) {
	db := setupWishlistTestDB(t)
	repo := NewWishlistRepo(db)

	// Add a new wishlist
	w := &service.Wishlist{UserID: 1, Name: "Mi lista"}
	err := repo.Add(w)
	assert.NoError(t, err)

	// List wishlists for user 1
	lists, err := repo.List(1)
	assert.NoError(t, err)
	assert.Len(t, lists, 1)
	assert.Equal(t, "Mi lista", lists[0].Name)

	// Delete the wishlist
	err = repo.Delete(1, lists[0].ID)
	assert.NoError(t, err)

	// Verify that the list is now empty
	lists, err = repo.List(1)
	assert.NoError(t, err)
	assert.Empty(t, lists)
}
