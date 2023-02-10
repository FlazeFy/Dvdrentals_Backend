package controllers

import (
	"net/http"

	"dvdrentals_backend/model"

	"github.com/labstack/echo/v4"
	//"dvdrentals_backend/storage"
)

func GetAllCustomer(c echo.Context) error {
	result, err := model.GetAllCustomer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
