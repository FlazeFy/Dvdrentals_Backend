package routes

import (
	"net/http"

	"dvdrentals_backend/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to dvdrentals analytic")
	})

	//================== Handler ==================//
	//Customer
	e.GET("api/v1/customer", controllers.GetAllCustomer)

	//Staff
	e.GET("api/v1/staff", controllers.GetAllStaff)

	//Film
	e.GET("api/v1/film", controllers.GetAllFilm)

	//Payment
	e.GET("api/v1/payment", controllers.GetAllPayment)

	return e
}
