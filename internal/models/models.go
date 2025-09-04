package models

import "gorm.io/gorm"

// User representa un usuario registrado en el sistema
type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	PasswordHash string
}

// Wishlist representa una lista de deseos de un usuario
type Wishlist struct {
	gorm.Model
	Name   string
	UserID uint
	Books  []Book
}

// Book representa un libro dentro de una lista de deseos
type Book struct {
	gorm.Model
	GoogleID   string
	Title      string
	Author     string
	Publisher  string
	WishlistID uint
}
