package repository

import "github.com/deividmendozatech-stack/wishlist/internal/domain"

// BookRepository defines the basic persistence methods for Book.
type BookRepository interface {
	// Create saves a new book record.
	Create(b *domain.Book) error

	// ListByWishlist returns all books for a given wishlist ID.
	ListByWishlist(wishlistID uint) ([]domain.Book, error)

	// Delete removes a book by wishlist and book ID.
	Delete(wishlistID, bookID uint) error
}
