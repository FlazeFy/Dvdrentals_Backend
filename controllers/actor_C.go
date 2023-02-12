package controllers

import (
	"net/http"

	"dvdrentals_backend/model"

	"github.com/labstack/echo/v4"
)

func GetActorWMostFilm(c echo.Context) error {
	result, err := model.GetActorWMostFilm()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAverageActorFilmCategory(c echo.Context) error {
	result, err := model.GetTotalAverageActorFilmCategory()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
