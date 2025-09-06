package repository

import "github.com/deividmendozatech-stack/wishlist/internal/domain"

type WishlistRepository interface {
	Create(w *domain.Wishlist) error
	FindByUser(userID uint) ([]domain.Wishlist, error)
	Delete(userID, wishlistID uint) error
	AddBook(b *domain.Book) error
	ListBooks(userID, wishlistID uint) ([]domain.Book, error)
	RemoveBook(userID, wishlistID, bookID uint) error
}
