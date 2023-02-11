package controllers

import (
	"net/http"

	"dvdrentals_backend/model"

	"github.com/labstack/echo/v4"
)

func GetAllStaff(c echo.Context) error {
	result, err := model.GetAllStaff()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetStaffPerformance(c echo.Context) error {
	result, err := model.GetStaffPerformance()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"map": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
