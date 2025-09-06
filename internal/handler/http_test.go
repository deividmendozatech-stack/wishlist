package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/handler"
)

type mockSvc struct{ mock.Mock }

func (m *mockSvc) Create(userID uint, name string) error {
	return m.Called(userID, name).Error(0)
}
func (m *mockSvc) List(userID uint) ([]domain.Wishlist, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Wishlist), args.Error(1)
}
func (m *mockSvc) Delete(userID, id uint) error {
	return m.Called(userID, id).Error(0)
}

func TestCreateWishlistHandler(t *testing.T) {
	ms := new(mockSvc)
	ms.On("Create", uint(1), "Mi lista").Return(nil)

	h := handler.NewHTTPHandler(ms)
	r := mux.NewRouter()
	h.RegisterRoutes(r)

	body := []byte(`{"name":"Mi lista"}`)
	req := httptest.NewRequest(http.MethodPost, "/wishlist", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	ms.AssertExpectations(t)
}

func TestListWishlistsHandler(t *testing.T) {
	ms := new(mockSvc)
	// Preparamos la respuesta simulada
	expected := []domain.Wishlist{
		{ID: 1, Name: "Lista A", UserID: 1},
		{ID: 2, Name: "Lista B", UserID: 1},
	}
	ms.On("List", uint(1)).Return(expected, nil)

	h := handler.NewHTTPHandler(ms)
	r := mux.NewRouter()
	h.RegisterRoutes(r)

	req := httptest.NewRequest(http.MethodGet, "/wishlist", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// decodificamos el JSON devuelto
	var got []domain.Wishlist
	err := json.NewDecoder(rr.Body).Decode(&got)
	assert.NoError(t, err)
	assert.Len(t, got, 2)
	assert.Equal(t, "Lista A", got[0].Name)

	ms.AssertExpectations(t)
}

func TestDeleteWishlistHandler(t *testing.T) {
	ms := new(mockSvc)
	ms.On("Delete", uint(1), uint(99)).Return(nil)

	h := handler.NewHTTPHandler(ms)
	r := mux.NewRouter()
	h.RegisterRoutes(r)

	req := httptest.NewRequest(http.MethodDelete, "/wishlist/99", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	ms.AssertExpectations(t)
}
