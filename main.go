package main

import (
	_ "go-crud/docs"
	"go-crud/src/config"
	"go-crud/src/controller"
	"go-crud/src/helper"
	"go-crud/src/model"
	"go-crud/src/repository"
	"go-crud/src/router"
	"go-crud/src/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title Tag Service API
// @version 1.0.0
// @description A Tag Service API in Golang using Gin Framework

// @host localhost:8888
// @BasePath /api/v1
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
