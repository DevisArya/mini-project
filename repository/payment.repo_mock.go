package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	Mock mock.Mock
}

var dataPayment = []models.Payment{
	{},
}

func (um *PaymentRepositoryMock) CreatePayment(payment *models.Payment) error {
	argumnents := um.Mock.Called(payment)
	return argumnents.Error(0)
}

func (um *PaymentRepositoryMock) GetPayment(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataPayment {
		if val.Id == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataPayment[index]
}

func (um *PaymentRepositoryMock) GetPayments() (error, interface{}) {
	if dataPayment == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataPayment
}

func (um *PaymentRepositoryMock) DeletePayment(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataPayment {
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

func (um *PaymentRepositoryMock) UpdatePayment(dataUpdate *models.Payment, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataPayment {
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
