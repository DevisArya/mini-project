package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type IAreaService interface {
	CreateArea(area *models.Area) error
	GetAreaName(name string) error
	GetAreas() (error, interface{})
	DeleteArea(id int) error
	UpdateArea(dataUpdate *models.Area, id int) error
}

type AreaRepository struct {
	Func IAreaService
}

var areaRepository IAreaService

func init() {
	bg := &AreaRepository{}
	bg.Func = bg

	areaRepository = bg
}
func GetAreaRepository() IAreaService {
	return areaRepository
}
func SetAreaRepository(ur IAreaService) {
	areaRepository = ur
}

func (u *AreaRepository) CreateArea(area *models.Area) error {
	if err := config.DB.Save(&area).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *AreaRepository) GetAreaName(name string) (err error) {
	var area models.Area
	if err := config.DB.Where("name = ?", name).First(&area).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "area not found",
		})
	}
	return nil
}

func (u *AreaRepository) GetAreas() (err error, res interface{}) {
	var areas []models.Area

	if err := config.DB.Preload("Store").Find(&areas).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	var areaResponse []models.AreaResponseStore

	// for i, res := range areas {
	// 	areaRes := models.AreaResponse{Id: res.Id, Name: res.Name, Store: res.Store}
	// 	areaResponse = append(areaResponse, areaRes)
	// }

	// log.Println(areas[0].Store[0].Id)

	for _, res := range areas {
		if len(res.Store) == 0 {
			storeResponse := []models.StoreResponse{}
			areaRes := models.AreaResponseStore{Id: res.Id, Name: res.Name, Store: storeResponse}
			areaResponse = append(areaResponse, areaRes)
		}
		for _, val := range res.Store {
			var storeResponse []models.StoreResponse
			store := models.StoreResponse{
				Id:      val.Id,
				AreaId:  val.AreaId,
				Address: val.Address,
				Phone:   val.Phone,
				Email:   val.Email,
			}
			storeResponse = append(storeResponse, store)
			areaRes := models.AreaResponseStore{Id: res.Id, Name: res.Name, Store: storeResponse}
			areaResponse = append(areaResponse, areaRes)
		}
	}

	return nil, areaResponse
}

func (u *AreaRepository) DeleteArea(id int) error {
	result := config.DB.Delete(&models.Area{}, id)

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

func (u *AreaRepository) UpdateArea(updateData *models.Area, id int) error {
	result := config.DB.Model(&models.Area{}).Where("id = ?", id).Updates(&updateData)

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
