package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type TransactionDetailRepositoryMock struct {
	Mock mock.Mock
}

var dataTransactionDetail = []models.TransactionDetail{
	{},
}

func (um *TransactionDetailRepositoryMock) CreateTransactionDetail(transactiondetaildetail *models.TransactionDetail) error {
	argumnents := um.Mock.Called(transactiondetaildetail)
	return argumnents.Error(0)
}

func (um *TransactionDetailRepositoryMock) GetTransactionDetail(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataTransactionDetail {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataTransactionDetail[index]
}

func (um *TransactionDetailRepositoryMock) GetTransactionDetails() (error, interface{}) {
	if dataTransactionDetail == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataTransactionDetail
}

func (um *TransactionDetailRepositoryMock) DeleteTransactionDetail(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataTransactionDetail {
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

func (um *TransactionDetailRepositoryMock) UpdateTransactionDetail(dataUpdate *models.TransactionDetail, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataTransactionDetail {
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
