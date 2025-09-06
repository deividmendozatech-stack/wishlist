package domain

type Wishlist struct {
	ID     uint
	Name   string
	UserID uint
	Books  []Book
}
