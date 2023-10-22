package service

import (
	"github.com/bayujuara01/learn-golang-basic-unit-test/entity"
	"github.com/bayujuara01/learn-golang-basic-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get_NotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", 1).Return(nil)
	category, err := categoryService.Get(1)
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_Get(t *testing.T) {
	resultCategory := entity.CategoryEntity{
		Id:   2,
		Name: "Berbahaya",
	}
	categoryRepository.Mock.On("FindById", 2).Return(resultCategory)
	result, err := categoryService.Get(2)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, int(result.Id))
	assert.Equal(t, "Berbahaya", result.Name)
}
