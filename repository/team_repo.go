package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type ITeamService interface {
	CreateTeam(team *models.Team) error
	GetTeam(id int) (error, interface{})
	GetTeamName(name string) error
	GetTeams() (error, interface{})
	DeleteTeam(id int) error
	UpdateTeam(dataUpdate *models.Team, id int) error
}

type TeamRepository struct {
	Func ITeamService
}

var teamRepository ITeamService

func init() {
	bg := &TeamRepository{}
	bg.Func = bg

	teamRepository = bg
}
func GetTeamRepository() ITeamService {
	return teamRepository
}
func SetTeamRepository(ur ITeamService) {
	teamRepository = ur
}

func (u *TeamRepository) CreateTeam(team *models.Team) error {
	if err := config.DB.Save(&team).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *TeamRepository) GetTeam(id int) (err error, res interface{}) {
	var team models.Team
	if err := config.DB.Preload("Cleaner").Where("id = ?", id).First(&team).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "team not found",
		}), nil
	}
	teamRes := models.TeamResponse{Name: team.Name, Cleaner: team.Cleaner}
	return nil, teamRes
}
func (u *TeamRepository) GetTeamName(name string) (err error) {
	var team models.Team
	if err := config.DB.Where("name = ?", name).First(&team).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "team not found",
		})
	}
	return nil
}

func (u *TeamRepository) GetTeams() (err error, res interface{}) {
	var teams []models.Team

	if err := config.DB.Preload("Cleaner").Find(&teams).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}

	var teamResponse []models.TeamResponse

	for _, res := range teams {
		teamRes := models.TeamResponse{Name: res.Name, Cleaner: res.Cleaner}
		teamResponse = append(teamResponse, teamRes)
	}
	return nil, teamResponse
}

func (u *TeamRepository) DeleteTeam(id int) error {
	result := config.DB.Delete(&models.Team{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}

func (u *TeamRepository) UpdateTeam(updateData *models.Team, id int) error {
	result := config.DB.Model(&models.Team{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return nil
}
