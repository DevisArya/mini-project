package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type ITransactionDetailService interface {
	CreateTransactionDetail(transactiondetail *models.TransactionDetail) error
	GetTransactionDetail(id int) (error, interface{})
	GetTransactionDetails() (error, interface{})
	DeleteTransactionDetail(id int) error
	UpdateTransactionDetail(dataUpdate *models.TransactionDetail, id int) error
}

type TransactionDetailRepository struct {
	Func ITransactionDetailService
}

var transactiondetailRepository ITransactionDetailService

func init() {
	bg := &TransactionDetailRepository{}
	bg.Func = bg

	transactiondetailRepository = bg
}
func GetTransactionDetailRepository() ITransactionDetailService {
	return transactiondetailRepository
}
func SetTransactionDetailRepository(ur ITransactionDetailService) {
	transactiondetailRepository = ur
}

func (u *TransactionDetailRepository) CreateTransactionDetail(transactiondetail *models.TransactionDetail) error {
	if err := config.DB.Save(&transactiondetail).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *TransactionDetailRepository) GetTransactionDetail(id int) (err error, res interface{}) {
	var transactiondetail models.TransactionDetail
	if err := config.DB.Where("id = ?", id).First(&transactiondetail).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "transactiondetail not found",
		}), nil
	}
	return nil, transactiondetail
}
func (u *TransactionDetailRepository) GetTransactionDetails() (err error, res interface{}) {
	var transactiondetails []models.TransactionDetail

	if err := config.DB.Find(&transactiondetails).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	return nil, transactiondetails
}

func (u *TransactionDetailRepository) DeleteTransactionDetail(id int) error {
	result := config.DB.Delete(&models.TransactionDetail{}, id)

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

func (u *TransactionDetailRepository) UpdateTransactionDetail(updateData *models.TransactionDetail, id int) error {
	result := config.DB.Model(&models.TransactionDetail{}).Where("id = ?", id).Updates(&updateData)

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
