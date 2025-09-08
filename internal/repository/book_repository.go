package repository

import "github.com/deividmendozatech-stack/wishlist/internal/domain"

// BookRepository define las operaciones de persistencia para Book
type BookRepository interface {
	Create(b *domain.Book) error
	ListByWishlist(wishlistID uint) ([]domain.Book, error)
	Delete(wishlistID, bookID uint) error
}
