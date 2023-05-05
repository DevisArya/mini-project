package controller

import (
	"miniproject/config"
	m "miniproject/models"
	u "miniproject/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAreas(c echo.Context) error {

	var areas []m.Area

	if err := config.DB.Find(&areas).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all area",
		"areas":   areas,
	})
}

func CreateArea(c echo.Context) error {

	area := m.Area{}
	c.Bind(&area)

	valid := u.PostAreaValidation(area)
	if valid != nil {
		return echo.NewHTTPError(http.StatusBadRequest, valid.Error())
	}

	if err := config.DB.Save(&area).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create new area",
		"area":    area,
	})
}

func DeleteArea(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	result := config.DB.Delete(&m.Area{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete area",
	})
}

func UpdateArea(c echo.Context) error {
	updateData := m.Area{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	result := config.DB.Model(&m.Area{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update area",
	})
}
