package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	svc "github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/gorilla/mux"
)

//
// ──────────────── MOCKS ────────────────
//

// mockWishlist implementa service.WishlistUsecase
type mockWishlist struct{}

var _ svc.WishlistUsecase = (*mockWishlist)(nil)

func (m *mockWishlist) Create(userID uint, name string) error {
	return nil
}
func (m *mockWishlist) List(userID uint) ([]domain.Wishlist, error) {
	return []domain.Wishlist{
		{ID: 1, UserID: userID, Name: "TestList"},
	}, nil
}
func (m *mockWishlist) Delete(userID, id uint) error {
	return nil
}

// mockUser implementa service.UserUsecase
type mockUser struct{}

var _ svc.UserUsecase = (*mockUser)(nil)

func (m *mockUser) Register(username, password string) error {
	return nil
}
func (m *mockUser) List() ([]domain.User, error) {
	return []domain.User{
		{ID: 1, Username: "david"},
	}, nil
}

// mockBook implementa service.BookUsecase
type mockBook struct{}

var _ svc.BookUsecase = (*mockBook)(nil)

func (m *mockBook) Add(wishlistID uint, title, author string) error {
	return nil
}
func (m *mockBook) List(wishlistID uint) ([]domain.Book, error) {
	return []domain.Book{
		{ID: 1, WishlistID: wishlistID, Title: "BookTest", Author: "Anon"},
	}, nil
}
func (m *mockBook) Delete(wishlistID, bookID uint) error {
	return nil
}

//
// ──────────────── HELPERS ────────────────
//

func setupRouter() *mux.Router {
	wSvc := &mockWishlist{}
	uSvc := &mockUser{}
	bSvc := &mockBook{}

	mainHandler := NewHTTPHandler(wSvc, uSvc)
	bookHandler := NewBookHTTP(bSvc)

	r := mux.NewRouter()
	mainHandler.RegisterRoutes(r)
	bookHandler.RegisterBookRoutes(r)

	return r
}

//
// ──────────────── TESTS ────────────────
//

// TestEndpoints prueba endpoints básicos de usuarios, wishlist y libros
func TestEndpoints(t *testing.T) {
	router := setupRouter()

	// ---------- USER REGISTER ----------
	bodyUser, _ := json.Marshal(RegisterUserRequest{Username: "david", Password: "1234"})
	req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(bodyUser))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}

	// ---------- CREATE WISHLIST ----------
	bodyWL, _ := json.Marshal(CreateWishlistRequest{Name: "MyList"})
	req = httptest.NewRequest(http.MethodPost, "/wishlist", bytes.NewBuffer(bodyWL))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}

	// ---------- LIST WISHLIST ----------
	req = httptest.NewRequest(http.MethodGet, "/wishlist", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}

	// ---------- ADD BOOK ----------
	bookPayload := `{"title":"Go 101","author":"Unknown"}`
	req = httptest.NewRequest(http.MethodPost, "/wishlist/1/books", bytes.NewBufferString(bookPayload))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", resp.Code)
	}

	// ---------- LIST BOOKS ----------
	req = httptest.NewRequest(http.MethodGet, "/wishlist/1/books", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}

	// ---------- LIST USERS ----------
	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.Code)
	}

}
