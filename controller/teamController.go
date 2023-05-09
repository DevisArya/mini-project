package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTeam(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	err, res := repository.GetTeamRepository().GetTeam(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get team",
		"Team":    res,
	})

}

func GetTeams(c echo.Context) error {

	err, res := repository.GetTeamRepository().GetTeams()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success get all teams",
		"Teams":   res,
	})
}

func CreateTeam(c echo.Context) error {

	team := m.Team{}
	c.Bind(&team)

	if err := repository.GetTeamRepository().GetTeamName(team.Name); err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Team Already Exists",
		})
	}

	valid := md.PostTeamValidation(team)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetTeamRepository().CreateTeam(&team); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.TeamResponse{
		Id:   team.Id,
		Name: team.Name,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "succes create new team",
		"Team":    result,
	})
}

func DeleteTeam(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetTeamRepository().DeleteTeam(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete team",
	})
}

func UpdateTeam(c echo.Context) error {
	updateData := m.Team{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetTeamRepository().UpdateTeam(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update team",
	})
}
