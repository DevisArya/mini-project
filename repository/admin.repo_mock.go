package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type AdminRepositoryMock struct {
	Mock mock.Mock
}

var dataAdmin = []models.Admin{
	{
		Id:       1,
		Name:     "Devis",
		Address:  "Jombang",
		Phone:    "12345678910",
		Email:    "devis@gmail.com",
		Role:     true,
		Password: "123",
	},
	{
		Id:       2,
		Name:     "Devis2",
		Address:  "Jombang",
		Phone:    "12345678910",
		Email:    "devis@gmail.com",
		Role:     true,
		Password: "123",
	},
}

func (um *AdminRepositoryMock) CreateAdmin(admin *models.Admin) error {
	argumnents := um.Mock.Called(admin)
	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	return nil
}

func (um *AdminRepositoryMock) GetAdmin(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataAdmin {
		if int(val.Id) == id {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataAdmin[index]
}
func (um *AdminRepositoryMock) GetAdminEmail(email string) error {
	argumnents := um.Mock.Called(email)

	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	var index int
	for i, val := range dataAdmin {
		if val.Email == email {
			index = i
			break
		}
	}

	if index == 0 {
		return errors.New("email not found")
	}

	return nil
}

func (um *AdminRepositoryMock) GetAdmins() (error, interface{}) {
	if dataAdmin == nil {
		return errors.New("internal server error"), nil

	}
	return nil, dataAdmin
}

func (um *AdminRepositoryMock) DeleteAdmin(id int) error {

	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	var index int
	for i, val := range dataAdmin {
		if val.Id == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found")
	}

	return nil
}

func (um *AdminRepositoryMock) UpdateAdmin(dataUpdate *models.Admin, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataAdmin {
		if val.Id == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found")
	}

	return nil
}

func (um *AdminRepositoryMock) LoginAdmin(admin *models.Admin) error {
	argumnents := um.Mock.Called(admin)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataAdmin {
		if val.Email == admin.Email {
			index = i
			break
		}
	}

	if index == 0 {
		return errors.New("user not found")
	}

	if admin.Password != dataAdmin[index].Password {
		return errors.New("mismatch password")
	}

	return nil
}
