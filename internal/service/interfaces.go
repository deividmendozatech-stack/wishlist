package service

//
// ─────────────────────────── USE CASE INTERFACES ───────────────────────────
//

// UserUsecase defines the business logic for users.
type UserUsecase interface {
	// Register creates a new user with username and password.
	Register(username, password string) error

	// List retrieves all registered users.
	List() ([]User, error)
}

// WishlistUsecase defines the business logic for wishlists.
type WishlistUsecase interface {
	// Create adds a new wishlist for a given user.
	Create(userID uint, name string) error

	// List retrieves all wishlists for a given user.
	List(userID uint) ([]Wishlist, error)

	// Delete removes a wishlist by its ID for a given user.
	Delete(userID, wishlistID uint) error
}

// BookUsecase defines the business logic for books inside wishlists.
type BookUsecase interface {
	// Add inserts a new book into a wishlist.
	Add(wishlistID uint, title, author string) error

	// List retrieves all books in a given wishlist.
	List(wishlistID uint) ([]Book, error)

	// Delete removes a book by its ID from a wishlist.
	Delete(wishlistID, bookID uint) error
}

// GoogleBooksUsecase defines the contract for searching books via Google Books API.
type GoogleBooksUsecase interface {
	// Search performs a query against the Google Books API
	// and returns a simplified list of results.
	Search(query string) ([]GoogleBook, error)
}

//
// ─────────────────────────── REPOSITORY INTERFACES ───────────────────────────
//

// UserRepository defines persistence operations for users.
type UserRepository interface {
	// Add saves a new user to the database.
	Add(u *User) error

	// List retrieves all users from the database.
	List() ([]User, error)
}

// WishlistRepository defines persistence operations for wishlists.
type WishlistRepository interface {
	// Add saves a new wishlist to the database.
	Add(w *Wishlist) error

	// List retrieves all wishlists for a given user.
	List(userID uint) ([]Wishlist, error)

	// Delete removes a wishlist by its ID for a given user.
	Delete(userID, wishlistID uint) error
}

// BookRepository defines persistence operations for books.
type BookRepository interface {
	// Add saves a new book to the database.
	Add(b *Book) error

	// List retrieves all books in a given wishlist.
	List(wishlistID uint) ([]Book, error)

	// Delete removes a book by its ID from a wishlist.
	Delete(wishlistID, bookID uint) error
}
