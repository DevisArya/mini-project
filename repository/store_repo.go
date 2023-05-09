package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type IStoreService interface {
	CreateStore(store *models.Store) error
	GetStore(id int) (error, interface{})
	GetStoreEmail(email string) error
	GetStores() (error, interface{})
	DeleteStore(id int) error
	UpdateStore(dataUpdate *models.Store, id int) error
}

type StoreRepository struct {
	Func IStoreService
}

var storeRepository IStoreService

func init() {
	bg := &StoreRepository{}
	bg.Func = bg

	storeRepository = bg
}
func GetStoreRepository() IStoreService {
	return storeRepository
}
func SetStoreRepository(ur IStoreService) {
	storeRepository = ur
}

func (u *StoreRepository) CreateStore(store *models.Store) error {
	if err := config.DB.Save(&store).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *StoreRepository) GetStore(id int) (err error, res interface{}) {
	var store models.Store
	if err := config.DB.Preload("Cleaner").Where("id = ?", id).First(&store).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "store not found",
		}), nil
	}
	return nil, store
}
func (u *StoreRepository) GetStoreEmail(email string) (err error) {
	var store models.Store
	if err := config.DB.Where("email = ?", email).First(&store).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "store not found",
		})
	}
	return nil
}

func (u *StoreRepository) GetStores() (err error, res interface{}) {
	var stores []models.Store

	if err := config.DB.Preload("Cleaner").Find(&stores).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	return nil, stores
}

func (u *StoreRepository) DeleteStore(id int) error {
	result := config.DB.Delete(&models.Store{}, id)

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

func (u *StoreRepository) UpdateStore(updateData *models.Store, id int) error {
	result := config.DB.Model(&models.Store{}).Where("id = ?", id).Updates(&updateData)

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
