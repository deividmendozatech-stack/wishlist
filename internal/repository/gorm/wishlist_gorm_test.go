package gormrepo_test

import (
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	gormrepo "github.com/deividmendozatech-stack/wishlist/internal/repository/gorm"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB opens an in-memory SQLite DB and migrates Wishlist.
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("cannot open sqlite in-memory: %v", err)
	}
	if err = db.AutoMigrate(&domain.Wishlist{}); err != nil {
		t.Fatalf("cannot migrate: %v", err)
	}
	return db
}

// TestWishlistGormIntegration checks create, fetch, and delete for WishlistRepo.
func TestWishlistGormIntegration(t *testing.T) {
	db := setupTestDB(t)
	repo := gormrepo.NewWishlistRepo(db)

	// Create
	err := repo.Create(&domain.Wishlist{UserID: 1, Name: "Lista A"})
	assert.NoError(t, err)

	// FindByUser
	lists, err := repo.FindByUser(1)
	assert.NoError(t, err)
	assert.Len(t, lists, 1)
	assert.Equal(t, "Lista A", lists[0].Name)

	// Delete
	err = repo.Delete(1, lists[0].ID)
	assert.NoError(t, err)

	lists, err = repo.FindByUser(1)
	assert.NoError(t, err)
	assert.Len(t, lists, 0)
}
