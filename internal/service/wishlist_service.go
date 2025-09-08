package service

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
)

// WishlistUsecase defines core wishlist operations.
type WishlistUsecase interface {
	Create(userID uint, name string) error
	List(userID uint) ([]domain.Wishlist, error)
	Delete(userID, id uint) error
}

// WishlistService implements WishlistUsecase.
type WishlistService struct {
	repo repository.WishlistRepository
}

// NewWishlistService returns a WishlistService using the provided repository.
func NewWishlistService(r repository.WishlistRepository) *WishlistService {
	return &WishlistService{repo: r}
}

// Create adds a new wishlist for the given user.
func (s *WishlistService) Create(userID uint, name string) error {
	wl := &domain.Wishlist{UserID: userID, Name: name}
	return s.repo.Create(wl)
}

// List fetches all wishlists for a user.
func (s *WishlistService) List(userID uint) ([]domain.Wishlist, error) {
	return s.repo.FindByUser(userID)
}

// Delete removes a wishlist by user and ID.
func (s *WishlistService) Delete(userID, id uint) error {
	return s.repo.Delete(userID, id)
}
