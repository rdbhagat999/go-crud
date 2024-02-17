package main

import (
	"go-crud/config"
	"go-crud/controller"
	"go-crud/helper"
	"go-crud/model"
	"go-crud/repository"
	"go-crud/router"
	"go-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Server started")

	db := config.DatabaseConnection()

	db.Table("tags").AutoMigrate(&model.Tag{})

	tagRepo := repository.NewTagRepositoryImpl(db)
	validate := validator.New()
	tagS := service.NewTagServiceImpl(tagRepo, validate)
	tagController := controller.NewTagController(tagS)

	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()

	helper.ErrorPanic(err)
}
