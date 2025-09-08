package service

import (
	"errors"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
)

type UserUsecase interface {
	Register(username, password string) error
	List() ([]domain.User, error) // 👈 NUEVO
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserUsecase {
	return &userService{repo: r}
}

func (s *userService) Register(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password required")
	}
	u := domain.User{Username: username, Password: password}
	return s.repo.Create(&u)
}

func (s *userService) List() ([]domain.User, error) {
	return s.repo.FindAll()
}
