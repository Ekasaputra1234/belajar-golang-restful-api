package controller

import (
	"gitlab.com/voltunes/api-master-project/auth"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Create(context *gin.Context, auth *auth.AccessDetails)
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	FindByID(context *gin.Context, auth *auth.AccessDetails)
	Delete(context *gin.Context, auth *auth.AccessDetails)
	Update(context *gin.Context, auth *auth.AccessDetails)
}
