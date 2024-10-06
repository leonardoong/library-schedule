package domain

import (
	"context"
)

type Schedule struct {
	Book       Book       `json:"book"`
	Username   string     `json:"username"`
	PickUpDate string     `json:"pickup_date"`
	PickUpTime PickupTime `json:"pickup_time"`
	DueDate    string     `json:"due_date"`
}

type SaveScheduleRequest struct {
	BookID     string     `json:"book_id"`
	Username   string     `json:"username"`
	PickUpDate string     `json:"pickup_date"`
	PickUpTime PickupTime `json:"pickup_time"`
}

type SaveScheduleResponse struct {
	Message  string   `json:"message"`
	Schedule Schedule `json:"schedule"`
}

type ScheduleUsecase interface {
	SaveSchedule(ctx context.Context, req SaveScheduleRequest) (SaveScheduleResponse, error)
	GetSchedules(ctx context.Context) ([]Schedule, error)
}

type ScheduleRepository interface {
	SaveSchedule(ctx context.Context, req Schedule) error
	GetSchedules(ctx context.Context) ([]Schedule, error)
}
