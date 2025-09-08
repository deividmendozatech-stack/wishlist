package service

import (
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/repository"
)

// BookUsecase expone las operaciones de negocio para libros
type BookUsecase interface {
	Add(wishlistID uint, title, author string) error
	List(wishlistID uint) ([]domain.Book, error)
	Delete(wishlistID, bookID uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookUsecase {
	return &bookService{repo: r}
}

func (s *bookService) Add(wishlistID uint, title, author string) error {
	book := &domain.Book{
		Title:      title,
		Author:     author,
		WishlistID: wishlistID,
	}
	return s.repo.Create(book)
}

func (s *bookService) List(wishlistID uint) ([]domain.Book, error) {
	return s.repo.ListByWishlist(wishlistID)
}

func (s *bookService) Delete(wishlistID, bookID uint) error {
	return s.repo.Delete(wishlistID, bookID)
}
