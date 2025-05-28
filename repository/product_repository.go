package repository

import (
	goHelper "gitlab.com/vneu/go-helper/helper"
	"gitlab.com/voltunes/api-master-project/model/domain"
)

type ProductRepository interface {
	Create(db *goHelper.DatabaseResolver, product *domain.Product) *domain.Product
	Delete(db *goHelper.DatabaseResolver, id *int, deletedByID *string)
	FindAll(db *goHelper.DatabaseResolver, filters *map[string]string) domain.Products
	FindByID(db *goHelper.DatabaseResolver, id *int) domain.Product
	Update(db *goHelper.DatabaseResolver, product *domain.Product) *domain.Product
}
