package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GoogleBook represents a simplified book model returned by the API.
type GoogleBook struct {
	Title  string   `json:"title"`
	Author []string `json:"author"`
}

// GoogleBooksUsecase defines book search operations via Google Books.
type GoogleBooksUsecase interface {
	Search(query string) ([]GoogleBook, error)
}

// googleBooksService implements GoogleBooksUsecase.
type googleBooksService struct{}

// NewGoogleBooksService creates a GoogleBooksUsecase instance.
func NewGoogleBooksService() GoogleBooksUsecase {
	return &googleBooksService{}
}

// Search queries the Google Books API and returns simplified results.
func (s *googleBooksService) Search(query string) ([]GoogleBook, error) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResp struct {
		Items []struct {
			VolumeInfo struct {
				Title   string   `json:"title"`
				Authors []string `json:"authors"`
			} `json:"volumeInfo"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, err
	}

	books := make([]GoogleBook, 0, len(apiResp.Items))
	for _, item := range apiResp.Items {
		books = append(books, GoogleBook{
			Title:  item.VolumeInfo.Title,
			Author: item.VolumeInfo.Authors,
		})
	}
	return books, nil
}
