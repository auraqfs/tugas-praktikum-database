package controller

import (
	"go-postgres/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func FetchAllDrama(c echo.Context) error {
	result, err := models.FetchAllDrama()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func InsertDataDrama(c echo.Context) error {
	title := c.FormValue("title")
	director := c.FormValue("director")
	release_date := c.FormValue("release_date")

	result, err := models.InsertDataDrama(title, director, release_date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDrama(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	director := c.FormValue("director")
	release_date := c.FormValue("release_date")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	result, err := models.UpdateDrama(idInt, title, director, release_date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteDrama(c echo.Context) error {
	id := c.FormValue("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	result, err := models.DeleteDrama(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}
