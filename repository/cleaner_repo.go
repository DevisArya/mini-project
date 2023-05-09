package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type ICleanerService interface {
	CreateCleaner(cleaner *models.Cleaner) error
	GetCleaner(id int) (error, interface{})
	GetCleanerEmail(email string) error
	GetCleaners() (error, interface{})
	DeleteCleaner(id int) error
	UpdateCleaner(dataUpdate *models.Cleaner, id int) error
}

type CleanerRepository struct {
	Func ICleanerService
}

var cleanerRepository ICleanerService

func init() {
	bg := &CleanerRepository{}
	bg.Func = bg

	cleanerRepository = bg
}
func GetCleanerRepository() ICleanerService {
	return cleanerRepository
}
func SetCleanerRepository(ur ICleanerService) {
	cleanerRepository = ur
}

func (u *CleanerRepository) CreateCleaner(cleaner *models.Cleaner) error {
	if err := config.DB.Save(&cleaner).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *CleanerRepository) GetCleaner(id int) (err error, res interface{}) {
	var cleaner models.Cleaner
	if err := config.DB.Where("id = ?", id).First(&cleaner).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "cleaner not found",
		}), nil
	}
	return nil, cleaner
}
func (u *CleanerRepository) GetCleanerEmail(email string) (err error) {
	var cleaner models.Cleaner
	if err := config.DB.Where("email = ?", email).First(&cleaner).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "cleaner not found",
		})
	}
	return nil
}

func (u *CleanerRepository) GetCleaners() (err error, res interface{}) {
	var cleaners []models.Cleaner

	if err := config.DB.Find(&cleaners).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	return nil, cleaners
}

func (u *CleanerRepository) DeleteCleaner(id int) error {
	result := config.DB.Delete(&models.Cleaner{}, id)

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

func (u *CleanerRepository) UpdateCleaner(updateData *models.Cleaner, id int) error {
	result := config.DB.Model(&models.Cleaner{}).Where("id = ?", id).Updates(&updateData)

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
