package controller

import (
	"github.com/labstack/echo/v4"
	"gomvc/service"
	"gomvc/util"
	"net/http"
	"strings"
)

type Controllers struct {
	userService *service.UserService
}

func NewEchoServer() *echo.Echo {
	e := echo.New()
	setup(e)
	return e
}

func setup(e *echo.Echo) {
	mainController := Controllers{userService: service.NewUserService()}
	e.POST("/login", mainController.loginUser)
	e.POST("/validate", mainController.validateToken)

	apiController := e.Group("/api")
	apiController.Use(FilterRequest)

	NewBookController(apiController.Group("/book"))
}

func FilterRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		tokenString := e.Request().Header.Get("Authorization")

		if strings.TrimSpace(tokenString) == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please pass the token")
		}

		err, _ := util.VerifyToken(tokenString)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		} else {
			return next(e)
		}
	}
}

func (c *Controllers) loginUser(e echo.Context) error {
	res, err := c.userService.LoginUser(e.QueryParam("email"), e.QueryParam("password"))
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	} else {
		return e.JSON(http.StatusOK, res)
	}
}

func (c *Controllers) validateToken(e echo.Context) error {
	err := c.userService.ValidateToken(e.Request().Header.Get("Authorization"))

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	} else {
		return e.JSON(http.StatusOK, "Valid token")
	}
}
