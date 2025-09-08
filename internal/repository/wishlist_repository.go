package repository

import "github.com/deividmendozatech-stack/wishlist/internal/domain"

// WishlistRepository defines basic operations for persisting wishlists.
type WishlistRepository interface {
	// Create saves a new wishlist.
	Create(w *domain.Wishlist) error

	// FindByUser returns all wishlists belonging to a user.
	FindByUser(userID uint) ([]domain.Wishlist, error)

	// Delete removes a wishlist by user and ID.
	Delete(userID, id uint) error
}
