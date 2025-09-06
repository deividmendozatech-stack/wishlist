// @title Wishlist API
// @version 1.0
// @description API para gestionar listas de deseos de libros
// @host localhost:8080
// @BasePath /api

package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/deividmendozatech-stack/wishlist/docs" // swagger embed files
	"github.com/deividmendozatech-stack/wishlist/internal/handler"
	"github.com/deividmendozatech-stack/wishlist/internal/platform/storage"
	gormrepo "github.com/deividmendozatech-stack/wishlist/internal/repository/gorm"
	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db, err := storage.NewConnection(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	repo := gormrepo.NewWishlistRepo(db)
	svc := service.NewWishlistService(repo)
	h := handler.NewHTTPHandler(svc)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	h.RegisterRoutes(api)
	//Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Servidor en http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
