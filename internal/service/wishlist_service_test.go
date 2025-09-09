package service_test

import (
	"errors"
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/stretchr/testify/assert"
)

// mockWishlistRepo is a lightweight mock implementation of service.WishlistRepository.
// It allows injecting custom behavior for Add, List, and Delete methods in tests.
type mockWishlistRepo struct {
	addFn    func(*service.Wishlist) error
	listFn   func(uint) ([]service.Wishlist, error)
	deleteFn func(uint, uint) error
}

func (m *mockWishlistRepo) Add(w *service.Wishlist) error {
	if m.addFn != nil {
		return m.addFn(w)
	}
	return nil
}

func (m *mockWishlistRepo) List(userID uint) ([]service.Wishlist, error) {
	if m.listFn != nil {
		return m.listFn(userID)
	}
	return []service.Wishlist{}, nil
}

func (m *mockWishlistRepo) Delete(userID, wishlistID uint) error {
	if m.deleteFn != nil {
		return m.deleteFn(userID, wishlistID)
	}
	return nil
}

// TestWishlistService_Create verifies that a wishlist can be created without errors.
func TestWishlistService_Create(t *testing.T) {
	mockRepo := &mockWishlistRepo{}
	svc := service.NewWishlistService(mockRepo)

	err := svc.Create(1, "Mi lista")
	assert.NoError(t, err)
}

// TestWishlistService_List verifies that wishlists are correctly retrieved from the repository.
func TestWishlistService_List(t *testing.T) {
	mockRepo := &mockWishlistRepo{
		listFn: func(userID uint) ([]service.Wishlist, error) {
			return []service.Wishlist{{ID: 1, Name: "Lista de prueba"}}, nil
		},
	}
	svc := service.NewWishlistService(mockRepo)

	lists, err := svc.List(1)
	assert.NoError(t, err)
	assert.Len(t, lists, 1)
	assert.Equal(t, "Lista de prueba", lists[0].Name)
}

// TestWishlistService_Delete verifies that a wishlist can be deleted successfully.
func TestWishlistService_Delete(t *testing.T) {
	mockRepo := &mockWishlistRepo{}
	svc := service.NewWishlistService(mockRepo)

	err := svc.Delete(1, 1)
	assert.NoError(t, err)
}

// TestWishlistService_DeleteError ensures errors from the repository are propagated properly.
func TestWishlistService_DeleteError(t *testing.T) {
	mockRepo := &mockWishlistRepo{
		deleteFn: func(userID, wishlistID uint) error {
			return errors.New("db error")
		},
	}
	svc := service.NewWishlistService(mockRepo)

	err := svc.Delete(1, 1)
	assert.Error(t, err)
}
