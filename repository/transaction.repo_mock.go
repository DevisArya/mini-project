package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	Mock mock.Mock
}

var dataTransaction = []models.Transaction{
	{},
}

func (um *TransactionRepositoryMock) CreateTransaction(transaction *models.Transaction) error {
	argumnents := um.Mock.Called(transaction)
	return argumnents.Error(0)
}

func (um *TransactionRepositoryMock) GetTransaction(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataTransaction {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataTransaction[index]
}

func (um *TransactionRepositoryMock) GetTransactions() (error, interface{}) {
	if dataTransaction == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataTransaction
}

func (um *TransactionRepositoryMock) DeleteTransaction(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataTransaction {
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

func (um *TransactionRepositoryMock) UpdateTransaction(dataUpdate *models.Transaction, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataTransaction {
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
