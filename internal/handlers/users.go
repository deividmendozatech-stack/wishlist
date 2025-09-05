package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deividmendozatech-stack/wishlist/internal/models"
	"github.com/deividmendozatech-stack/wishlist/internal/storage"
)

// ListUsersHandler devuelve todos los usuarios (sólo para pruebas)
func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := storage.DB.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
