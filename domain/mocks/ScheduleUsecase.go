// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	domain "case-study/leo/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ScheduleUsecase is an autogenerated mock type for the ScheduleUsecase type
type ScheduleUsecase struct {
	mock.Mock
}

// GetSchedules provides a mock function with given fields: ctx
func (_m *ScheduleUsecase) GetSchedules(ctx context.Context) ([]domain.Schedule, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSchedules")
	}

	var r0 []domain.Schedule
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Schedule, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Schedule); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Schedule)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveSchedule provides a mock function with given fields: ctx, req
func (_m *ScheduleUsecase) SaveSchedule(ctx context.Context, req domain.SaveScheduleRequest) (domain.SaveScheduleResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for SaveSchedule")
	}

	var r0 domain.SaveScheduleResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.SaveScheduleRequest) (domain.SaveScheduleResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.SaveScheduleRequest) domain.SaveScheduleResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.SaveScheduleResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.SaveScheduleRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewScheduleUsecase creates a new instance of ScheduleUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScheduleUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScheduleUsecase {
	mock := &ScheduleUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
