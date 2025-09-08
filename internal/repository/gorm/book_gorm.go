package gormrepo

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
	"gorm.io/gorm"
)

type bookGorm struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) repository.BookRepository {
	return &bookGorm{db: db}
}

func (r *bookGorm) Create(b *domain.Book) error {
	return r.db.Create(b).Error
}

func (r *bookGorm) ListByWishlist(wishlistID uint) ([]domain.Book, error) {
	var books []domain.Book
	if err := r.db.Where("wishlist_id = ?", wishlistID).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookGorm) Delete(wishlistID, bookID uint) error {
	return r.db.Where("wishlist_id = ? AND id = ?", wishlistID, bookID).
		Delete(&domain.Book{}).Error
}
