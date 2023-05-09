package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type ServiceTypeRepositoryMock struct {
	Mock mock.Mock
}

var dataServiceType = []models.ServiceType{
	{},
}

func (um *ServiceTypeRepositoryMock) CreateServiceType(servicetype *models.ServiceType) error {
	argumnents := um.Mock.Called(servicetype)
	return argumnents.Error(0)
}

func (um *ServiceTypeRepositoryMock) GetServiceType(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataServiceType {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataServiceType[index]
}

func (um *ServiceTypeRepositoryMock) GetServiceTypes() (error, interface{}) {
	if dataServiceType == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataServiceType
}

func (um *ServiceTypeRepositoryMock) DeleteServiceType(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataServiceType {
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

func (um *ServiceTypeRepositoryMock) UpdateServiceType(dataUpdate *models.ServiceType, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataServiceType {
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
