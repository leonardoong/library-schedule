package repository

import (
	"case-study/leo/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

type bookRepository struct {
	OpenLibUrl string
	storage    map[string]domain.Book
	mu         sync.RWMutex
}

type openLibraryResponse struct {
	WorkCount int64 `json:"work_count"`
	Works     []struct {
		Key              string `json:"key"`
		Title            string `json:"title"`
		EditionCount     int64  `json:"edition_count"`
		CoverID          int64  `json:"cover_id"`
		FirstPublishYear int64  `json:"first_publish_year"`
		Authors          []struct {
			Name string `json:"name"`
		} `json:"authors"`
		Availability struct {
			AvailableToBorrow bool `json:"available_to_borrow"`
		} `json:"availability,omitempty"`
	} `json:"works"`
}

func NewBookRepository(url string) domain.BookRepository {
	return &bookRepository{
		OpenLibUrl: url,
		storage:    make(map[string]domain.Book),
	}
}

func (br *bookRepository) GetBySubject(ctx context.Context, req domain.BookGetBySubjectRequest) (res domain.BookGetBySubjectResponse, err error) {
	url := fmt.Sprintf("%s%s.json?limit=%d&offset=%d", br.OpenLibUrl, req.Subject, req.Limit, req.Offset)

	resp, err := http.Get(url)
	if err != nil {
		return res, fmt.Errorf("failed to fetch data from Open Library: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return res, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	var result openLibraryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return res, fmt.Errorf("failed to decode Open Library response: %w", err)
	}

	var books []domain.Book
	for _, work := range result.Works {
		authors := []domain.Author{}
		for _, author := range work.Authors {
			authors = append(authors, domain.Author{
				Name: author.Name,
			})
		}

		coverImageURL := ""
		if work.CoverID > 0 {
			coverImageURL = fmt.Sprintf("https://covers.openlibrary.org/b/id/%d-L.jpg", work.CoverID)
		}

		books = append(books, domain.Book{
			ID:               work.Key,
			Title:            work.Title,
			Author:           authors,
			EditionCount:     work.EditionCount,
			FirstPublishYear: work.FirstPublishYear,
			CanBorrow:        work.Availability.AvailableToBorrow,
			CoverImage:       coverImageURL,
		})
	}

	return domain.BookGetBySubjectResponse{
		TotalBook: result.WorkCount,
		Books:     books,
	}, err
}

func (br *bookRepository) SaveCanBorrowBook(ctx context.Context, books []domain.Book) (err error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	for _, book := range books {
		br.storage[book.ID] = book
	}

	return
}

func (br *bookRepository) GetCanBorrowBookByID(ctx context.Context, id string) (res domain.Book, err error) {
	br.mu.Lock()
	defer br.mu.Unlock()

	book, exists := br.storage[id]
	if !exists {
		return domain.Book{}, errors.New("book not found")
	}

	return book, err
}
