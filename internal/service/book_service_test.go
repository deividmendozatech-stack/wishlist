package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
// ─────────────────────────── MOCK REPOSITORY ───────────────────────────
//

// mockBookRepo simulates the Book repository layer for testing purposes.
// It tracks whether methods were called and can inject errors to validate behavior.
type mockBookRepo struct {
	addCalled    bool
	listCalled   bool
	deleteCalled bool
	books        []Book
	err          error
}

// Add simulates inserting a book into the repository.
func (m *mockBookRepo) Add(book *Book) error {
	m.addCalled = true
	if m.err != nil {
		return m.err
	}
	book.ID = 1
	m.books = append(m.books, *book)
	return nil
}

// List simulates fetching all books by wishlist ID.
func (m *mockBookRepo) List(wishlistID uint) ([]Book, error) {
	m.listCalled = true
	if m.err != nil {
		return nil, m.err
	}
	return m.books, nil
}

// Delete simulates removing a book by wishlist ID and book ID.
func (m *mockBookRepo) Delete(wishlistID, bookID uint) error {
	m.deleteCalled = true
	if m.err != nil {
		return m.err
	}
	for i, b := range m.books {
		if b.ID == bookID && b.WishlistID == wishlistID {
			m.books = append(m.books[:i], m.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

//
// ─────────────────────────── UNIT TESTS ───────────────────────────
//

// TestBookService_Add ensures that a book can be added successfully
// and verifies repository interaction.
func TestBookService_Add(t *testing.T) {
	mockRepo := &mockBookRepo{}
	svc := NewBookService(mockRepo)

	err := svc.Add(1, "Go Programming", "Alice")
	assert.NoError(t, err)
	assert.True(t, mockRepo.addCalled)
	assert.Len(t, mockRepo.books, 1)
	assert.Equal(t, "Go Programming", mockRepo.books[0].Title)
}

// TestBookService_List ensures that books are correctly listed
// from the repository by wishlist ID.
func TestBookService_List(t *testing.T) {
	mockRepo := &mockBookRepo{
		books: []Book{{ID: 1, WishlistID: 1, Title: "Go 101", Author: "Bob"}},
	}
	svc := NewBookService(mockRepo)

	books, err := svc.List(1)
	assert.NoError(t, err)
	assert.True(t, mockRepo.listCalled)
	assert.Len(t, books, 1)
	assert.Equal(t, "Go 101", books[0].Title)
}

// TestBookService_Delete ensures that a book can be deleted
// and validates repository interaction.
func TestBookService_Delete(t *testing.T) {
	mockRepo := &mockBookRepo{
		books: []Book{{ID: 1, WishlistID: 1, Title: "Go 101", Author: "Bob"}},
	}
	svc := NewBookService(mockRepo)

	err := svc.Delete(1, 1)
	assert.NoError(t, err)
	assert.True(t, mockRepo.deleteCalled)
	assert.Len(t, mockRepo.books, 0)
}
