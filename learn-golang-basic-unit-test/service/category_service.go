package service

import (
	"errors"
	"github.com/bayujuara01/learn-golang-basic-unit-test/entity"
	"github.com/bayujuara01/learn-golang-basic-unit-test/repository"
)

type CategoryService struct {
	Repository repository.Category
}

func (service CategoryService) Get(id int) (*entity.CategoryEntity, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category Not Found")
	} else {
		return category, nil
	}
}

func (service CategoryService) GetByName(name string) (*[]entity.CategoryEntity, error) {
	category := service.Repository.FindByName(name)

	if category == nil {
		return nil, errors.New("category Not Found")
	} else {
		return category, nil
	}
}
