package handlers

import (
	"encoding/json"
	"net/http"

	// bcrypt para verificar contraseña
	"golang.org/x/crypto/bcrypt"

	// JWT propio
	"github.com/deividmendozatech-stack/wishlist/internal/auth"

	// Modelos y DB
	"github.com/deividmendozatech-stack/wishlist/internal/models"
	"github.com/deividmendozatech-stack/wishlist/internal/storage"
)

// LoginRequest estructura de datos de entrada
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler autentica al usuario y genera un token JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Buscar usuario
	var user models.User
	if result := storage.DB.First(&user, "username = ?", req.Username); result.Error != nil {
		http.Error(w, "Usuario no encontrado", http.StatusUnauthorized)
		return
	}

	// Verificar contraseña
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		http.Error(w, "Contraseña incorrecta", http.StatusUnauthorized)
		return
	}

	// Generar token
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, "Error generando token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}
