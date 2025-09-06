package gormrepo

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
	"gorm.io/gorm"
)

type wishlistGormRepo struct {
	db *gorm.DB
}

func NewWishlistRepo(db *gorm.DB) repository.WishlistRepository {
	return &wishlistGormRepo{db: db}
}

func (r *wishlistGormRepo) Create(w *domain.Wishlist) error {
	return r.db.Create(w).Error
}

func (r *wishlistGormRepo) FindByUser(userID uint) ([]domain.Wishlist, error) {
	var lists []domain.Wishlist
	err := r.db.Where("user_id = ?", userID).Preload("Books").Find(&lists).Error
	return lists, err
}

func (r *wishlistGormRepo) Delete(userID, wishlistID uint) error {
	return r.db.Where("id = ? AND user_id = ?", wishlistID, userID).
		Delete(&domain.Wishlist{}).Error
}

// AddBook, ListBooks, RemoveBook ... (implementar igual que arriba)
func (r *wishlistGormRepo) AddBook(b *domain.Book) error {
	return r.db.Create(b).Error
}
func (r *wishlistGormRepo) ListBooks(userID, wishlistID uint) ([]domain.Book, error) {
	var books []domain.Book
	err := r.db.Where("wishlist_id = ?", wishlistID).Find(&books).Error
	return books, err
}
func (r *wishlistGormRepo) RemoveBook(userID, wishlistID, bookID uint) error {
	return r.db.Where("id = ? AND wishlist_id = ?", bookID, wishlistID).
		Delete(&domain.Book{}).Error
}
