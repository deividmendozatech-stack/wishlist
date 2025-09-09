package storage

import (
	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"gorm.io/gorm"
)

// UserRepo is the GORM-based implementation of the UserRepository interface.
// It provides persistence operations for User entities.
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo creates a new instance of UserRepo.
func NewUserRepo(db *gorm.DB) service.UserRepository {
	return &UserRepo{db: db}
}

// Add inserts a new user into the database.
//
// Params:
//   - u: pointer to a User entity
//
// Returns:
//   - error: if the database operation fails
func (r *UserRepo) Add(u *service.User) error {
	return r.db.Create(u).Error
}

// List retrieves all users from the database.
//
// Returns:
//   - []service.User: slice of all users
//   - error: if the database operation fails
func (r *UserRepo) List() ([]service.User, error) {
	var users []service.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
