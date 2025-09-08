package domain

import "time"

// Book representa un libro dentro de una wishlist
type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description,omitempty"`
	WishlistID  uint      `json:"wishlist_id" gorm:"index"` // relación con Wishlist
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
