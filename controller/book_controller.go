package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gomvc/exception"
	"gomvc/model"
	"gomvc/resources"
	"gomvc/service"
	"gomvc/util"
	"gomvc/views"
	"io"
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
		return e.JSON(http.StatusUnauthorized, exception.NewExceptionResponse("Only Admins can add book"))
	}

	req := e.Request().Body

	var payload views.BookDTO

	payloadBytes, err := io.ReadAll(req)

	if err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewExceptionResponse("Payload is incorrect"))
	}

	json.Unmarshal(payloadBytes, &payload)

	err = b.bookService.AddBook(payload)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, exception.NewExceptionResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, e.JSON(http.StatusOK, "Book Added"))
}

func (b *BookController) deleteBook(e echo.Context) error {

	userEmail, err := util.GetUserEmailFromToken(e.Request().Header.Get("Authorization"))

	if err != nil {
		return err
	}

	user, err := resources.GetUser(userEmail)

	if user.UserType != model.Admin {
		return e.JSON(http.StatusUnauthorized, exception.NewExceptionResponse("Only admins can delete book"))
	}

	return nil
}
