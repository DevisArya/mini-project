package repository

import (
	"miniproject/config"
	"miniproject/models"
	"net/http"

	"github.com/labstack/echo"
)

type IPaymentService interface {
	CreatePayment(payment *models.Payment) error
	GetPayment(id int) (error, interface{})
	GetPaymentName(name string) error
	GetPayments() (error, interface{})
	DeletePayment(id int) error
	UpdatePayment(dataUpdate *models.Payment, id int) error
}

type PaymentRepository struct {
	Func IPaymentService
}

var paymentRepository IPaymentService

func init() {
	bg := &PaymentRepository{}
	bg.Func = bg

	paymentRepository = bg
}
func GetPaymentRepository() IPaymentService {
	return paymentRepository
}
func SetPaymentRepository(ur IPaymentService) {
	paymentRepository = ur
}

func (u *PaymentRepository) CreatePayment(payment *models.Payment) error {
	if err := config.DB.Save(&payment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *PaymentRepository) GetPayment(id int) (err error, res interface{}) {
	var payment models.Payment
	if err := config.DB.Where("id = ?", id).First(&payment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "payment not found",
		}), nil
	}

	paymentRes := models.PaymentResponse{Id: payment.Id, Name: payment.Name, PaymentType: payment.PaymentType, PaymentNumber: payment.PaymentNumber}
	return nil, paymentRes
}
func (u *PaymentRepository) GetPaymentName(name string) (err error) {
	var area models.Payment
	if err := config.DB.Where("name = ?", name).First(&area).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "area not found",
		})
	}
	return nil
}

func (u *PaymentRepository) GetPayments() (err error, res interface{}) {
	var payments []models.Payment

	if err := config.DB.Find(&payments).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	var paymentResponse []models.PaymentResponse

	for _, res := range payments {
		paymentRes := models.PaymentResponse{Id: res.Id, Name: res.Name, PaymentType: res.PaymentType, PaymentNumber: res.PaymentNumber}
		paymentResponse = append(paymentResponse, paymentRes)
	}

	return nil, paymentResponse
}

func (u *PaymentRepository) DeletePayment(id int) error {
	result := config.DB.Delete(&models.Payment{}, id)

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

func (u *PaymentRepository) UpdatePayment(updateData *models.Payment, id int) error {
	result := config.DB.Model(&models.Payment{}).Where("id = ?", id).Updates(&updateData)

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
