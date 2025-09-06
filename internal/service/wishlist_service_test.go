package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/service"
)

// mockRepo implementa repository.WishlistRepository
type mockRepo struct{ mock.Mock }

func (m *mockRepo) Create(w *domain.Wishlist) error {
	return m.Called(w).Error(0)
}
func (m *mockRepo) FindByUser(userID uint) ([]domain.Wishlist, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Wishlist), args.Error(1)
}
func (m *mockRepo) Delete(userID, id uint) error {
	return m.Called(userID, id).Error(0)
}

func TestCreateWishlist(t *testing.T) {
	mr := new(mockRepo)
	svc := service.NewWishlistService(mr)

	mr.On("Create", mock.AnythingOfType("*domain.Wishlist")).Return(nil)

	err := svc.Create(1, "Mi lista")
	assert.NoError(t, err)
	mr.AssertExpectations(t)
}

func TestListWishlist(t *testing.T) {
	mr := new(mockRepo)
	svc := service.NewWishlistService(mr)

	expected := []domain.Wishlist{
		{ID: 1, Name: "Lista A", UserID: 1},
		{ID: 2, Name: "Lista B", UserID: 1},
	}
	mr.On("FindByUser", uint(1)).Return(expected, nil)

	got, err := svc.List(1)
	assert.NoError(t, err)
	assert.Len(t, got, 2)
	assert.Equal(t, "Lista A", got[0].Name)

	mr.AssertExpectations(t)
}

func TestDeleteWishlist(t *testing.T) {
	mr := new(mockRepo)
	svc := service.NewWishlistService(mr)

	mr.On("Delete", uint(1), uint(99)).Return(nil)

	err := svc.Delete(1, 99)
	assert.NoError(t, err)

	mr.AssertExpectations(t)
}
