package storage

import (
	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"gorm.io/gorm"
)

// WishlistRepo is the GORM-based implementation of service.WishlistRepository.
// It provides persistence operations for wishlists.
type WishlistRepo struct {
	db *gorm.DB
}

// NewWishlistRepo creates a new WishlistRepo instance.
//
// Params:
//   - db: the GORM database connection
//
// Returns:
//   - service.WishlistRepository: a repository for wishlists
func NewWishlistRepo(db *gorm.DB) service.WishlistRepository {
	return &WishlistRepo{db: db}
}

// Add inserts a new wishlist into the database.
//
// Params:
//   - w: pointer to a Wishlist entity
//
// Returns:
//   - error: any database error encountered during insertion
func (r *WishlistRepo) Add(w *service.Wishlist) error {
	return r.db.Create(w).Error
}

// List retrieves all wishlists belonging to a given user.
//
// Params:
//   - userID: the ID of the user
//
// Returns:
//   - []service.Wishlist: the list of wishlists
//   - error: any database error encountered
func (r *WishlistRepo) List(userID uint) ([]service.Wishlist, error) {
	var wishlists []service.Wishlist
	if err := r.db.Where("user_id = ?", userID).Find(&wishlists).Error; err != nil {
		return nil, err
	}
	return wishlists, nil
}

// Delete removes a wishlist by its ID and associated user ID.
//
// Params:
//   - userID: the ID of the user who owns the wishlist
//   - wishlistID: the ID of the wishlist to delete
//
// Returns:
//   - error: any database error encountered during deletion
func (r *WishlistRepo) Delete(userID, wishlistID uint) error {
	return r.db.Where("id = ? AND user_id = ?", wishlistID, userID).
		Delete(&service.Wishlist{}).Error
}
