package controller

import (
	"github.com/labstack/echo/v4"
	"gomvc/model"
	"gomvc/resources"
	"gomvc/service"
	"gomvc/util"
	"net/http"
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
	userEmail, err := util.GetUserEmailFromToken(e.Request().Header.Get("Authorization"))

	if err != nil {
		return err
	}

	user, err := resources.GetUser(userEmail)

	return e.JSON(http.StatusOK, b.bookService.GetBooks(user.UserType))
}

func (b *BookController) addBook(e echo.Context) error {
	userEmail, err := util.GetUserEmailFromToken(e.Request().Header.Get("Authorization"))

	if err != nil {
		return err
	}

	user, err := resources.GetUser(userEmail)

	if user.UserType != model.Admin {
		return e.JSON(http.StatusUnauthorized, "Only admin can add book")
	}

	return nil
}

func (b *BookController) deleteBook(e echo.Context) error {

	userEmail, err := util.GetUserEmailFromToken(e.Request().Header.Get("Authorization"))

	if err != nil {
		return err
	}

	user, err := resources.GetUser(userEmail)

	if user.UserType != model.Admin {
		return e.JSON(http.StatusUnauthorized, "Only admin can add book")
	}

	return nil
}
