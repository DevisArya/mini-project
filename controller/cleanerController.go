package controller

import (
	"miniproject/config"
	m "miniproject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCleaner(c echo.Context) error {

	var cleaner m.Cleaner

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := config.DB.Where("id = ?", id).First(&cleaner).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "cleaner not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get cleaner",
		"cleaner": cleaner,
	})

}

func GetCleaners(c echo.Context) error {

	var cleaners []m.Cleaner

	if err := config.DB.Find(&cleaners).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all cleaners",
		"cleaners": cleaners,
	})
}

func CreateCleaner(c echo.Context) error {

	cleaner := m.Cleaner{}
	c.Bind(&cleaner)

	if err := config.DB.Save(&cleaner).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create new cleaner",
		"cleaner": cleaner,
	})
}

func DeleteCleaner(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	result := config.DB.Delete(&m.Cleaner{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete cleaner",
	})
}

func UpdateCleaner(c echo.Context) error {
	updateData := m.Cleaner{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	result := config.DB.Model(&m.Cleaner{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update cleaner",
	})
}
