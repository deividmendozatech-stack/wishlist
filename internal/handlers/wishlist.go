package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/deividmendozatech_stack/wishlist/internal/auth"
	"github.com/deividmendozatech_stack/wishlist/internal/models"
	"github.com/deividmendozatech_stack/wishlist/internal/storage"

	"github.com/gorilla/mux"
)

type CreateWishlistRequest struct {
	Name string `json:"name"`
}

func CreateWishlistHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateWishlistRequest
	json.NewDecoder(r.Body).Decode(&req)

	userID := r.Context().Value(auth.UserIDKey).(uint)
	wl := models.Wishlist{Name: req.Name, UserID: userID}
	storage.DB.Create(&wl)
	w.WriteHeader(http.StatusCreated)
}

func ListWishlistsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)
	var lists []models.Wishlist
	storage.DB.Where("user_id = ?", userID).Find(&lists)
	json.NewEncoder(w).Encode(lists)
}

func DeleteWishlistHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)
	id := mux.Vars(r)["id"]
	storage.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&models.Wishlist{})
	w.WriteHeader(http.StatusNoContent)
}

type AddBookRequest struct {
	GoogleID  string `json:"google_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)
	wishlistID := mux.Vars(r)["id"]

	var wl models.Wishlist
	if storage.DB.First(&wl, "id = ? AND user_id = ?", wishlistID, userID).Error != nil {
		http.Error(w, "Lista no encontrada", http.StatusNotFound)
		return
	}

	var req AddBookRequest
	json.NewDecoder(r.Body).Decode(&req)

	book := models.Book{
		GoogleID:   req.GoogleID,
		Title:      req.Title,
		Author:     req.Author,
		Publisher:  req.Publisher,
		WishlistID: wl.ID,
	}
	storage.DB.Create(&book)
	w.WriteHeader(http.StatusCreated)
}

func ListBooksHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)
	wishlistID := mux.Vars(r)["id"]

	var wl models.Wishlist
	if storage.DB.First(&wl, "id = ? AND user_id = ?", wishlistID, userID).Error != nil {
		http.Error(w, "Lista no encontrada", http.StatusNotFound)
		return
	}
	var books []models.Book
	storage.DB.Where("wishlist_id = ?", wl.ID).Find(&books)
	json.NewEncoder(w).Encode(books)
}

func RemoveBookHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(uint)
	wishlistID := mux.Vars(r)["id"]
	bookID := mux.Vars(r)["bookID"]

	var wl models.Wishlist
	if storage.DB.First(&wl, "id = ? AND user_id = ?", wishlistID, userID).Error != nil {
		http.Error(w, "Lista no encontrada", http.StatusNotFound)
		return
	}
	storage.DB.Delete(&models.Book{}, "id = ? AND wishlist_id = ?", bookID, wl.ID)
	w.WriteHeader(http.StatusNoContent)
}

func SearchGoogleBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	key := r.URL.Query().Get("key")
	if query == "" || key == "" {
		http.Error(w, "Falta q o key", http.StatusBadRequest)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/books/v1/volumes?q=" + query + "&key=" + key)
	if err != nil {
		http.Error(w, "Error consultando Google Books", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	// Copiar la respuesta cruda de Google Books al cliente
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Error enviando respuesta", http.StatusInternalServerError)
		return
	}
}
