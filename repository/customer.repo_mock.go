package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	Mock mock.Mock
}

var dataCustomer = []models.Customer{
	{
		ID:       1,
		Name:     "Devis",
		Address:  "Jombang",
		Phone:    "12345678910",
		Email:    "devis@gmail.com",
		Role:     false,
		Password: "123",
	},
	{
		ID:       2,
		Name:     "Devis2",
		Address:  "Jombang",
		Phone:    "12345678910",
		Email:    "devis@gmail.com",
		Role:     false,
		Password: "123",
	},
}

func (um *CustomerRepositoryMock) CreateCustomer(customer *models.Customer) error {
	argumnents := um.Mock.Called(customer)
	return argumnents.Error(0)
}

func (um *CustomerRepositoryMock) GetCustomer(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error id"), nil
	}
	var index int
	for i, val := range dataCustomer {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataCustomer[index]
}

func (um *CustomerRepositoryMock) GetCustomerEmail(email string) error {
	argumnents := um.Mock.Called(email)

	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	var index int
	for i, val := range dataCustomer {
		if val.Email == email {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("email not found")
	}

	return nil
}

func (um *CustomerRepositoryMock) GetCustomers() (error, interface{}) {
	if dataCustomer == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataCustomer
}

func (um *CustomerRepositoryMock) DeleteCustomer(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataCustomer {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found")
	}

	return nil
}

func (um *CustomerRepositoryMock) UpdateCustomer(dataUpdate *models.Customer, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataCustomer {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found")
	}

	return nil
}

func (um *CustomerRepositoryMock) LoginCustomer(customer *models.Customer) error {
	argumnents := um.Mock.Called(customer)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataCustomer {
		if val.Email == customer.Email {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("user not found")
	}

	if customer.Password != dataCustomer[index].Password {
		return errors.New("mismatch password")
	}

	return nil
}
