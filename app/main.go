package main

import (
	"log"
	"time"

	_bookHandler "case-study/leo/book/delivery"
	_bookRepo "case-study/leo/book/repository"
	_bookUsecase "case-study/leo/book/usecases"

	_scheduleHandler "case-study/leo/pickup_schedule/delivery"
	_scheduleRepo "case-study/leo/pickup_schedule/repository"
	_scheduleUsecase "case-study/leo/pickup_schedule/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	timeoutContext := time.Duration(5) * time.Second

	bookRepo := _bookRepo.NewBookRepository("https://openlibrary.org/subjects/")
	bookUsecase := _bookUsecase.NewBookUsecase(bookRepo, timeoutContext)
	_bookHandler.NewBookHandler(e, bookUsecase)

	scheduleRepo := _scheduleRepo.NewScheduleRepository()
	scheduleUsecase := _scheduleUsecase.NewScheduleUsecase(bookRepo, scheduleRepo, timeoutContext)
	_scheduleHandler.NewScheduleHandler(e, scheduleUsecase)

	log.Fatal(e.Start(":8080"))
}
