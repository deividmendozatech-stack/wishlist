package service

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
)

// BookUsecase defines book-related business operations.
type BookUsecase interface {
	Add(wishlistID uint, title, author string) error
	List(wishlistID uint) ([]domain.Book, error)
	Delete(wishlistID, bookID uint) error
}

// bookService implements BookUsecase using a BookRepository.
type bookService struct {
	repo repository.BookRepository
}

// NewBookService creates a BookUsecase with the given repository.
func NewBookService(r repository.BookRepository) BookUsecase {
	return &bookService{repo: r}
}

// Add creates a new book linked to a wishlist.
func (s *bookService) Add(wishlistID uint, title, author string) error {
	book := &domain.Book{
		Title:      title,
		Author:     author,
		WishlistID: wishlistID,
	}
	return s.repo.Create(book)
}

// List retrieves all books for a wishlist.
func (s *bookService) List(wishlistID uint) ([]domain.Book, error) {
	return s.repo.ListByWishlist(wishlistID)
}

// Delete removes a book from a wishlist.
func (s *bookService) Delete(wishlistID, bookID uint) error {
	return s.repo.Delete(wishlistID, bookID)
}
