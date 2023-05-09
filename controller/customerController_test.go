package controller

import (
	"bytes"
	"encoding/json"
	"miniproject/models"
	r "miniproject/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomerValid(t *testing.T) {
	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	var dataCustomer = models.Customer{
		Name:     "Devis",
		Address:  "Jombang",
		Phone:    "12345678910",
		Email:    "devis@gmail.com",
		Password: "123",
	}

	customerRepository.Mock.On("GetCustomerEmail", dataCustomer.Email).Return(nil)
	customerRepository.Mock.On("CreateCustomer", &dataCustomer).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(dataCustomer)
	req := httptest.NewRequest(http.MethodPost, "/customer/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateCustomer(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestCreateCustomerInvalid(t *testing.T) {
	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	var dataCustomer = models.Customer{
		Name: "Devis",
	}

	customerRepository.Mock.On("GetCustomerEmail", dataCustomer.Email).Return(nil)
	customerRepository.Mock.On("CreateCustomer", &dataCustomer).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(dataCustomer)
	req := httptest.NewRequest(http.MethodPost, "/customer/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateCustomer(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetCustomerValid(t *testing.T) {

	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	customerRepository.Mock.On("GetCustomer", 1).Return(1)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/customers/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetCustomer(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestGetCustomerInvalid(t *testing.T) {

	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	customerRepository.Mock.On("GetCustomer", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/customers/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	GetCustomer(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetCustomers(t *testing.T) {
	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	customerRepository.Mock.On("GetCustomers").Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/customers/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetCustomers(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteCustomerValid(t *testing.T) {
	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	customerRepository.Mock.On("DeleteCustomer", 1).Return(1)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/customers/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteCustomer(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestDeleteCustomerInvalid(t *testing.T) {
	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	customerRepository.Mock.On("DeleteCustomer", 10).Return(10)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/customers/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("10")

	DeleteCustomer(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateCustomerValid(t *testing.T) {
	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	data := models.Customer{
		Name: "Test Update",
	}

	customerRepository.Mock.On("UpdateCustomer", &data, 1).Return(1)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/customers/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateCustomer(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestUpdateCustomerInvalid(t *testing.T) {
	customerRepository := &r.CustomerRepositoryMock{Mock: mock.Mock{}}
	r.SetCustomerRepository(customerRepository)

	data := models.Customer{
		Name: "TEST",
	}

	customerRepository.Mock.On("UpdateCustomer", &data, 10).Return(10)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/customers/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("10")

	UpdateCustomer(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
