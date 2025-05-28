package repository

import (
	goHelper "gitlab.com/vneu/go-helper/helper"
	"gitlab.com/voltunes/api-master-project/helper"
	"gitlab.com/voltunes/api-master-project/model/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) FindAll(db *goHelper.DatabaseResolver, filters *map[string]string) domain.Products {
	products := domain.Products{}
	tx := db.Read.Table("products")

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	err = tx.Find(&products).Error
	helper.PanicIfError(err)

	return products
}

func (repository *ProductRepositoryImpl) Create(db *goHelper.DatabaseResolver, product *domain.Product) *domain.Product {
	err := db.Write.Create(&product).Error
	helper.PanicIfError(err)
	return product
}

func (repository *ProductRepositoryImpl) Delete(db *goHelper.DatabaseResolver, id *int, deletedByID *string) {
	err := db.Read.First(&domain.Product{}, id).Error
	helper.PanicIfError(err)

	err = db.Write.Updates(
		&domain.Product{
			Model:       gorm.Model{ID: uint(*id)},
			DeletedByID: deletedByID,
		},
	).Delete(&domain.Product{}, id).Error
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) Update(db *goHelper.DatabaseResolver, product *domain.Product) *domain.Product {
	err := db.Write.Table("products").Where("id = ?", product.ID).Updates(&product).First(&product).Error
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) FindByID(db *goHelper.DatabaseResolver, id *int) domain.Product {
	var product domain.Product
	err := db.Read.Table("products").Where("products.id = ?", id).First(&product).Error
	helper.PanicIfError(err)
	return product
}
