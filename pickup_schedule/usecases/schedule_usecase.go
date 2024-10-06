package usecases

import (
	"case-study/leo/domain"
	"context"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
)

type scheduleUsecase struct {
	bookRepo       domain.BookRepository
	scheduleRepo   domain.ScheduleRepository
	contextTimeout time.Duration
}

func NewScheduleUsecase(br domain.BookRepository, sr domain.ScheduleRepository, timeout time.Duration) domain.ScheduleUsecase {
	return &scheduleUsecase{
		bookRepo:       br,
		scheduleRepo:   sr,
		contextTimeout: timeout,
	}
}

func (su *scheduleUsecase) SaveSchedule(ctx context.Context, req domain.SaveScheduleRequest) (resp domain.SaveScheduleResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	// get book data from can borrow data
	bookCanBorrow, err := su.bookRepo.GetCanBorrowBookByID(ctx, req.BookID)
	if err != nil {
		log.Warnf("Failed get book to be borrow : %s", err.Error())
		return domain.SaveScheduleResponse{
			Message: fmt.Sprintf("Failed get book to be borrow : %s", err.Error()),
		}, domain.ErrNotFound
	}

	// save into schedule
	pickupDate, err := time.Parse(time.DateOnly, req.PickUpDate)
	if err != nil {
		return resp, domain.ErrInternalServerError
	}
	dueDate := pickupDate.AddDate(0, 0, 14)
	schedule := domain.Schedule{
		Book:       bookCanBorrow,
		Username:   req.Username,
		PickUpDate: req.PickUpDate,
		PickUpTime: req.PickUpTime,
		DueDate:    dueDate.Format(time.DateOnly),
	}
	err = su.scheduleRepo.SaveSchedule(ctx, schedule)
	if err != nil {
		log.Warnf("Failed get save schedule  : %s", err.Error())
		return domain.SaveScheduleResponse{
			Message: fmt.Sprintf("Failed get save schedule : %s", err.Error()),
		}, domain.ErrInternalServerError
	}

	return domain.SaveScheduleResponse{
		Message:  "Pickup scheduled created",
		Schedule: schedule,
	}, nil
}

func (su *scheduleUsecase) GetSchedules(ctx context.Context) (schedules []domain.Schedule, err error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	schedules, err = su.scheduleRepo.GetSchedules(ctx)
	if err != nil {
		return schedules, domain.ErrInternalServerError
	}

	return schedules, err
}
