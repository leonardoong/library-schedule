package delivery

import (
	"fmt"
	"net/http"

	"case-study/leo/domain"
	"case-study/leo/util"

	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	ScheduleUsecase domain.ScheduleUsecase
}

func NewScheduleHandler(e *echo.Echo, us domain.ScheduleUsecase) {
	handler := &ScheduleHandler{
		ScheduleUsecase: us,
	}

	e.POST("/v1/pickup_schedule", handler.SaveSchedule)
	e.GET("/v1/pickup_schedule", handler.GetSchedules)
}

func (b *ScheduleHandler) SaveSchedule(c echo.Context) error {
	ctx := c.Request().Context()

	var req domain.SaveScheduleRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseError{Message: "Book ID, Username, Pick up date and time are required."})
	}

	if !domain.IsValidPickupTime(req.PickUpTime) {
		return c.JSON(http.StatusBadRequest, util.ResponseError{
			Message: fmt.Sprintf("Invalid pickup time, please select a valid time. Pick time from this time slot : %s",
				domain.AllValidPickUpTime()),
		})
	}

	resp, err := b.ScheduleUsecase.SaveSchedule(ctx, req)
	if err != nil {
		return c.JSON(util.GetStatusCode(err), util.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}

func (b *ScheduleHandler) GetSchedules(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := b.ScheduleUsecase.GetSchedules(ctx)
	if err != nil {
		return c.JSON(util.GetStatusCode(err), util.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
