package service_test

import (
	"errors"
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//
// ──────────────────────────────── MOCK ────────────────────────────────
//

// mockUserRepo is a mock implementation of the UserRepository interface.
// It uses testify/mock to simulate database operations.
type mockUserRepo struct{ mock.Mock }

func (m *mockUserRepo) Add(u *service.User) error {
	args := m.Called(u)
	return args.Error(0)
}

func (m *mockUserRepo) List() ([]service.User, error) {
	args := m.Called()
	if val, ok := args.Get(0).([]service.User); ok {
		return val, args.Error(1)
	}
	return nil, args.Error(1)
}

//
// ──────────────────────────────── TESTS ────────────────────────────────
//

// TestRegister_Success ensures that a new user can be registered successfully.
func TestRegister_Success(t *testing.T) {
	repo := new(mockUserRepo)
	repo.On("Add", mock.AnythingOfType("*service.User")).Return(nil)

	svc := service.NewUserService(repo)
	err := svc.Register("david", "12345")

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

// TestRegister_EmptyFields checks that registering with empty username/password returns an error.
func TestRegister_EmptyFields(t *testing.T) {
	repo := new(mockUserRepo)
	svc := service.NewUserService(repo)

	err := svc.Register("", "")
	assert.Error(t, err)
}

// TestRegister_AddError simulates a database error when adding a user.
func TestRegister_AddError(t *testing.T) {
	repo := new(mockUserRepo)
	repo.On("Add", mock.AnythingOfType("*service.User")).Return(errors.New("db error"))

	svc := service.NewUserService(repo)
	err := svc.Register("john", "pwd")

	assert.Error(t, err)
	repo.AssertExpectations(t)
}

// TestListUsers_Success validates that the service returns a list of users when the repository succeeds.
func TestListUsers_Success(t *testing.T) {
	repo := new(mockUserRepo)
	repo.On("List").Return([]service.User{
		{ID: 1, Username: "alice"},
		{ID: 2, Username: "bob"},
	}, nil)

	svc := service.NewUserService(repo)
	users, err := svc.List()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "alice", users[0].Username)
	repo.AssertExpectations(t)
}

// TestListUsers_Error ensures that repository errors are properly propagated.
func TestListUsers_Error(t *testing.T) {
	repo := new(mockUserRepo)
	repo.On("List").Return(nil, errors.New("db error"))

	svc := service.NewUserService(repo)
	users, err := svc.List()

	assert.Error(t, err)
	assert.Nil(t, users)
	repo.AssertExpectations(t)
}
