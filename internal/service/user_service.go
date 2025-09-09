package service

import "errors"

// Predefined business-level errors.
var (
	// ErrInvalidInput is returned when username or password are empty.
	ErrInvalidInput = errors.New("username and password cannot be empty")
)

// userService is the concrete implementation of the UserUsecase interface.
// It contains the business logic for user-related operations.
type userService struct {
	repo UserRepository
}

// NewUserService creates and returns a new UserUsecase implementation.
func NewUserService(repo UserRepository) UserUsecase {
	return &userService{repo: repo}
}

// Register validates input and registers a new user by delegating to the repository.
// Returns ErrInvalidInput if username or password are empty.
func (s *userService) Register(username, password string) error {
	if username == "" || password == "" {
		return ErrInvalidInput
	}
	user := &User{Username: username, Password: password}
	return s.repo.Add(user)
}

// List retrieves all registered users by delegating to the repository.
func (s *userService) List() ([]User, error) {
	return s.repo.List()
}
