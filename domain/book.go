package domain

import (
	"context"
)

type Book struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	Author           []Author `json:"author"`
	EditionCount     int64    `json:"edition_count"`
	FirstPublishYear int64    `json:"first_publish_year"`
	CoverImage       string   `json:"cover_image"`
	CanBorrow        bool     `json:"can_borrow"`
}

type BookGetBySubjectRequest struct {
	Subject string
	Offset  int64
	Limit   int64
}

type BookGetBySubjectResponse struct {
	TotalBook int64  `json:"total_book"`
	Books     []Book `json:"books"`
}
type BookUsecase interface {
	GetBySubject(ctx context.Context, req BookGetBySubjectRequest) (BookGetBySubjectResponse, error)
}

type BookRepository interface {
	GetBySubject(ctx context.Context, req BookGetBySubjectRequest) (BookGetBySubjectResponse, error)
	SaveCanBorrowBook(ctx context.Context, books []Book) error
	GetCanBorrowBookByID(ctx context.Context, id string) (Book, error)
}
