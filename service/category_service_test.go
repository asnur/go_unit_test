package service

import (
	"golang_testing/entity"
	"golang_testing/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	//Program mock
	categoryRepository.Mock.On("FindByid", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_Getfound(t *testing.T) {
	//Program mock
	category := &entity.Category{ID: "2", Name: "asnur"}
	categoryRepository.Mock.On("FindByid", "2").Return(category)

	category, err := categoryService.Get("2")
	assert.NotNil(t, category)
	assert.Nil(t, err)
}
