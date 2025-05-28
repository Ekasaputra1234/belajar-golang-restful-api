package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	goHelper "gitlab.com/vneu/go-helper/helper"
	"gitlab.com/voltunes/api-master-project/app"
	"gitlab.com/voltunes/api-master-project/helper"
)

func main() {
	configuration, err := goHelper.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	port := configuration.Port
	db := app.ConnectDatabase(configuration)

	// Validator
	validate := validator.New()
	helper.RegisterValidation(validate)

	router := app.NewRouter(db, validate)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server is running on port %s", port)

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
