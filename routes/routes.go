package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to dvdrentals analytic")
	})

	return e
}
