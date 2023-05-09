package repository

import (
	"errors"
	"miniproject/models"

	"github.com/stretchr/testify/mock"
)

type TeamRepositoryMock struct {
	Mock mock.Mock
}

var dataTeam = []models.Team{
	{},
}

func (um *TeamRepositoryMock) CreateTeam(team *models.Team) error {
	argumnents := um.Mock.Called(team)
	return argumnents.Error(0)
}

func (um *TeamRepositoryMock) GetTeam(id int) (error, interface{}) {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("error"), nil
	}
	var index int
	for i, val := range dataTeam {
		if val.ID == uint(id) {
			index = i + 1
			break
		}
	}

	if index == 0 {
		return errors.New("id not found"), nil
	}

	return nil, dataTeam[index]
}

func (um *TeamRepositoryMock) GetTeams() (error, interface{}) {
	if dataTeam == nil {
		return errors.New("internal server error"), nil
	}
	return nil, dataTeam
}

func (um *TeamRepositoryMock) DeleteTeam(id int) error {
	argumnents := um.Mock.Called(id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataTeam {
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

func (um *TeamRepositoryMock) UpdateTeam(dataUpdate *models.Team, id int) error {
	argumnents := um.Mock.Called(dataUpdate, id)

	if argumnents.Get(0) == nil {
		return errors.New("bad request")
	}
	var index int
	for i, val := range dataTeam {
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
