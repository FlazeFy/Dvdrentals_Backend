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
	e.GET("api/v1/customer/all/avgtotal", controllers.GetAllCustomerTotalAvgSpending)
	e.GET("api/v1/customer/country/most", controllers.GetCountryWMostCustomer)

	//Staff
	e.GET("api/v1/staff", controllers.GetAllStaff)
	e.GET("api/v1/staff/performance", controllers.GetStaffPerformance)

	//Film
	e.GET("api/v1/film", controllers.GetAllFilm)
	e.GET("api/v1/film/total/actor", controllers.GetFilmWMostActor)
	e.GET("api/v1/film/total/rate", controllers.GetTotalFilmByRating)
	e.GET("api/v1/film/total/category", controllers.GetTotalFilmByCat)
	e.GET("api/v1/film/duration/category", controllers.GetAverageFilmDurationByCat)

	//Payment
	e.GET("api/v1/payment", controllers.GetAllPayment)
	e.GET("api/v1/payment/total/amount", controllers.GetTotalTransactionByAmount)

	return e
}
