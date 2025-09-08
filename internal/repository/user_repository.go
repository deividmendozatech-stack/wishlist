package repository

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"gorm.io/gorm"
)

// UserRepository define los métodos de persistencia de usuarios
type UserRepository interface {
	Create(*domain.User) error
	FindByUsername(username string) (*domain.User, error)
	FindAll() ([]domain.User, error)
}

// userRepo implementa UserRepository
type userRepo struct {
	db *gorm.DB
}

// NewUserRepo crea un nuevo repositorio
func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

// Create inserta un usuario nuevo
func (r *userRepo) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

// FindByUsername busca un usuario por username
func (r *userRepo) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll retorna todos los usuarios
func (r *userRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
