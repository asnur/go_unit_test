package repository

import "golang_testing/entity"

type CategoryRepository interface {
	FindByid(id string) *entity.Category
}
