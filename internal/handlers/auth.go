package handlers

import (
	"encoding/json"
	"net/http"

	// bcrypt para encriptar contraseñas
	"golang.org/x/crypto/bcrypt"

	// paquetes internos de tu módulo
	"github.com/deividmendozatech_stack/wishlist/internal/models"
	"github.com/deividmendozatech_stack/wishlist/internal/storage"
)

// Estructura de la solicitud de registro
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterHandler crea un nuevo usuario con contraseña encriptada
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Generar hash de la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al procesar contraseña", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Username:     req.Username,
		PasswordHash: string(hash),
	}

	if result := storage.DB.Create(&user); result.Error != nil {
		http.Error(w, "Usuario ya existe", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
