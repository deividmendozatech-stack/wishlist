package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//
// ─────────────────────────── GOOGLE BOOKS SERVICE ───────────────────────────
//

// GoogleBook represents a simplified result from Google Books API.
type GoogleBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// googleBooksService implements the GoogleBooksUsecase interface.
// It provides methods to search for books using the Google Books API.
type googleBooksService struct{}

// NewGoogleBooksService creates a new instance of googleBooksService.
func NewGoogleBooksService() GoogleBooksUsecase {
	return &googleBooksService{}
}

// Search queries the Google Books API with the given search term
// and returns a simplified list of books (title and first author).
func (g *googleBooksService) Search(query string) ([]GoogleBook, error) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s", query)

	// Send HTTP request to Google Books API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response JSON into a simplified structure
	var data struct {
		Items []struct {
			VolumeInfo struct {
				Title   string   `json:"title"`
				Authors []string `json:"authors"`
			} `json:"volumeInfo"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// Extract relevant data: title and first author
	var books []GoogleBook
	for _, item := range data.Items {
		author := ""
		if len(item.VolumeInfo.Authors) > 0 {
			author = item.VolumeInfo.Authors[0]
		}
		books = append(books, GoogleBook{
			Title:  item.VolumeInfo.Title,
			Author: author,
		})
	}

	return books, nil
}
