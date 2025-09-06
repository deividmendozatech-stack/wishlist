package gormrepo

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
	"gorm.io/gorm"
)

type wishlistGorm struct {
	db *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) repository.WishlistRepository {
	return &wishlistGorm{db: db}
}

func (r *wishlistGorm) Create(w *domain.Wishlist) error {
	return r.db.Create(w).Error
}

func (r *wishlistGorm) FindByUser(userID uint) ([]domain.Wishlist, error) {
	var lists []domain.Wishlist
	if err := r.db.Where("user_id = ?", userID).Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *wishlistGorm) Delete(userID, id uint) error {
	return r.db.Where("user_id = ? AND id = ?", userID, id).Delete(&domain.Wishlist{}).Error
}
