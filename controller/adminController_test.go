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

func TestCreateAdminValid(t *testing.T) {
	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	var dataAdmin = models.Admin{
		Name:     "Devis",
		Address:  "Jombang",
		Phone:    "12345678910",
		Email:    "devis@gmail.com",
		Password: "123",
	}

	adminRepository.Mock.On("GetAdminEmail", dataAdmin.Email).Return(nil)
	adminRepository.Mock.On("CreateAdmin", &dataAdmin).Return(dataAdmin)

	e := echo.New()

	bData, _ := json.Marshal(dataAdmin)
	req := httptest.NewRequest(http.MethodPost, "/admin/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateAdmin(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestCreateAdminInvalid(t *testing.T) {
	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	var dataAdmin = models.Admin{
		Name: "Devis",
	}

	adminRepository.Mock.On("GetAdminEmail", dataAdmin.Email).Return(nil)
	adminRepository.Mock.On("CreateAdmin", &dataAdmin).Return(nil)

	e := echo.New()

	bData, _ := json.Marshal(dataAdmin)
	req := httptest.NewRequest(http.MethodPost, "/admin/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	CreateAdmin(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetAdminValid(t *testing.T) {

	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	adminRepository.Mock.On("GetAdmin", 1).Return(1)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/admins/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	GetAdmin(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestGetAdminInvalid(t *testing.T) {

	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	adminRepository.Mock.On("GetAdmin", 3).Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/admins/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("3")

	GetAdmin(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetAdmins(t *testing.T) {
	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	adminRepository.Mock.On("GetAdmins").Return(nil)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/admins/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetAdmins(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteAdminValid(t *testing.T) {
	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	adminRepository.Mock.On("DeleteAdmin", 1).Return(1)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/admins/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	DeleteAdmin(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestDeleteAdminInvalid(t *testing.T) {
	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	adminRepository.Mock.On("DeleteAdmin", 10).Return(10)

	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/admins/", nil)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("10")

	DeleteAdmin(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateAdminValid(t *testing.T) {
	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	data := models.Admin{
		Name: "Test Update",
	}

	adminRepository.Mock.On("UpdateAdmin", &data, 1).Return(1)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/admins/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("1")

	UpdateAdmin(c)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestUpdateAdminInvalid(t *testing.T) {
	adminRepository := &r.AdminRepositoryMock{Mock: mock.Mock{}}
	r.SetAdminRepository(adminRepository)

	data := models.Admin{
		Name: "TEST",
	}

	adminRepository.Mock.On("UpdateAdmin", &data, 10).Return(10)

	e := echo.New()
	bData, _ := json.Marshal(data)
	req := httptest.NewRequest(http.MethodPut, "/admins/", bytes.NewReader(bData))
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id/")
	c.SetParamNames("id")
	c.SetParamValues("10")

	UpdateAdmin(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
