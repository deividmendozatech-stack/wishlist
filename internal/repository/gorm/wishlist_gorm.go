package gormrepo

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
	"gorm.io/gorm"
)

// wishlistGorm is a GORM-based WishlistRepository.
type wishlistGorm struct {
	db *gorm.DB
}

// NewWishlistRepo returns a WishlistRepository using GORM.
func NewWishlistRepo(db *gorm.DB) repository.WishlistRepository {
	return &wishlistGorm{db: db}
}

// Create stores a new wishlist.
func (r *wishlistGorm) Create(w *domain.Wishlist) error {
	return r.db.Create(w).Error
}

// FindByUser gets all wishlists for a user.
func (r *wishlistGorm) FindByUser(userID uint) ([]domain.Wishlist, error) {
	var lists []domain.Wishlist
	if err := r.db.Where("user_id = ?", userID).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

// Delete removes a wishlist by user and ID.
func (r *wishlistGorm) Delete(userID, id uint) error {
	return r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&domain.Wishlist{}).Error
}
