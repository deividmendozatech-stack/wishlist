package domain

// Wishlist groups books under a user.
type Wishlist struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}
