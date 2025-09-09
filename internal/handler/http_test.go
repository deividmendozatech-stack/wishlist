package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/gorilla/mux"
)

//
// ──────────────── MOCKS ────────────────
//

// mockWishlist is a mock implementation of WishlistUsecase for testing purposes.
type mockWishlist struct{}

var _ service.WishlistUsecase = (*mockWishlist)(nil)

func (m *mockWishlist) Create(userID uint, name string) error { return nil }
func (m *mockWishlist) List(userID uint) ([]service.Wishlist, error) {
	return []service.Wishlist{{ID: 1, UserID: userID, Name: "TestList"}}, nil
}
func (m *mockWishlist) Delete(userID, id uint) error { return nil }

// mockUser is a mock implementation of UserUsecase for testing purposes.
type mockUser struct{}

var _ service.UserUsecase = (*mockUser)(nil)

func (m *mockUser) Register(username, password string) error { return nil }
func (m *mockUser) List() ([]service.User, error) {
	return []service.User{{ID: 1, Username: "david"}}, nil
}

// mockBook is a mock implementation of BookUsecase for testing purposes.
type mockBook struct{}

var _ service.BookUsecase = (*mockBook)(nil)

func (m *mockBook) Add(wishlistID uint, title, author string) error { return nil }
func (m *mockBook) List(wishlistID uint) ([]service.Book, error) {
	return []service.Book{{ID: 1, WishlistID: wishlistID, Title: "BookTest", Author: "Anon"}}, nil
}
func (m *mockBook) Delete(wishlistID, bookID uint) error { return nil }

//
// ──────────────── HELPERS ────────────────
//

// setupRouter builds a test HTTP router with mock services.
// It registers the same routes as in main.go but uses mock implementations.
func setupRouter() *mux.Router {
	wSvc := &mockWishlist{}
	uSvc := &mockUser{}
	bSvc := &mockBook{}

	mainHandler := NewHTTPHandler(wSvc, uSvc)
	bookHandler := NewBookHTTP(bSvc)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	// Same routes defined in main.go
	api.HandleFunc("/users/register", mainHandler.RegisterUser).Methods(http.MethodPost)
	api.HandleFunc("/users", mainHandler.ListUsers).Methods(http.MethodGet)
	api.HandleFunc("/wishlist", mainHandler.CreateWishlist).Methods(http.MethodPost)
	api.HandleFunc("/wishlist", mainHandler.ListWishlists).Methods(http.MethodGet)
	api.HandleFunc("/wishlist/{id}", mainHandler.DeleteWishlist).Methods(http.MethodDelete)

	api.HandleFunc("/wishlist/{id}/books", bookHandler.AddBook).Methods(http.MethodPost)
	api.HandleFunc("/wishlist/{id}/books", bookHandler.ListBooks).Methods(http.MethodGet)
	api.HandleFunc("/wishlist/{id}/books/{bookID}", bookHandler.DeleteBook).Methods(http.MethodDelete)

	return r
}

//
// ──────────────── TESTS ────────────────
//

// TestEndpoints validates core API routes for users, wishlists, and books
// using mocked services instead of a real database.
func TestEndpoints(t *testing.T) {
	router := setupRouter()

	// USER REGISTER
	bodyUser, _ := json.Marshal(RegisterUserRequest{Username: "david", Password: "1234"})
	req := httptest.NewRequest(http.MethodPost, "/api/users/register", bytes.NewBuffer(bodyUser))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}

	// CREATE WISHLIST
	bodyWL, _ := json.Marshal(CreateWishlistRequest{Name: "MyList"})
	req = httptest.NewRequest(http.MethodPost, "/api/wishlist", bytes.NewBuffer(bodyWL))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}

	// LIST WISHLIST
	req = httptest.NewRequest(http.MethodGet, "/api/wishlist", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}

	// ADD BOOK
	bookPayload := `{"title":"Go 101","author":"Unknown"}`
	req = httptest.NewRequest(http.MethodPost, "/api/wishlist/1/books", bytes.NewBufferString(bookPayload))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}

	// LIST BOOKS
	req = httptest.NewRequest(http.MethodGet, "/api/wishlist/1/books", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}

	// LIST USERS
	req = httptest.NewRequest(http.MethodGet, "/api/users", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}
}
