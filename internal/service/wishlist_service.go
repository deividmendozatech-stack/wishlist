package service

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
)

type WishlistUsecase interface {
	Create(userID uint, name string) error
	List(userID uint) ([]domain.Wishlist, error)
	Delete(userID, id uint) error
}

type WishlistService struct {
	repo repository.WishlistRepository
}

func NewWishlistService(r repository.WishlistRepository) *WishlistService {
	return &WishlistService{repo: r}
}

func (s *WishlistService) Create(userID uint, name string) error {
	wl := &domain.Wishlist{
		UserID: userID,
		Name:   name,
	}
	return s.repo.Create(wl)
}

func (s *WishlistService) List(userID uint) ([]domain.Wishlist, error) {
	return s.repo.FindByUser(userID)
}

func (s *WishlistService) Delete(userID, id uint) error {
	return s.repo.Delete(userID, id)
}
