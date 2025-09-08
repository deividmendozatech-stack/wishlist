// @title Wishlist API
// @version 1.0
// @description API para gestionar usuarios, wishlists y libros
// @host localhost:8080
// @BasePath /api

package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/deividmendozatech-stack/wishlist/docs" // swagger embed files
	"github.com/deividmendozatech-stack/wishlist/internal/domain"
	"github.com/deividmendozatech-stack/wishlist/internal/handler"
	"github.com/deividmendozatech-stack/wishlist/internal/platform/storage"
	gormrepo "github.com/deividmendozatech-stack/wishlist/internal/repository/gorm"
	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Conexión a SQLite
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "wishlist.db"
	}

	db, err := storage.NewConnection(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Migraciones
	if err := db.AutoMigrate(&domain.User{}, &domain.Wishlist{}, &domain.Book{}); err != nil {
		log.Fatal(err)
	}

	// Repositorios
	userRepo := gormrepo.NewUserRepo(db)
	wishlistRepo := gormrepo.NewWishlistRepo(db)
	bookRepo := gormrepo.NewBookRepo(db)

	// Servicios
	userSvc := service.NewUserService(userRepo)
	wishlistSvc := service.NewWishlistService(wishlistRepo)
	bookSvc := service.NewBookService(bookRepo)

	// Handlers
	mainHandler := handler.NewHTTPHandler(wishlistSvc, userSvc)
	bookHandler := handler.NewBookHTTP(bookSvc)

	// Router
	r := mux.NewRouter()

	// Rutas API
	api := r.PathPrefix("/api").Subrouter()
	mainHandler.RegisterRoutes(api)
	bookHandler.RegisterBookRoutes(api)

	// Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
