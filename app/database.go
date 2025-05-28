package app

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	goHelper "gitlab.com/vneu/go-helper/helper"

	"gorm.io/gorm"
)

func ConnectDatabase(configuration goHelper.Configuration) *gorm.DB {
	database := goHelper.ConnectMysqlDatabaseResolver(configuration)

	if err := database.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	err := database.AutoMigrate(
	// &domain.Product{},
	// &domain.ProductProduct{},
	// &domain.ProductTask{},
	// &domain.ProductTaskDetail{},
	)
	if err != nil {
		panic("failed to auto migrate schema")
	}

	return database
}
