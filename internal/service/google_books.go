package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GoogleBook es el modelo simplificado que devolverá tu API
type GoogleBook struct {
	Title  string   `json:"title"`
	Author []string `json:"author"`
}

// GoogleBooksUsecase define la interfaz
type GoogleBooksUsecase interface {
	Search(query string) ([]GoogleBook, error)
}

type googleBooksService struct{}

func NewGoogleBooksService() GoogleBooksUsecase {
	return &googleBooksService{}
}

// Search consulta la API de Google Books
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
