package service

import (
	"errors"
	"golang_testing/entity"
	"golang_testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindByid(id)
	if category == nil {
		return category, errors.New("Category not found")
	}
	return category, nil
}
