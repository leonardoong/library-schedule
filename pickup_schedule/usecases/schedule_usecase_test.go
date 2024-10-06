package usecases_test

import (
	"case-study/leo/domain"
	"case-study/leo/domain/mocks"
	"case-study/leo/pickup_schedule/usecases"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveSchedule_Success(t *testing.T) {
	bookRepoMock := new(mocks.BookRepository)
	scheduleRepoMock := new(mocks.ScheduleRepository)

	book := domain.Book{ID: "book_id", Title: "Test Book"}
	bookRepoMock.On("GetCanBorrowBookByID", mock.Anything, "book_id").Return(book, nil)
	scheduleRepoMock.On("SaveSchedule", mock.Anything, mock.Anything).Return(nil)

	scheduleUsecase := usecases.NewScheduleUsecase(bookRepoMock, scheduleRepoMock, time.Second*10)

	resp, err := scheduleUsecase.SaveSchedule(context.Background(), domain.SaveScheduleRequest{
		BookID:     "book_id",
		Username:   "test_user",
		PickUpDate: "2024-10-01",
		PickUpTime: "09:00",
	})

	assert.NoError(t, err)
	assert.Equal(t, "Pickup scheduled created", resp.Message)
	assert.Equal(t, "2024-10-01", resp.Schedule.PickUpDate)
	assert.Equal(t, "2024-10-15", resp.Schedule.DueDate)
	assert.Equal(t, "book_id", resp.Schedule.Book.ID)

	bookRepoMock.AssertExpectations(t)
	scheduleRepoMock.AssertExpectations(t)
}

func TestSaveSchedule_BookNotFound(t *testing.T) {
	bookRepoMock := new(mocks.BookRepository)
	scheduleRepoMock := new(mocks.ScheduleRepository)

	bookRepoMock.On("GetCanBorrowBookByID", mock.Anything, "book_id").Return(domain.Book{}, domain.ErrNotFound)

	scheduleUsecase := usecases.NewScheduleUsecase(bookRepoMock, scheduleRepoMock, time.Second*10)

	resp, err := scheduleUsecase.SaveSchedule(context.Background(), domain.SaveScheduleRequest{
		BookID:     "book_id",
		Username:   "test_user",
		PickUpDate: "2024-10-01",
		PickUpTime: "09:00",
	})

	assert.Error(t, err)
	assert.Equal(t, "Failed get book to be borrow : your requested Item is not found", resp.Message)

	// Verify mocks
	bookRepoMock.AssertExpectations(t)
}

func TestGetSchedules_Success(t *testing.T) {
	bookRepoMock := new(mocks.BookRepository)
	scheduleRepoMock := new(mocks.ScheduleRepository)

	schedules := []domain.Schedule{
		{Username: "user1", PickUpDate: "2024-10-01"},
		{Username: "user2", PickUpDate: "2024-10-02"},
	}
	scheduleRepoMock.On("GetSchedules", mock.Anything).Return(schedules, nil)

	scheduleUsecase := usecases.NewScheduleUsecase(bookRepoMock, scheduleRepoMock, time.Second*10)

	resp, err := scheduleUsecase.GetSchedules(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, schedules, resp)

	scheduleRepoMock.AssertExpectations(t)
}
