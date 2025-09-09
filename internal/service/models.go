package service

//
// ─────────────────────────── DOMAIN MODELS ───────────────────────────
//

// User represents a registered user in the system.
// Stored in the database using GORM with a unique username.
type User struct {
	ID       uint   `gorm:"primaryKey"` // Auto-increment primary key
	Username string `gorm:"unique"`     // Unique username
	Password string // Hashed password
}

// Wishlist represents a list of desired books created by a user.
type Wishlist struct {
	ID     uint   `gorm:"primaryKey"` // Auto-increment primary key
	UserID uint   // Reference to the owning user
	Name   string // Name of the wishlist
}

// Book represents a book stored inside a wishlist.
type Book struct {
	ID         uint   `gorm:"primaryKey"` // Auto-increment primary key
	WishlistID uint   // Reference to the parent wishlist
	Title      string // Book title
	Author     string // Book author
}
