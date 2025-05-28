package app

import (
	"time"

	helperLogger "gitlab.com/VNEU/logger/helper"
	goHelper "gitlab.com/vneu/go-helper/helper"
	"gitlab.com/voltunes/api-master-project/helper"
	"gitlab.com/voltunes/api-master-project/route"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {
	serviceName := "GO-VISIT-FLOW"
	goHelper.InitTracer(serviceName)
	router := gin.Default()
	router.Use(otelgin.Middleware(serviceName))
	router.Use(helperLogger.LoggerHandler())
	router.Use(helperLogger.ErrorHandler(helper.DatabaseErrors))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	route.ProductRoute(router, db, validate)

	return router
}
