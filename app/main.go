package main

import (
	"log"
	"time"

	_bookHandler "case-study/leo/book/delivery"
	_bookRepo "case-study/leo/book/repository"
	_bookUsecase "case-study/leo/book/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	bookRepo := _bookRepo.NewBookRepository("https://openlibrary.org/subjects/")

	timeoutContext := time.Duration(5) * time.Second
	bookUsecase := _bookUsecase.NewBookUsecase(bookRepo, timeoutContext)

	_bookHandler.NewBookHandler(e, bookUsecase)

	log.Fatal(e.Start(":8080"))
}
