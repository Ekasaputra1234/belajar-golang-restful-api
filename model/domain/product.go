package domain

import (
	"gitlab.com/voltunes/api-master-project/model/web"

	"gorm.io/gorm"
)

type Products []Product
type Product struct {
	// Required Fields
	gorm.Model
	CreatedByID string  `gorm:""`
	UpdatedByID string  `gorm:""`
	DeletedByID *string `gorm:""`

	// Fields
	Name     string  `gorm:""`
	Price    float64 `gorm:""`
	Stock    int     `gorm:""`
	Category string  `gorm:""`

	//Linked Foreign Key
}

func (product *Product) ToProductResponse() web.ProductResponse {
	return web.ProductResponse{
		// Required Fields
		ID: product.ID,

		// Fields
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Category: product.Category,
	}
}

func (products Products) ToProductResponses() []web.ProductResponse {
	productResponses := []web.ProductResponse{}
	for _, product := range products {
		productResponses = append(productResponses, product.ToProductResponse())
	}
	return productResponses
}
