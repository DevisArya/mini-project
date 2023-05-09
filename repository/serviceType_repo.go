package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type IServiceTypeService interface {
	CreateServiceType(servicetype *models.ServiceType) error
	GetServiceType(id int) (error, interface{})
	GetServiceTypeName(name string) error
	GetServiceTypes() (error, interface{})
	DeleteServiceType(id int) error
	UpdateServiceType(dataUpdate *models.ServiceType, id int) error
}

type ServiceTypeRepository struct {
	Func IServiceTypeService
}

var servicetypeRepository IServiceTypeService

func init() {
	bg := &ServiceTypeRepository{}
	bg.Func = bg

	servicetypeRepository = bg
}
func GetServiceTypeRepository() IServiceTypeService {
	return servicetypeRepository
}
func SetServiceTypeRepository(ur IServiceTypeService) {
	servicetypeRepository = ur
}

func (u *ServiceTypeRepository) CreateServiceType(servicetype *models.ServiceType) error {
	if err := config.DB.Save(&servicetype).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *ServiceTypeRepository) GetServiceType(id int) (err error, res interface{}) {
	var servicetype models.ServiceType
	if err := config.DB.Where("id = ?", id).First(&servicetype).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "servicetype not found",
		}), nil
	}
	servicetypeRes := models.ServiceTypeResponse{Name: servicetype.Name, Price: servicetype.Price}
	return nil, servicetypeRes
}
func (u *ServiceTypeRepository) GetServiceTypeName(name string) (err error) {
	var servicetype models.ServiceType
	if err := config.DB.Where("name = ?", name).First(&servicetype).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "servicetype not found",
		})
	}
	return nil
}

func (u *ServiceTypeRepository) GetServiceTypes() (err error, res interface{}) {
	var servicetypes []models.ServiceType

	if err := config.DB.Find(&servicetypes).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}

	var servicetypeResponse []models.ServiceTypeResponse

	for _, res := range servicetypes {
		servicetypeRes := models.ServiceTypeResponse{Name: res.Name, Price: res.Price}
		servicetypeResponse = append(servicetypeResponse, servicetypeRes)
	}
	return nil, servicetypeResponse
}

func (u *ServiceTypeRepository) DeleteServiceType(id int) error {
	result := config.DB.Delete(&models.ServiceType{}, id)

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

func (u *ServiceTypeRepository) UpdateServiceType(updateData *models.ServiceType, id int) error {
	result := config.DB.Model(&models.ServiceType{}).Where("id = ?", id).Updates(&updateData)

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
