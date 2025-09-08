package gormrepo

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
	"gorm.io/gorm"
)

// userGormRepo is a GORM-based UserRepository.
type userGormRepo struct {
	db *gorm.DB
}

// NewUserRepo returns a UserRepository using GORM.
func NewUserRepo(db *gorm.DB) repository.UserRepository {
	return &userGormRepo{db: db}
}

// Create stores a new user.
func (r *userGormRepo) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

// FindByUsername fetches a user by username.
func (r *userGormRepo) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll returns all users.
func (r *userGormRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
