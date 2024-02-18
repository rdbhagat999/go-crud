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

	validate := validator.New()

	tagRepository := repository.NewTagRepositoryImpl(db)
	tagService := service.NewTagServiceImpl(tagRepository, validate)
	tagController := controller.NewTagController(tagService)

	router, apiVersion1 := router.NewRouter()

	tagRouter := apiVersion1.Group("/tags")
	tagRouter.GET("/", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("/", tagController.Create)
	tagRouter.PUT("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	err := server.ListenAndServe()

	helper.ErrorPanic(err)
}
