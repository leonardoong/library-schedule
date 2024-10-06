package delivery_test

import (
	"case-study/leo/domain"
	"case-study/leo/domain/mocks"
	"case-study/leo/pickup_schedule/delivery"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveSchedule_Success(t *testing.T) {
	e := echo.New()

	mockUsecase := new(mocks.ScheduleUsecase)
	mockUsecaseResp := domain.SaveScheduleResponse{Message: "Schedule saved", Schedule: domain.Schedule{
		Book: domain.Book{
			ID: "123",
		},
		Username:   "user1",
		PickUpDate: "2024-10-01",
		PickUpTime: "10:00 - 11:00",
		DueDate:    "2024-10-15",
	}}
	mockUsecase.On("SaveSchedule", mock.Anything, mock.AnythingOfType("domain.SaveScheduleRequest")).
		Return(mockUsecaseResp, nil)

	handler := delivery.ScheduleHandler{ScheduleUsecase: mockUsecase}

	reqBody := `{"book_id":"123", "username":"user1", "pickup_date":"2024-10-01", "pickup_time":"10:00 - 11:00"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/pickup_schedule", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.SaveSchedule(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"message":"Schedule saved","schedule":{"book":{"id":"123","title":"","author":null,"edition_count":0,"first_publish_year":0,"cover_image":"","can_borrow":false},"username":"user1","pickup_date":"2024-10-01","pickup_time":"10:00 - 11:00","due_date":"2024-10-15"}}`, rec.Body.String())
}

func TestSaveSchedule_InvalidPickupTime(t *testing.T) {
	e := echo.New()

	mockUsecase := new(mocks.ScheduleUsecase)

	handler := delivery.ScheduleHandler{ScheduleUsecase: mockUsecase}

	reqBody := `{"book_id":"123", "username":"user1", "pickup_date":"2024-10-01", "pickup_time":"00:00"}`
	req := httptest.NewRequest(http.MethodPost, "/v1/pickup_schedule", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.SaveSchedule(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Invalid pickup time")
}

func TestGetSchedules_Success(t *testing.T) {
	e := echo.New()

	mockUsecase := new(mocks.ScheduleUsecase)

	mockSchedulesResp := []domain.Schedule{{Book: domain.Book{ID: "book1"}, Username: "user1", PickUpDate: "2024-10-01", PickUpTime: "10:00 - 11:00"}}
	b, _ := json.Marshal(mockSchedulesResp)
	fmt.Println(string(b))

	mockUsecase.On("GetSchedules", mock.Anything).
		Return(mockSchedulesResp, nil)

	handler := delivery.ScheduleHandler{ScheduleUsecase: mockUsecase}

	req := httptest.NewRequest(http.MethodGet, "/v1/pickup_schedule", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.GetSchedules(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"book":{"id":"book1","title":"","author":null,"edition_count":0,"first_publish_year":0,"cover_image":"","can_borrow":false},"username":"user1","pickup_date":"2024-10-01","pickup_time":"10:00 - 11:00","due_date":""}]`, rec.Body.String())

	mockUsecase.AssertExpectations(t)
}

func TestScheduleHandler_GetSchedules_Error(t *testing.T) {
	e := echo.New()

	mockUsecase := new(mocks.ScheduleUsecase)

	mockUsecase.On("GetSchedules", mock.Anything).
		Return([]domain.Schedule{}, errors.New("error"))

	handler := delivery.ScheduleHandler{ScheduleUsecase: mockUsecase}

	req := httptest.NewRequest(http.MethodGet, "/v1/pickup_schedule", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler.GetSchedules(c)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "error")

	// Assert that the expectations were met
	mockUsecase.AssertExpectations(t)
}
