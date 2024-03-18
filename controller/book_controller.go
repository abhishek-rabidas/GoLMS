package controller

import (
	"github.com/labstack/echo/v4"
	"gomvc/service"
)

type BookController struct {
	bookService *service.BookService
}

func NewBookController(e *echo.Group) *BookController {
	bookController := BookController{
		bookService: service.NewBookService(),
	}
	e.GET("/home", bookController.getBooks)
	e.POST("/addBook", bookController.addBook)
	e.DELETE("/deleteBook", bookController.deleteBook)
	return &bookController
}

func (b *BookController) getBooks(e echo.Context) error {
	return nil
}

func (b *BookController) addBook(e echo.Context) error {
	return nil
}

func (b *BookController) deleteBook(e echo.Context) error {
	return nil
}
