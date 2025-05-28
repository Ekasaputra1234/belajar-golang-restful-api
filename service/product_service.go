package service

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/voltunes/api-master-project/auth"
	"gitlab.com/voltunes/api-master-project/model/web"
)

type ProductService interface {
	Create(auth *auth.AccessDetails, request *web.ProductCreateRequest, c *gin.Context) web.ProductResponse
	Delete(auth *auth.AccessDetails, id *int, c *gin.Context)
	Update(auth *auth.AccessDetails, id *int, request *web.ProductUpdateRequest, c *gin.Context) web.ProductResponse
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse
	FindByID(auth *auth.AccessDetails, id *int, c *gin.Context) web.ProductResponse
}
