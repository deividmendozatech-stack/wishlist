package service

import (
	"errors"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
)

// UserUsecase defines basic user operations.
type UserUsecase interface {
	Register(username, password string) error
	List() ([]domain.User, error)
}

// userService implements UserUsecase.
type userService struct {
	repo repository.UserRepository
}

// NewUserService returns a UserUsecase backed by the given repository.
func NewUserService(r repository.UserRepository) UserUsecase {
	return &userService{repo: r}
}

// Register validates input and stores a new user.
func (s *userService) Register(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password required")
	}
	u := domain.User{Username: username, Password: password}
	return s.repo.Create(&u)
}

// List fetches all registered users.
func (s *userService) List() ([]domain.User, error) {
	return s.repo.FindAll()
}
