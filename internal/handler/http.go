package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/deividmendozatech-stack/wishlist/internal/service"
	"github.com/gorilla/mux"
)

//
// ───────────────────────── MODELS FOR SWAGGER ─────────────────────────
//

// CreateWishlistRequest represents the payload to create a wishlist.
// Used in Swagger documentation.
type CreateWishlistRequest struct {
	Name string `json:"name" example:"My book list"`
}

// RegisterUserRequest represents the payload to register a new user.
// Used in Swagger documentation.
type RegisterUserRequest struct {
	Username string `json:"username" example:"david"`
	Password string `json:"password" example:"1234"`
}

// AddBookRequest represents the payload to add a book into a wishlist.
// Used in Swagger documentation.
type AddBookRequest struct {
	Title  string `json:"title"  example:"The Little Prince"`
	Author string `json:"author" example:"Antoine de Saint-Exupéry"`
}

//
// ───────────────────────── HANDLERS ─────────────────────────
//

// HTTPHandler groups endpoints related to users and wishlists.
type HTTPHandler struct {
	wishlist service.WishlistUsecase
	users    service.UserUsecase
}

// BookHTTP groups endpoints related to books inside wishlists.
type BookHTTP struct {
	book service.BookUsecase
}

//
// ───────────────────────── CONSTRUCTORS ─────────────────────────
//

// NewHTTPHandler builds a handler for wishlist and user endpoints.
func NewHTTPHandler(w service.WishlistUsecase, u service.UserUsecase) *HTTPHandler {
	return &HTTPHandler{wishlist: w, users: u}
}

// NewBookHTTP builds a handler for book endpoints.
func NewBookHTTP(b service.BookUsecase) *BookHTTP {
	return &BookHTTP{book: b}
}

//
// ───────────────────────── USERS ─────────────────────────
//

// RegisterUser handles POST /users/register
// @Summary Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body RegisterUserRequest true "User data"
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

// ListUsers handles GET /users
// @Summary List registered users
// @Tags users
// @Produce json
// @Success 200 {array} service.User
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

//
// ───────────────────────── WISHLIST ─────────────────────────
//

// CreateWishlist handles POST /wishlist
// @Summary Create a new wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Param data body CreateWishlistRequest true "Wishlist data"
// @Success 201
// @Failure 400
// @Router /wishlist [post]
func (h *HTTPHandler) CreateWishlist(w http.ResponseWriter, r *http.Request) {
	var req CreateWishlistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	userID := uint(1) // In real scenarios, this would come from JWT claims
	if err := h.wishlist.Create(userID, req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListWishlists handles GET /wishlist
// @Summary List all wishlists for a user
// @Tags wishlist
// @Produce json
// @Success 200 {array} service.Wishlist
// @Router /wishlist [get]
func (h *HTTPHandler) ListWishlists(w http.ResponseWriter, r *http.Request) {
	userID := uint(1) // Hardcoded for now, should come from JWT
	lists, err := h.wishlist.List(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lists)
}

// DeleteWishlist handles DELETE /wishlist/{id}
// @Summary Delete a wishlist by ID
// @Tags wishlist
// @Param id path int true "Wishlist ID"
// @Success 204
// @Failure 400
// @Failure 500
// @Router /wishlist/{id} [delete]
func (h *HTTPHandler) DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	userID := uint(1) // Hardcoded for now
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

// AddBook handles POST /wishlist/{id}/books
// @Summary Add a book to the wishlist
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Wishlist ID"
// @Param data body AddBookRequest true "Book data"
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

// ListBooks handles GET /wishlist/{id}/books
// @Summary List all books from a wishlist
// @Tags books
// @Produce json
// @Param id path int true "Wishlist ID"
// @Success 200 {array} service.Book
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

// DeleteBook handles DELETE /wishlist/{id}/books/{bookID}
// @Summary Remove a book from a wishlist
// @Tags books
// @Param id path int true "Wishlist ID"
// @Param bookID path int true "Book ID"
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

//
// ───────────────────────── GOOGLE BOOKS ─────────────────────────
//

// GoogleBooksHTTP handles external Google Books searches.
type GoogleBooksHTTP struct {
	api service.GoogleBooksUsecase
}

// NewGoogleBooksHTTP builds the Google Books handler.
func NewGoogleBooksHTTP(api service.GoogleBooksUsecase) *GoogleBooksHTTP {
	return &GoogleBooksHTTP{api: api}
}

// RegisterGoogleRoutes registers Google Books API routes under /books.
func (h *GoogleBooksHTTP) RegisterGoogleRoutes(r *mux.Router) {
	r.HandleFunc("/books/search", h.SearchBooks).Methods(http.MethodGet)
}

// SearchBooks handles GET /books/search
// @Summary Search books using Google Books API
// @Tags books
// @Produce json
// @Param q query string true "Search term"
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
