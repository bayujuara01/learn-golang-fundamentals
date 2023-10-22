package repository

import (
	"github.com/bayujuara01/learn-golang-basic-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id int) *entity.CategoryEntity {

	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.CategoryEntity)
		return &category
	}

}

func (repository *CategoryRepositoryMock) FindByName(name string) *[]entity.CategoryEntity {
	//TODO implement me
	panic("implement me")
}
