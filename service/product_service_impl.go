package service

import (
	goHelper "gitlab.com/vneu/go-helper/helper"
	"gitlab.com/voltunes/api-master-project/auth"
	"gitlab.com/voltunes/api-master-project/helper"
	"gitlab.com/voltunes/api-master-project/model/domain"
	"gitlab.com/voltunes/api-master-project/model/web"
	"gitlab.com/voltunes/api-master-project/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewProductService(
	city repository.ProductRepository,
	db *gorm.DB,
	validate *validator.Validate,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository: city,
		DB:                db,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse {
	goHelper.SignozSpan("product", c)

	tx := goHelper.CreateTransaction(service.DB, c)
	defer goHelper.CommitOrRollback(tx.Write)

	products := service.ProductRepository.FindAll(tx, filters)
	return products.ToProductResponses()
}

func (service *ProductServiceImpl) Create(auth *auth.AccessDetails, request *web.ProductCreateRequest, c *gin.Context) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	goHelper.SignozSpan("product", c)

	tx := goHelper.CreateTransaction(service.DB, c)
	defer goHelper.CommitOrRollback(tx.Write)

	product := &domain.Product{
		// Required Fields
		CreatedByID: auth.UserID,
		UpdatedByID: auth.UserID,

		// Fields
		Name:     request.Name,
		Price:    request.Price,
		Stock:    request.Stock,
		Category: request.Category,
	}
	product = service.ProductRepository.Create(tx, product)

	return product.ToProductResponse()
}

func (service *ProductServiceImpl) Delete(auth *auth.AccessDetails, id *int, c *gin.Context) {
	goHelper.SignozSpan("product", c)

	tx := goHelper.CreateTransaction(service.DB, c)
	defer goHelper.CommitOrRollback(tx.Write)
	service.ProductRepository.Delete(tx, id, &auth.UserID)
}

func (service *ProductServiceImpl) Update(auth *auth.AccessDetails, id *int, request *web.ProductUpdateRequest, c *gin.Context) web.ProductResponse {
	goHelper.SignozSpan("product", c)

	tx := goHelper.CreateTransaction(service.DB, c)
	defer goHelper.CommitOrRollback(tx.Write)

	subject := &domain.Product{
		// Required Fields
		Model:       gorm.Model{ID: uint(*id)},
		UpdatedByID: auth.UserID,

		//  Fields
		Name:     request.Name,
		Price:    request.Price,
		Stock:    request.Stock,
		Category: request.Category,
	}
	subject = service.ProductRepository.Update(tx, subject)
	return subject.ToProductResponse()
}

func (service *ProductServiceImpl) FindByID(auth *auth.AccessDetails, id *int, c *gin.Context) web.ProductResponse {
	goHelper.SignozSpan("product", c)

	tx := goHelper.CreateTransaction(service.DB, c)
	defer goHelper.CommitOrRollback(tx.Write)

	product := service.ProductRepository.FindByID(tx, id)
	return product.ToProductResponse()
}
