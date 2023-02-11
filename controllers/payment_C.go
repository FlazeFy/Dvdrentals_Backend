package controllers

import (
	"net/http"

	"dvdrentals_backend/model"

	"github.com/labstack/echo/v4"
)

func GetAllPayment(c echo.Context) error {
	result, err := model.GetAllPayment()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalTransactionByAmount(c echo.Context) error {
	result, err := model.GetTotalTransactionByAmount()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
