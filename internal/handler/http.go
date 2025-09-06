package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/gorilla/mux"
)

// HTTPHandler agrupa los endpoints HTTP.
type HTTPHandler struct {
	wishlist service.WishlistUsecase
}

// CreateWishlistRequest representa el body de creación de wishlist
type CreateWishlistRequest struct {
	Name string `json:"name"`
}

// NewHTTPHandler crea una nueva instancia del handler.
func NewHTTPHandler(w service.WishlistUsecase) *HTTPHandler {
	return &HTTPHandler{wishlist: w}
}

// RegisterRoutes registra todas las rutas en el router dado.
func (h *HTTPHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/wishlist", h.CreateWishlist).Methods(http.MethodPost)
	r.HandleFunc("/wishlist", h.ListWishlists).Methods(http.MethodGet)
	r.HandleFunc("/wishlist/{id}", h.DeleteWishlist).Methods(http.MethodDelete)
}

// CreateWishlist godoc
// @Summary Crea una nueva wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Param data body CreateWishlistRequest true "Datos de la lista"
// @Success 201
// @Failure 400
// @Router /wishlist [post]
func (h *HTTPHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {
	var req CreateWishlistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	userID := uint(1)
	if err := h.wishlist.Create(userID, req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListWishlists godoc
// @Summary Lista las wishlists del usuario
// @Tags wishlist
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /wishlist [get]
func (h *HTTPHandler) ListWishlists(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)
	lists, err := h.wishlist.List(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

// DeleteWishlist godoc
// @Summary Elimina una wishlist por ID
// @Tags wishlist
// @Param id path int true "ID de la wishlist"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /wishlist/{id} [delete]
func (h *HTTPHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.wishlist.Delete(userID, uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
