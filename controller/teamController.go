package controller

import (
	"miniproject/config"
	m "miniproject/models"
	u "miniproject/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTeam(c echo.Context) error {

	var team m.Team

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := config.DB.Where("id = ?", id).First(&team).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "team not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get team",
		"team":    team,
	})

}

func GetTeams(c echo.Context) error {

	var teams []m.Team

	if err := config.DB.Find(&teams).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all teams",
		"teams":   teams,
	})
}

func CreateTeam(c echo.Context) error {

	team := m.Team{}
	c.Bind(&team)

	valid := u.PostTeamValidation(team)
	if valid != nil {
		return echo.NewHTTPError(http.StatusBadRequest, valid.Error())
	}

	if err := config.DB.Save(&team).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes create new team",
		"team":    team,
	})
}

func DeleteTeam(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	result := config.DB.Delete(&m.Team{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete team",
	})
}

func UpdateTeam(c echo.Context) error {
	updateData := m.Team{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	result := config.DB.Model(&m.Team{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update team",
	})
}
