package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type CleanerRepositoryMock struct {
	Mock mock.Mock
}

var dataCleaner = []models.Cleaner{
	{},
}

func (um *CleanerRepositoryMock) CreateCleaner(cleaner *models.Cleaner) error {
	argumnents := um.Mock.Called(cleaner)
	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	return nil
}

func (um *CleanerRepositoryMock) GetCleaner(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataCleaner {
		if val.Id == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataCleaner[index]
}
func (um *CleanerRepositoryMock) GetCleanerEmail(email string) error {
	argumnents := um.Mock.Called(email)

	if argumnents.Get(0) == nil {
		return errors.New("error")
	}
	var index int
	for i, val := range dataCleaner {
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

func (um *CleanerRepositoryMock) GetCleaners() (error, interface{}) {
	if dataCleaner == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataCleaner
}

func (um *CleanerRepositoryMock) DeleteCleaner(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataCleaner {
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

func (um *CleanerRepositoryMock) UpdateCleaner(dataUpdate *models.Cleaner, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataCleaner {
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
