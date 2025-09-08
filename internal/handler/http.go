package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/gorilla/mux"
)

//
// ───────────────────────── MODELOS PARA SWAGGER ─────────────────────────
//

// CreateWishlistRequest representa el body de creación de wishlist
type CreateWishlistRequest struct {
	Name string `json:"name" example:"Mi lista de libros"`
}

// RegisterUserRequest representa el body de registro de usuario
type RegisterUserRequest struct {
	Username string `json:"username" example:"david"`
	Password string `json:"password" example:"1234"`
}

// AddBookRequest representa el body para agregar un libro
type AddBookRequest struct {
	Title  string `json:"title"  example:"El Principito"`
	Author string `json:"author" example:"Antoine de Saint-Exupéry"`
}

//
// ───────────────────────── HANDLERS ─────────────────────────
//

// HTTPHandler agrupa endpoints de usuarios y wishlists
type HTTPHandler struct {
	wishlist service.WishlistUsecase
	users    service.UserUsecase
}

// BookHTTP agrupa endpoints de libros
type BookHTTP struct {
	book service.BookUsecase
}

//
// ───────────────────────── CONSTRUCTORES ─────────────────────────
//

func NewHTTPHandler(w service.WishlistUsecase, u service.UserUsecase) *HTTPHandler {
	return &HTTPHandler{wishlist: w, users: u}
}

func NewBookHTTP(b service.BookUsecase) *BookHTTP {
	return &BookHTTP{book: b}
}

//
// ───────────────────────── RUTAS ─────────────────────────
//

func (h *HTTPHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", h.RegisterUser).Methods(http.MethodPost)

	r.HandleFunc("/wishlist", h.CreateWishlist).Methods(http.MethodPost)
	r.HandleFunc("/wishlist", h.ListWishlists).Methods(http.MethodGet)
	r.HandleFunc("/wishlist/{id}", h.DeleteWishlist).Methods(http.MethodDelete)
	r.HandleFunc("/users", h.ListUsers).Methods(http.MethodGet)

}

func (h *BookHTTP) RegisterBookRoutes(r *mux.Router) {
	r.HandleFunc("/wishlist/{id}/books", h.AddBook).Methods(http.MethodPost)
	r.HandleFunc("/wishlist/{id}/books", h.ListBooks).Methods(http.MethodGet)
	r.HandleFunc("/wishlist/{id}/books/{bookID}", h.DeleteBook).Methods(http.MethodDelete)
}

//
// ───────────────────────── USERS ─────────────────────────
//

// RegisterUser godoc
// @Summary Registrar nuevo usuario
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterUserRequest true "Usuario"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /users/register [post]
func (h *HTTPHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	if err := h.users.Register(req.Username, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//
// ───────────────────────── WISHLIST ─────────────────────────
//

// CreateWishlist godoc
// @Summary Crea una nueva wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Param data body CreateWishlistRequest true "Datos de la lista"
// @Success 201
// @Failure 400
// @Router /wishlist [post]
func (h *HTTPHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {
	var req CreateWishlistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	userID := uint(1) // en real vendría del token JWT
	if err := h.wishlist.Create(userID, req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListWishlists godoc
// @Summary Lista las wishlists del usuario
// @Tags wishlist
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /wishlist [get]
func (h *HTTPHandler) ListWishlists(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)
	lists, err := h.wishlist.List(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

// DeleteWishlist godoc
// @Summary Elimina una wishlist por ID
// @Tags wishlist
// @Param id path int true "ID de la wishlist"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /wishlist/{id} [delete]
func (h *HTTPHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	userID := uint(1)
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.wishlist.Delete(userID, uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//
// ───────────────────────── BOOKS ─────────────────────────
//

// AddBook godoc
// @Summary Agrega un libro a la wishlist
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID de la wishlist"
// @Param data body AddBookRequest true "Datos del libro"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /wishlist/{id}/books [post]
func (h *BookHTTP) AddBook(w http.ResponseWriter, r *http.Request) {
	wishlistIDStr := mux.Vars(r)["id"]
	wishlistID, err := strconv.Atoi(wishlistIDStr)
	if err != nil {
		http.Error(w, "invalid wishlist id", http.StatusBadRequest)
		return
	}

	var req AddBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	if err := h.book.Add(uint(wishlistID), req.Title, req.Author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListBooks godoc
// @Summary Lista los libros de una wishlist
// @Tags books
// @Produce json
// @Param id path int true "ID de la wishlist"
// @Success 200 {array} map[string]interface{}
// @Router /wishlist/{id}/books [get]
func (h *BookHTTP) ListBooks(w http.ResponseWriter, r *http.Request) {
	wishlistIDStr := mux.Vars(r)["id"]
	wishlistID, err := strconv.Atoi(wishlistIDStr)
	if err != nil {
		http.Error(w, "invalid wishlist id", http.StatusBadRequest)
		return
	}

	books, err := h.book.List(uint(wishlistID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// DeleteBook godoc
// @Summary Elimina un libro de una wishlist
// @Tags books
// @Param id path int true "ID de la wishlist"
// @Param bookID path int true "ID del libro"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /wishlist/{id}/books/{bookID} [delete]
func (h *BookHTTP) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wishlistID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid wishlist id", http.StatusBadRequest)
		return
	}
	bookID, err := strconv.Atoi(vars["bookID"])
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	if err := h.book.Delete(uint(wishlistID), uint(bookID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ListUsers godoc
// @Summary Lista los usuarios registrados
// @Tags users
// @Produce json
// @Success 200 {array} domain.User
// @Router /users [get]
func (h *HTTPHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.users.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GoogleBooksHTTP maneja las búsquedas externas en Google Books
type GoogleBooksHTTP struct {
	api service.GoogleBooksUsecase
}

func NewGoogleBooksHTTP(api service.GoogleBooksUsecase) *GoogleBooksHTTP {
	return &GoogleBooksHTTP{api: api}
}

// RegisterGoogleRoutes registra la ruta GET /books/search
func (h *GoogleBooksHTTP) RegisterGoogleRoutes(r *mux.Router) {
	r.HandleFunc("/books/search", h.SearchBooks).Methods(http.MethodGet)
}

// SearchBooks godoc
// @Summary Busca libros en Google Books
// @Tags books
// @Produce json
// @Param q query string true "Término de búsqueda"
// @Success 200 {array} service.GoogleBook
// @Failure 400
// @Failure 500
// @Router /books/search [get]
func (h *GoogleBooksHTTP) SearchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing query param q", http.StatusBadRequest)
		return
	}

	results, err := h.api.Search(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
