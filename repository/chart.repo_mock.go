package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type ChartRepositoryMock struct {
	Mock mock.Mock
}

var dataChart = []models.Chart{
	{},
}

func (um *ChartRepositoryMock) CreateChart(chartdetail *models.Chart) error {
	argumnents := um.Mock.Called(chartdetail)
	return argumnents.Error(0)
}

func (um *ChartRepositoryMock) GetChart(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataChart {
		if val.Id == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataChart[index]
}

func (um *ChartRepositoryMock) GetCharts() (error, interface{}) {
	if dataChart == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataChart
}

func (um *ChartRepositoryMock) DeleteChart(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataChart {
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

func (um *ChartRepositoryMock) UpdateChart(dataUpdate *models.Chart, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataChart {
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
