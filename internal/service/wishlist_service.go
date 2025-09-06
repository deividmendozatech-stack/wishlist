package service

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
)

type WishlistService struct {
	repo repository.WishlistRepository
}

func NewWishlistService(r repository.WishlistRepository) *WishlistService {
	return &WishlistService{repo: r}
}

func (s *WishlistService) Create(userID uint, name string) error {
	w := &domain.Wishlist{Name: name, UserID: userID}
	return s.repo.Create(w)
}

func (s *WishlistService) List(userID uint) ([]domain.Wishlist, error) {
	return s.repo.FindByUser(userID)
}

func (s *WishlistService) Delete(userID, id uint) error {
	return s.repo.Delete(userID, id)
}

// AddBook, ListBooks, RemoveBook ... se implementan igual
