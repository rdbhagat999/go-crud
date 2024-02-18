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
	"time"

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

	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		evt := log.Fatal()
		evt.Msg("🚀 Could not load environment variables")
	}

	db := config.DatabaseConnection(&loadConfig)

	db.Table("tags").AutoMigrate(&model.Tag{})
	db.Table("users").AutoMigrate(&model.User{})

	validate := validator.New()

	//Init Repository
	tagRepository := repository.NewTagRepositoryImpl(db)
	userRepository := repository.NewUserRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagServiceImpl(tagRepository, validate)
	userService := service.NewUserServiceImpl(userRepository, validate)

	//Init Controller
	tagController := controller.NewTagController(tagService)
	userController := controller.NewUserController(userService)

	router, apiVersion1 := router.NewRouter()

	tagRouter := apiVersion1.Group("/tags")
	tagRouter.GET("/", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("/", tagController.Create)
	tagRouter.PUT("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	userRouter := apiVersion1.Group("/users")
	userRouter.GET("/", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("/", userController.Create)
	userRouter.PUT("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	server := &http.Server{
		Addr:           ":" + loadConfig.SERVER_PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()

	helper.ErrorPanic(server_err)
}
