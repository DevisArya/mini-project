package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type AreaRepositoryMock struct {
	Mock mock.Mock
}

var dataArea = []models.Area{
	{},
}

func (um *AreaRepositoryMock) CreateArea(area *models.Area) error {
	argumnents := um.Mock.Called(area)
	return argumnents.Error(0)
}

func (um *AreaRepositoryMock) GetArea(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error id"), nil
	}
	var index int
	for i, val := range dataArea {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataArea[index]
}
func (um *AreaRepositoryMock) GetAreaName(name string) error {
	argumnents := um.Mock.Called(name)

	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	var index int
	for i, val := range dataArea {
		if val.Name == name {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found")
	}

	return nil
}

func (um *AreaRepositoryMock) GetAreas() (error, interface{}) {
	if dataArea == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataArea
}

func (um *AreaRepositoryMock) DeleteArea(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataArea {
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

func (um *AreaRepositoryMock) UpdateArea(dataUpdate *models.Area, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataArea {
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
