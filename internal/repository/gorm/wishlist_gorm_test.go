package gormrepo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	gormrepo "github.com/deividmendozatech-stack/wishlist/internal/repository/gorm"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("cannot open sqlite in-memory: %v", err)
	}
	// Migramos el modelo que se usará en las pruebas
	err = db.AutoMigrate(&domain.Wishlist{})
	if err != nil {
		t.Fatalf("cannot migrate: %v", err)
	}
	return db
}

func TestWishlistGormIntegration(t *testing.T) {
	db := setupTestDB(t)
	repo := gormrepo.NewWishlistRepo(db)

	// 1. Create
	err := repo.Create(&domain.Wishlist{UserID: 1, Name: "Lista A"})
	assert.NoError(t, err)

	// 2. FindByUser
	lists, err := repo.FindByUser(1)
	assert.NoError(t, err)
	assert.Len(t, lists, 1)
	assert.Equal(t, "Lista A", lists[0].Name)

	// 3. Delete
	err = repo.Delete(1, lists[0].ID)
	assert.NoError(t, err)

	lists, err = repo.FindByUser(1)
	assert.NoError(t, err)
	assert.Len(t, lists, 0)
}
