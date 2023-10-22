package repository

import "github.com/bayujuara01/learn-golang-basic-unit-test/entity"

type Category interface {
	FindById(id int) *entity.CategoryEntity
	FindByName(name string) *[]entity.CategoryEntity
}
