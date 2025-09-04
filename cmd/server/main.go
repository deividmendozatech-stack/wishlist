package main

import (
	"log"
	"net/http"

	"wishlist/internal/auth"
	"wishlist/internal/handlers"
	"wishlist/internal/models"
	"wishlist/internal/storage"

	"github.com/gorilla/mux"
)

func main() {
	// DB y migraciones
	storage.Init()
	storage.DB.AutoMigrate(&models.User{}, &models.Wishlist{}, &models.Book{})

	r := mux.NewRouter()

	// Rutas públicas
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Rutas protegidas
	api := r.PathPrefix("/api").Subrouter()
	api.Use(auth.AuthMiddleware)
	api.HandleFunc("/wishlist", handlers.CreateWishlistHandler).Methods("POST")
	api.HandleFunc("/wishlist", handlers.ListWishlistsHandler).Methods("GET")
	api.HandleFunc("/wishlist/{id}", handlers.DeleteWishlistHandler).Methods("DELETE")
	api.HandleFunc("/wishlist/{id}/books", handlers.AddBookHandler).Methods("POST")
	api.HandleFunc("/wishlist/{id}/books", handlers.ListBooksHandler).Methods("GET")
	api.HandleFunc("/wishlist/{id}/books/{bookID}", handlers.RemoveBookHandler).Methods("DELETE")
	api.HandleFunc("/books/search", handlers.SearchGoogleBooksHandler).Methods("GET")

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
