package controllers

import (
	"net/http"

	"dvdrentals_backend/model"

	"github.com/labstack/echo/v4"
)

func GetAllFilm(c echo.Context) error {
	result, err := model.GetAllFilm()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetFilmWMostActor(c echo.Context) error {
	result, err := model.GetFilmWMostActor()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalFilmByRating(c echo.Context) error {
	result, err := model.GetTotalFilmByRating()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalFilmByCat(c echo.Context) error {
	result, err := model.GetTotalFilmByCat()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAverageFilmDurationByCat(c echo.Context) error {
	result, err := model.GetAverageFilmDurationByCat()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalAverageFilmReplacementCost(c echo.Context) error {
	result, err := model.GetTotalAverageFilmReplacementCost()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
