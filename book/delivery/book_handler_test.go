package delivery_test

import (
	"case-study/leo/book/delivery"
	"case-study/leo/domain"
	"case-study/leo/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBySubject_Success(t *testing.T) {
	mockUsecase := new(mocks.BookUsecase)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/books?subject=test&offset=0&limit=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockResponse := domain.BookGetBySubjectResponse{
		TotalBook: 2,
		Books: []domain.Book{
			{ID: "1", Title: "Test Book 1", CanBorrow: true},
			{ID: "2", Title: "Test Book 2", CanBorrow: false},
		},
	}

	mockUsecase.On("GetBySubject", mock.Anything, domain.BookGetBySubjectRequest{
		Subject: "test",
		Offset:  0,
		Limit:   10,
	}).Return(mockResponse, nil)

	handler := delivery.BookHandler{BookUsecase: mockUsecase}

	err := handler.GetBySubject(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{
		"total_book" : 2,
		"books": [
			{"id": "1", "title": "Test Book 1", "can_borrow": true, "author" : null, "cover_image" : "", "edition_count" : 0, "first_publish_year" : 0},
			{"id": "2", "title": "Test Book 2", "can_borrow": false, "author" : null, "cover_image" : "", "edition_count" : 0, "first_publish_year" : 0}
		]
	}`, rec.Body.String())

	mockUsecase.AssertExpectations(t)
}

func TestGetBySubject_MissingSubject(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := delivery.BookHandler{}

	err := handler.GetBySubject(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, `{"message": "Subject is required"}`, rec.Body.String())
}

func TestGetBySubject_UsecaseError(t *testing.T) {
	mockUsecase := new(mocks.BookUsecase)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/v1/books?subject=test", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase.On("GetBySubject", mock.Anything, mock.Anything).
		Return(domain.BookGetBySubjectResponse{}, domain.ErrInternalServerError).Once()

	// Initialize the handler with the mock use case
	handler := delivery.BookHandler{BookUsecase: mockUsecase}

	// Call the handler method
	err := handler.GetBySubject(c)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.JSONEq(t, `{"message": "internal Server Error"}`, rec.Body.String())

	// Ensure the mock expectations were met
	mockUsecase.AssertExpectations(t)
}
