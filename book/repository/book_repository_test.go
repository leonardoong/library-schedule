package repository_test

import (
	"case-study/leo/book/repository"
	"case-study/leo/domain"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookRepository_GetBySubject_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Contains(t, r.URL.Path, "/test-subject.json")
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "0", r.URL.Query().Get("offset"))

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp := `{
			"work_count": 1,
			"works": [{
				"key": "/works/OL1234W",
				"title": "Test Book",
				"edition_count": 2,
				"cover_id": 12345,
				"first_publish_year": 2020,
				"authors": [{"name": "John Doe"}],
				"availability": {"available_to_borrow": true}
			}]
		}`
		w.Write([]byte(jsonResp))
	}))
	defer mockServer.Close()

	repo := repository.NewBookRepository(mockServer.URL + "/")

	req := domain.BookGetBySubjectRequest{
		Subject: "test-subject",
		Limit:   10,
		Offset:  0,
	}

	res, err := repo.GetBySubject(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), res.TotalBook)
	assert.Len(t, res.Books, 1)
	assert.Equal(t, "Test Book", res.Books[0].Title)
	assert.Equal(t, "John Doe", res.Books[0].Author[0].Name)
	assert.Equal(t, true, res.Books[0].CanBorrow)
}

func TestBookRepository_SaveCanBorrowBook_Success(t *testing.T) {
	repo := repository.NewBookRepository("")

	books := []domain.Book{
		{ID: "1", Title: "Book 1", CanBorrow: true},
		{ID: "2", Title: "Book 2", CanBorrow: true},
	}

	err := repo.SaveCanBorrowBook(context.Background(), books)

	assert.NoError(t, err)

	book, err := repo.GetCanBorrowBookByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, "Book 1", book.Title)
}

func TestBookRepository_GetCanBorrowBookByID_NotFound(t *testing.T) {
	repo := repository.NewBookRepository("")

	_, err := repo.GetCanBorrowBookByID(context.Background(), "non-existent-id")

	assert.EqualError(t, err, "book not found")
}
