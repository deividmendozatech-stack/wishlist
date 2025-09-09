package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/deividmendozatech-stack/wishlist/docs"

	"github.com/deividmendozatech-stack/wishlist/internal/handler"
	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/deividmendozatech-stack/wishlist/internal/storage"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// main is the entry point of the Wishlist API.
// It connects to the database, runs migrations, initializes repositories,
// services, handlers, sets up routes, and starts the HTTP server.
func main() {
	// Load database path from environment variable (default: wishlist.db)
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "wishlist.db"
	}

	// Connect to SQLite database
	db, err := storage.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations for User, Wishlist, and Book models
	if err := db.AutoMigrate(&service.User{}, &service.Wishlist{}, &service.Book{}); err != nil {
		log.Fatal(err)
	}

	// Initialize repositories
	userRepo := storage.NewUserRepo(db)
	wishlistRepo := storage.NewWishlistRepo(db)
	bookRepo := storage.NewBookRepo(db)

	// Initialize services (business logic layer)
	userSvc := service.NewUserService(userRepo)
	wishlistSvc := service.NewWishlistService(wishlistRepo)
	bookSvc := service.NewBookService(bookRepo)
	googleSvc := service.NewGoogleBooksService()

	// Initialize HTTP handlers
	mainHandler := handler.NewHTTPHandler(wishlistSvc, userSvc)
	bookHandler := handler.NewBookHTTP(bookSvc)
	googleHandler := handler.NewGoogleBooksHTTP(googleSvc)

	// Create a new router
	r := mux.NewRouter()

	// API routes (grouped under /api)
	api := r.PathPrefix("/api").Subrouter()

	// User and Wishlist routes
	api.HandleFunc("/users/register", mainHandler.RegisterUser).Methods(http.MethodPost)    // Register a new user
	api.HandleFunc("/users", mainHandler.ListUsers).Methods(http.MethodGet)                 // List all users
	api.HandleFunc("/wishlist", mainHandler.CreateWishlist).Methods(http.MethodPost)        // Create a new wishlist
	api.HandleFunc("/wishlist", mainHandler.ListWishlists).Methods(http.MethodGet)          // List all wishlists
	api.HandleFunc("/wishlist/{id}", mainHandler.DeleteWishlist).Methods(http.MethodDelete) // Delete a wishlist by ID

	// Book routes (within a wishlist)
	api.HandleFunc("/wishlist/{id}/books", bookHandler.AddBook).Methods(http.MethodPost)               // Add a book to a wishlist
	api.HandleFunc("/wishlist/{id}/books", bookHandler.ListBooks).Methods(http.MethodGet)              // List books in a wishlist
	api.HandleFunc("/wishlist/{id}/books/{bookID}", bookHandler.DeleteBook).Methods(http.MethodDelete) // Delete a book from a wishlist

	// Google Books routes (search integration)
	googleHandler.RegisterGoogleRoutes(api)

	// Swagger UI (API documentation)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start HTTP server
	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
