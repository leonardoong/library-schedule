package delivery

import (
	"net/http"
	"strconv"

	"case-study/leo/domain"
	"case-study/leo/util"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookUsecase domain.BookUsecase
}

func NewBookHandler(e *echo.Echo, us domain.BookUsecase) {
	handler := &BookHandler{
		BookUsecase: us,
	}

	e.GET("/v1/books", handler.GetBySubject)
}

func (b *BookHandler) GetBySubject(c echo.Context) error {
	ctx := c.Request().Context()

	subject := c.QueryParam("subject")

	if subject == "" {
		return c.JSON(http.StatusBadRequest, util.ResponseError{Message: "Subject is required"})
	}

	// set default value for offset and limit
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}

	book, err := b.BookUsecase.GetBySubject(ctx, domain.BookGetBySubjectRequest{
		Subject: subject,
		Offset:  int64(offset),
		Limit:   int64(limit),
	})
	if err != nil {
		return c.JSON(util.GetStatusCode(err), util.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, book)
}
