package controller

import (
	md "miniproject/middleware"
	m "miniproject/models"
	"miniproject/repository"
	u "miniproject/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCustomer(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	err, res := repository.GetCustomerRepository().GetCustomer(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":   "200",
		"Message":  "success get customer",
		"Customer": res,
	})

}

func GetCustomers(c echo.Context) error {

	err, res := repository.GetCustomerRepository().GetCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":    "200",
		"Message":   "success get all customers",
		"Customers": res,
	})
}

func CreateCustomer(c echo.Context) error {

	customer := m.Customer{}
	c.Bind(&customer)

	if err := repository.GetCustomerRepository().GetCustomerEmail(customer.Email); err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "Email Account Already Exists",
		})
	}

	valid := md.PostCustValidation(customer)
	if valid != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": valid.Error(),
		})
	}

	if err := repository.GetCustomerRepository().CreateCustomer(&customer); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Status":  "500",
			"Message": err.Error(),
		})
	}
	result := m.CustomerResponse{
		Id:      customer.Id,
		Name:    customer.Name,
		Address: customer.Address,
		Phone:   customer.Phone,
		Email:   customer.Email,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":   "200",
		"Message":  "succes create new customer",
		"Customer": result,
	})
}

func DeleteCustomer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if err := repository.GetCustomerRepository().DeleteCustomer(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success delete customer",
	})
}

func UpdateCustomer(c echo.Context) error {
	updateData := m.Customer{}
	c.Bind(&updateData)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": "invalid id",
		})
	}

	if updateData.Password != "" {
		hash, err := u.HashPassword(updateData.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  "400",
				"Message": err.Error(),
			})
		}

		updateData.Password = hash
	}
	if updateData.Email != "" {
		valid := md.EmailValidation(updateData.Email)
		if valid != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Status":  "400",
				"Message": valid.Error(),
			})
		}
	}

	if err := repository.GetCustomerRepository().UpdateCustomer(&updateData, id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Message": "success update customer",
	})
}

func LoginCustomer(c echo.Context) error {
	customer := m.Customer{}
	c.Bind(&customer)

	if err := repository.GetCustomerRepository().LoginCustomer(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	token, err := md.CreateToken(int(customer.Id), customer.Name, customer.Role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Status":  "400",
			"Message": err.Error(),
		})
	}

	customerResponse := m.AdminResponseLogin{Id: int(customer.Id), Name: customer.Name, Email: customer.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Status":   "200",
		"Message":  "success login",
		"Customer": customerResponse,
	})
}
