package controller

import (
	"github.com/labstack/echo/v4"
)

type Controllers struct {
	userController *UserController
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	setup(e)
	return e
}

func setup(e *echo.Echo) {
	apiController := e.Group("/api")
	NewUserController(apiController.Group("/user"))

}
