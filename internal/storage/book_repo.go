package storage

import (
	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"gorm.io/gorm"
)

// BookRepo is the GORM-based implementation of the BookRepository interface.
// It handles persistence of books in a relational database.
type BookRepo struct {
	db *gorm.DB
}

// NewBookRepo creates a new BookRepo instance using the provided GORM DB connection.
func NewBookRepo(db *gorm.DB) service.BookRepository {
	return &BookRepo{db: db}
}

// Add inserts a new book into the database.
func (r *BookRepo) Add(b *service.Book) error {
	return r.db.Create(b).Error
}

// List retrieves all books associated with a given wishlist ID.
func (r *BookRepo) List(wishlistID uint) ([]service.Book, error) {
	var books []service.Book
	if err := r.db.Where("wishlist_id = ?", wishlistID).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// Delete removes a book by its ID, ensuring it belongs to the specified wishlist.
func (r *BookRepo) Delete(wishlistID, bookID uint) error {
	return r.db.Where("id = ? AND wishlist_id = ?", bookID, wishlistID).
		Delete(&service.Book{}).Error
}
