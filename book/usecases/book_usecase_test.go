package usecases_test

import (
	"case-study/leo/book/usecases"
	"case-study/leo/domain"
	"case-study/leo/domain/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBySubject(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBookRepoResp := domain.BookGetBySubjectResponse{
		TotalBook: 100,
		Books: []domain.Book{
			{
				ID:    "id_1",
				Title: "title_1",
				Author: []domain.Author{
					{
						Name: "author_1",
					},
				},
				EditionCount:     1,
				FirstPublishYear: 2024,
				CoverImage:       "url",
				CanBorrow:        true,
			},
			{
				ID:    "id_2",
				Title: "title_2",
				Author: []domain.Author{
					{
						Name: "author_2",
					},
				},
				EditionCount:     1,
				FirstPublishYear: 2024,
				CoverImage:       "url",
				CanBorrow:        false,
			},
		},
	}

	mockReqBySubject := domain.BookGetBySubjectRequest{
		Subject: "test",
		Offset:  0,
		Limit:   10,
	}

	bookUsecase := usecases.NewBookUsecase(mockBookRepo, 2*time.Second)

	expectedResponse := mockBookRepoResp

	mockBookRepo.On("GetBySubject", mock.Anything, mockReqBySubject).Return(expectedResponse, nil).Once()

	mockBookRepo.On("SaveCanBorrowBook", mock.Anything, []domain.Book{expectedResponse.Books[0]}).Return(nil).Once()

	res, err := bookUsecase.GetBySubject(context.Background(), mockReqBySubject)

	assert.NoError(t, err)
	assert.Equal(t, res, mockBookRepoResp)

	mockBookRepo.AssertExpectations(t)
}

func TestGetBySubject_GetBySubjectError(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)

	mockReqBySubject := domain.BookGetBySubjectRequest{
		Subject: "test",
		Offset:  0,
		Limit:   10,
	}

	bookUsecase := usecases.NewBookUsecase(mockBookRepo, 2*time.Second)

	expectedResponse := domain.BookGetBySubjectResponse{}

	mockBookRepo.On("GetBySubject", mock.Anything, mockReqBySubject).Return(domain.BookGetBySubjectResponse{}, assert.AnError).Once()

	res, err := bookUsecase.GetBySubject(context.Background(), mockReqBySubject)

	assert.Error(t, err)
	assert.Equal(t, res, expectedResponse)

	mockBookRepo.AssertExpectations(t)
}

func TestGetBySubject_SaveCanBorrowBookIgnoreError(t *testing.T) {
	mockBookRepo := new(mocks.BookRepository)
	mockBookRepoResp := domain.BookGetBySubjectResponse{
		TotalBook: 100,
		Books: []domain.Book{
			{
				ID:    "id_1",
				Title: "title_1",
				Author: []domain.Author{
					{
						Name: "author_1",
					},
				},
				EditionCount:     1,
				FirstPublishYear: 2024,
				CoverImage:       "url",
				CanBorrow:        true,
			},
			{
				ID:    "id_2",
				Title: "title_2",
				Author: []domain.Author{
					{
						Name: "author_2",
					},
				},
				EditionCount:     1,
				FirstPublishYear: 2024,
				CoverImage:       "url",
				CanBorrow:        false,
			},
		},
	}

	mockReqBySubject := domain.BookGetBySubjectRequest{
		Subject: "test",
		Offset:  0,
		Limit:   10,
	}

	bookUsecase := usecases.NewBookUsecase(mockBookRepo, 2*time.Second)

	expectedResponse := mockBookRepoResp

	mockBookRepo.On("GetBySubject", mock.Anything, mockReqBySubject).Return(expectedResponse, nil).Once()

	mockBookRepo.On("SaveCanBorrowBook", mock.Anything, []domain.Book{expectedResponse.Books[0]}).Return(assert.AnError).Once()

	res, err := bookUsecase.GetBySubject(context.Background(), mockReqBySubject)

	assert.NoError(t, err)
	assert.Equal(t, res, mockBookRepoResp)

	mockBookRepo.AssertExpectations(t)
}
