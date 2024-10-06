package usecases

import (
	"case-study/leo/domain"
	"context"
	"time"

	"github.com/labstack/gommon/log"
)

type bookUsecase struct {
	bookRepository domain.BookRepository
	contextTimeout time.Duration
}

func NewBookUsecase(b domain.BookRepository, timeout time.Duration) domain.BookUsecase {
	return &bookUsecase{
		bookRepository: b,
		contextTimeout: timeout,
	}
}

func (bu *bookUsecase) GetBySubject(ctx context.Context, req domain.BookGetBySubjectRequest) (res domain.BookGetBySubjectResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, bu.contextTimeout)
	defer cancel()

	res, err = bu.bookRepository.GetBySubject(ctx, req)
	if err != nil {
		return res, domain.ErrInternalServerError
	}

	// save book data if can_borrow true
	var canBorrowBook []domain.Book
	for _, book := range res.Books {
		if book.CanBorrow {
			canBorrowBook = append(canBorrowBook, book)
		}
	}

	if len(canBorrowBook) != 0 {
		err = bu.bookRepository.SaveCanBorrowBook(ctx, canBorrowBook)
		if err != nil {
			// still return book list, but didn't save borrow book
			log.Warnf("fail save can borrow book : %s", err.Error())
			return res, nil
		}
	}

	return
}
