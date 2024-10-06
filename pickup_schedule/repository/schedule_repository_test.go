package repository_test

import (
	"case-study/leo/domain"
	"case-study/leo/pickup_schedule/repository"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduleRepository_SaveSchedulesAndGetSchedules(t *testing.T) {
	repo := repository.NewScheduleRepository()
	schedule1 := domain.Schedule{
		Book: domain.Book{
			ID: "book1",
		},
		Username:   "user1",
		PickUpDate: "2024-10-01",
		PickUpTime: "10:00",
	}

	schedule2 := domain.Schedule{
		Book: domain.Book{
			ID: "book2",
		},
		Username:   "user2",
		PickUpDate: "2024-10-02",
		PickUpTime: "11:00",
	}

	repo.SaveSchedule(context.Background(), schedule1)
	repo.SaveSchedule(context.Background(), schedule2)

	schedules, err := repo.GetSchedules(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(schedules))
	assert.Contains(t, schedules, schedule1)
	assert.Contains(t, schedules, schedule2)
}

func TestScheduleRepository_Concurrency(t *testing.T) {
	repo := repository.NewScheduleRepository()
	schedule := domain.Schedule{
		Book: domain.Book{
			ID: "book1",
		},
		Username:   "user1",
		PickUpDate: "2024-10-01",
		PickUpTime: "10:00",
	}

	for i := 0; i < 10; i++ {
		go repo.SaveSchedule(context.Background(), schedule)
	}

	time.Sleep(1 * time.Second)

	schedules, err := repo.GetSchedules(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 10, len(schedules))
}
