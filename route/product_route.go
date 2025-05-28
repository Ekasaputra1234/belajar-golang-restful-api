package route

import (
	"gitlab.com/voltunes/api-master-project/auth"
	"gitlab.com/voltunes/api-master-project/controller"
	"gitlab.com/voltunes/api-master-project/repository"
	"gitlab.com/voltunes/api-master-project/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ProductRoute(router *gin.Engine, DB *gorm.DB, validate *validator.Validate) {

	productService := service.NewProductService(
		repository.NewProductRepository(),
		DB,
		validate,
	)
	productController := controller.NewProductController(productService)

	router.DELETE("/products/:id", auth.Auth(productController.Delete, []string{}))
	router.GET("/products", auth.Auth(productController.FindAll, []string{}))
	router.GET("/products/:id", auth.Auth(productController.FindByID, []string{}))
	router.POST("/products", auth.Auth(productController.Create, []string{}))
	router.PUT("/products/:id", auth.Auth(productController.Update, []string{}))
}
