package repository

import "github.com/deividmendozatech-stack/wishlist/internal/domain"

type WishlistRepository interface {
	Create(w *domain.Wishlist) error
	FindByUser(userID uint) ([]domain.Wishlist, error)
	Delete(userID, id uint) error
}
