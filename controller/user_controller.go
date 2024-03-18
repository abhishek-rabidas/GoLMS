package controller

import (
	"github.com/labstack/echo/v4"
	"gomvc/service"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(e *echo.Group) *UserController {
	userController := UserController{userService: service.NewUserService()}
	e.POST("/login", userController.loginUser)
	return &userController
}

func (u *UserController) loginUser(e echo.Context) error {
	res, err := u.userService.LoginUser(e.QueryParam("email"), e.QueryParam("password"))
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	} else {
		return e.JSON(http.StatusOK, res)
	}

}
