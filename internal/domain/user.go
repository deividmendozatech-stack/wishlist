package domain

// User represents an application user.
type User struct {
	ID       uint
	Username string
	Password string // hashed password
}
