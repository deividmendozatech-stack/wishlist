package service

// wishlistService is the concrete implementation of the WishlistUsecase interface.
// It contains the business logic for managing user wishlists.
type wishlistService struct {
	repo WishlistRepository
}

// NewWishlistService creates a new instance of wishlistService.
// It requires a WishlistRepository to handle persistence.
func NewWishlistService(r WishlistRepository) WishlistUsecase {
	return &wishlistService{repo: r}
}

// Create adds a new wishlist for the given user.
func (s *wishlistService) Create(userID uint, name string) error {
	w := &Wishlist{UserID: userID, Name: name}
	return s.repo.Add(w)
}

// List retrieves all wishlists that belong to the given user.
func (s *wishlistService) List(userID uint) ([]Wishlist, error) {
	return s.repo.List(userID)
}

// Delete removes a wishlist by its ID, ensuring it belongs to the given user.
func (s *wishlistService) Delete(userID, wishlistID uint) error {
	return s.repo.Delete(userID, wishlistID)
}
