package repository

import (
	"miniproject/config"
	"miniproject/models"
	"miniproject/utils"
	"net/http"

	"github.com/labstack/echo"
)

type ICustomerService interface {
	CreateCustomer(customer *models.Customer) error
	GetCustomer(id int) (error, interface{})
	GetCustomerEmail(email string) error
	GetCustomers() (error, interface{})
	DeleteCustomer(id int) error
	UpdateCustomer(dataUpdate *models.Customer, id int) error
	LoginCustomer(customer *models.Customer) error
}

type CustomerRepository struct {
	Func ICustomerService
}

var customerRepository ICustomerService

func init() {
	bg := &CustomerRepository{}
	bg.Func = bg

	customerRepository = bg
}
func GetCustomerRepository() ICustomerService {
	return customerRepository
}
func SetCustomerRepository(ur ICustomerService) {
	customerRepository = ur
}

func (u *CustomerRepository) CreateCustomer(customer *models.Customer) error {

	customer.Role = false

	password := customer.Password

	hash, err := utils.HashPassword(password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	customer.Password = hash
	if err := config.DB.Save(&customer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (u *CustomerRepository) GetCustomer(id int) (err error, res interface{}) {
	var customer models.Customer
	if err := config.DB.Preload("Transaction").Where("id = ?", id).First(&customer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "customer not found",
		}), nil
	}
	return nil, customer
}

func (u *CustomerRepository) GetCustomerEmail(email string) (err error) {
	var customer models.Customer
	if err := config.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "customer not found",
		})
	}
	return nil
}

func (u *CustomerRepository) GetCustomers() (err error, res interface{}) {
	var customers []models.Customer

	if err := config.DB.Preload("Transaction").Find(&customers).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()), nil
	}
	return nil, customers
}

func (u *CustomerRepository) DeleteCustomer(id int) error {
	result := config.DB.Delete(&models.Customer{}, id)

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

func (u *CustomerRepository) UpdateCustomer(updateData *models.Customer, id int) error {
	result := config.DB.Model(&models.Customer{}).Where("id = ?", id).Updates(&updateData)

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
func (u *CustomerRepository) LoginCustomer(customer *models.Customer) error {
	password := customer.Password

	if err := config.DB.Where("email = ?", customer.Email).First(&customer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "user not found",
			"error":   err.Error(),
		})
	}
	if match := utils.CheckPasswordHash(password, customer.Password); match == false {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "mismatch password",
		})
	}
	return nil
}
