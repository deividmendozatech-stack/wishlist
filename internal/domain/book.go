package domain

import "time"

// Book is an item stored inside a wishlist.
type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description,omitempty"`
	WishlistID  uint      `json:"wishlist_id" gorm:"index"` // linked wishlist
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
