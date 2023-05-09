package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type StoreRepositoryMock struct {
	Mock mock.Mock
}

var dataStore = []models.Store{
	{},
}

func (um *StoreRepositoryMock) CreateStore(store *models.Store) error {
	argumnents := um.Mock.Called(store)
	return argumnents.Error(0)
}

func (um *StoreRepositoryMock) GetStore(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataStore {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataStore[index]
}
func (um *StoreRepositoryMock) GetStoreEmail(email string) error {
	argumnents := um.Mock.Called(email)

	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	var index int
	for i, val := range dataStore {
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

func (um *StoreRepositoryMock) GetStores() (error, interface{}) {
	if dataStore == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataStore
}

func (um *StoreRepositoryMock) DeleteStore(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataStore {
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

func (um *StoreRepositoryMock) UpdateStore(dataUpdate *models.Store, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataStore {
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
