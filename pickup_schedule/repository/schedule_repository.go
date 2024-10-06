package repository

import (
	"case-study/leo/domain"
	"context"
	"sync"
)

type scheduleRepository struct {
	scheduleList []domain.Schedule
	mu           sync.RWMutex
}

func NewScheduleRepository() domain.ScheduleRepository {
	return &scheduleRepository{
		scheduleList: []domain.Schedule{},
	}
}

func (sr *scheduleRepository) SaveSchedule(ctx context.Context, schedule domain.Schedule) (err error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.scheduleList = append(sr.scheduleList, schedule)

	return
}

func (sr *scheduleRepository) GetSchedules(ctx context.Context) ([]domain.Schedule, error) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	return sr.scheduleList, nil
}
