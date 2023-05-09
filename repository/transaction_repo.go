package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type ITransactionService interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransaction(id int) (error, interface{})
	GetTransactions() (error, interface{})
	DeleteTransaction(id int) error
	UpdateTransaction(dataUpdate *models.Transaction, id int) error
	UpdateRating(dataUpdate *models.TransactionUpdateRating, id int) error
	UpdateStatus(dataUpdate *models.TransactionUpdateStatus, id int) error
}

type TransactionRepository struct {
	Func ITransactionService
}

var transactionRepository ITransactionService

func init() {
	bg := &TransactionRepository{}
	bg.Func = bg

	transactionRepository = bg
}
func GetTransactionRepository() ITransactionService {
	return transactionRepository
}
func SetTransactionRepository(ur ITransactionService) {
	transactionRepository = ur
}

func (u *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	if err := config.DB.Save(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *TransactionRepository) GetTransaction(id int) (err error, res interface{}) {
	var transaction models.Transaction
	if err := config.DB.Preload("TransactionDetails").Where("id = ?", id).First(&transaction).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "transaction not found",
		}), nil
	}
	return nil, transaction
}
func (u *TransactionRepository) GetTransactions() (err error, res interface{}) {
	var transactions []models.Transaction

	if err := config.DB.Preload("TransactionDetails").Find(&transactions).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()), nil
	}

	return nil, transactions
}

func (u *TransactionRepository) DeleteTransaction(id int) error {
	result := config.DB.Delete(&models.Transaction{}, id)

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

func (u *TransactionRepository) UpdateTransaction(updateData *models.Transaction, id int) error {
	result := config.DB.Model(&models.Transaction{}).Where("id = ?", id).Updates(&updateData)

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
func (u *TransactionRepository) UpdateRating(updateData *models.TransactionUpdateRating, id int) error {
	result := config.DB.Model(&models.Transaction{}).Where("id = ?", id).Updates(&updateData)

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
func (u *TransactionRepository) UpdateStatus(updateData *models.TransactionUpdateStatus, id int) error {
	result := config.DB.Model(&models.Transaction{}).Where("id = ?", id).Updates(&updateData)

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
