package service

//
// ─────────────────────────── SERVICE IMPLEMENTATION ───────────────────────────
//

// bookService implements the BookUsecase interface.
// It contains the business logic for managing books inside wishlists.
type bookService struct {
	repo BookRepository
}

// NewBookService creates a new instance of bookService with the provided repository.
func NewBookService(r BookRepository) BookUsecase {
	return &bookService{repo: r}
}

// Add creates and stores a new book in the given wishlist.
func (s *bookService) Add(wishlistID uint, title, author string) error {
	book := Book{WishlistID: wishlistID, Title: title, Author: author}
	return s.repo.Add(&book)
}

// List retrieves all books associated with a given wishlist ID.
func (s *bookService) List(wishlistID uint) ([]Book, error) {
	return s.repo.List(wishlistID)
}

// Delete removes a book from the repository using its wishlist ID and book ID.
func (s *bookService) Delete(wishlistID, bookID uint) error {
	return s.repo.Delete(wishlistID, bookID)
}
