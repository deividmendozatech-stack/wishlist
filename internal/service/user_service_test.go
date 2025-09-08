package service_test

import (
	"errors"
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockUserRepo is a lightweight mock of UserRepository.
type mockUserRepo struct{ mock.Mock }

func (m *mockUserRepo) Create(u *domain.User) error {
	args := m.Called(u)
	return args.Error(0)
}
func (m *mockUserRepo) FindByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if val, ok := args.Get(0).(*domain.User); ok {
		return val, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockUserRepo) FindAll() ([]domain.User, error) {
	args := m.Called()
	if val, ok := args.Get(0).([]domain.User); ok {
		return val, args.Error(1)
	}
	return nil, args.Error(1)
}

// TestRegister_Success verifies Register persists a valid user.
func TestRegister_Success(t *testing.T) {
	repo := new(mockUserRepo)
	repo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)

	svc := service.NewUserService(repo)
	err := svc.Register("david", "12345")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

// TestRegister_EmptyFields expects an error when username/password are empty.
func TestRegister_EmptyFields(t *testing.T) {
	repo := new(mockUserRepo)
	svc := service.NewUserService(repo)

	err := svc.Register("", "")
	assert.Error(t, err)
}

// TestRegister_CreateError ensures repo errors propagate.
func TestRegister_CreateError(t *testing.T) {
	repo := new(mockUserRepo)
	repo.On("Create", mock.AnythingOfType("*domain.User")).Return(errors.New("db error"))

	svc := service.NewUserService(repo)
	err := svc.Register("john", "pwd")

	assert.Error(t, err)
	repo.AssertExpectations(t)
}
