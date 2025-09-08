package gormrepo

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
	"gorm.io/gorm"
)

// userGormRepo implementa repository.UserRepository
type userGormRepo struct {
	db *gorm.DB
}

// NewUserRepo crea un nuevo repositorio basado en GORM
func NewUserRepo(db *gorm.DB) repository.UserRepository {
	return &userGormRepo{db: db}
}

// Create guarda un nuevo usuario
func (r *userGormRepo) Create(u *domain.User) error {
	return r.db.Create(u).Error
}

// FindByUsername busca un usuario por su nombre
func (r *userGormRepo) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll devuelve todos los usuarios registrados
func (r *userGormRepo) FindAll() ([]domain.User, error) {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
