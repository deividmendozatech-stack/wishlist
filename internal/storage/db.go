package storage

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// Importamos el driver SQLite puro Go (sin CGO)
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

// Init abre o crea la base de datos SQLite.
func Init() {
	// Lee la ruta desde la variable de entorno (útil para Docker)
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "wishlist.db"
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al abrir base de datos: %v", err)
	}

	DB = db
	log.Println("Base de datos inicializada en", path)
}
