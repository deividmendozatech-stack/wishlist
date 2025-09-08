package repository

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"gorm.io/gorm"
)

// UserRepository defines the contract for user persistence.
type UserRepository interface {
	Create(*domain.User) error
	FindByUsername(username string) (*domain.User, error)
	FindAll() ([]domain.User, error)
}

// userRepo is a GORM-based implementation of UserRepository.
type userRepo struct {
	db *gorm.DB
}

// NewUserRepo returns a GORM-backed UserRepository.
func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// Create saves a new user record.
func (r *userRepo) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

// FindByUsername retrieves a user by username.
func (r *userRepo) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll returns all stored users.
func (r *userRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
