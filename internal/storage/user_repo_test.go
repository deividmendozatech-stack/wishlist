package storage

import (
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates an in-memory SQLite database for testing
// and applies migrations for the User model.
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to sqlite: %v", err)
	}
	// Migrate User model
	if err := db.AutoMigrate(&service.User{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

// TestUserRepo_AddAndList verifies that a user can be added
// and later retrieved from the repository.
func TestUserRepo_AddAndList(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepo(db)

	// Create user
	user := &service.User{Username: "david", Password: "12345"}
	err := repo.Add(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	// List users
	users, err := repo.List()
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "david", users[0].Username)
}
