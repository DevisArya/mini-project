package controller

import (
	"miniproject/config"
	md "miniproject/middleware"
	m "miniproject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCustomer(c echo.Context) error {

	var customer m.Customer

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	if err := config.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "customer not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get customer",
		"customer": customer,
	})

}

func GetCustomers(c echo.Context) error {

	var customers []m.Customer

	if err := config.DB.Find(&customers).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get all customers",
		"customers": customers,
	})
}

func CreateCustomer(c echo.Context) error {

	customer := m.Customer{}
	c.Bind(&customer)

	if err := config.DB.Save(&customer).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "succes create new customer",
		"customer": customer,
	})
}

func DeleteCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "invalid id",
		})
	}

	result := config.DB.Delete(&m.Customer{}, id)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete customer",
	})
}

func UpdateCustomer(c echo.Context) error {
	updateData := m.Customer{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	result := config.DB.Model(&m.Customer{}).Where("id = ?", id).Updates(&updateData)

	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if result.RowsAffected < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "id not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update customer",
	})
}

func LoginCustomerController(c echo.Context) error {
	customer := m.Customer{}
	c.Bind(&customer)

	if err := config.DB.Where("email = ? AND password = ?", customer.Email, customer.Password).First(&customer).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := md.CreateToken(int(customer.ID), customer.Name)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	customerResponse := m.CustomerResponse{ID: int(customer.ID), Name: customer.Name, Email: customer.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success login",
		"customer": customerResponse,
	})
}
