package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// paquetes internos
	"github.com/deividmendozatech-stack/wishlist/internal/auth"
	"github.com/deividmendozatech-stack/wishlist/internal/handlers"
	"github.com/deividmendozatech-stack/wishlist/internal/models"
	"github.com/deividmendozatech-stack/wishlist/internal/storage"
)

func main() {
	// carga variables de entorno
	_ = godotenv.Load()

	// Inicializa la base de datos y migra modelos
	storage.Init()
	storage.DB.AutoMigrate(&models.User{}, &models.Wishlist{}, &models.Book{})

	r := mux.NewRouter()

	// subrouter con prefijo /api
	api := r.PathPrefix("/api").Subrouter()

	// Rutas públicas
	api.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	api.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Rutas protegidas (usa middleware JWT)
	secured := api.NewRoute().Subrouter()
	secured.Use(auth.AuthMiddleware)
	secured.HandleFunc("/wishlist", handlers.CreateWishlistHandler).Methods("POST")
	secured.HandleFunc("/wishlist", handlers.ListWishlistsHandler).Methods("GET")
	secured.HandleFunc("/wishlist/{id}", handlers.DeleteWishlistHandler).Methods("DELETE")
	secured.HandleFunc("/wishlist/{id}/books", handlers.AddBookHandler).Methods("POST")
	secured.HandleFunc("/wishlist/{id}/books", handlers.ListBooksHandler).Methods("GET")
	secured.HandleFunc("/wishlist/{id}/books/{bookID}", handlers.RemoveBookHandler).Methods("DELETE")
	secured.HandleFunc("/books/search", handlers.SearchGoogleBooksHandler).Methods("GET")

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
