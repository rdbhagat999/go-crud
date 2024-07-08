package main

import (
	"fmt"
	_ "go-crud/docs"
	"go-crud/src/config"
	"go-crud/src/controller"
	"go-crud/src/dsa"
	"go-crud/src/helper"
	"go-crud/src/model"
	"go-crud/src/repository"
	"go-crud/src/router"
	"go-crud/src/service"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func greetFunction() {
	fmt.Println("Inside greetFunction")
	fmt.Println("Hello world!")
}

func calFactorialFunction() {
	fmt.Println("Inside calFactorialFunction")
	factorial := dsa.FactorialRecursive(5)
	println("factorial: ", factorial)
}

func calFibonacciFunction() {
	fmt.Println("Inside calFibonacciFunction")
	fibonacci := dsa.FibonacciRecursive(5)
	println("fibonacci: ", fibonacci)
}

func executeGreetFunction() {
	fmt.Println("Inside executeGreetFunction")

	defer greetFunction()
	defer calFactorialFunction()
	defer calFibonacciFunction()

	triggerPanic()
}

func triggerPanic() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover from panic with message:", err)
		}
	}()

	panic("This is a Panic message")
}

// @title Tag Service API
// @version 1.0.0
// @description A Tag Service API in Golang using Gin Framework

// @host localhost:8888
// @BasePath /api/v1
func main() {

	executeGreetFunction()

	log.Info().Msg("Server started")

	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		evt := log.Fatal()
		evt.Msg("🚀 Could not load environment variables")
	}

	db := config.DatabaseConnection(&loadConfig)

	db.Table("tags").AutoMigrate(&model.Tag{})
	db.Table("posts").AutoMigrate(&model.Post{})
	db.Table("users").AutoMigrate(&model.User{})

	validate := validator.New()

	//Init Repository
	tagRepository := repository.NewTagRepositoryImpl(db)
	postRepository := repository.NewPostRepositoryImpl(db)
	userRepository := repository.NewUserRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagServiceImpl(tagRepository, validate)
	postService := service.NewPostServiceImpl(postRepository, validate)
	userService := service.NewUserServiceImpl(userRepository, validate)

	//Init Controller
	tagController := controller.NewTagController(tagService)
	postController := controller.NewPostController(postService)
	userController := controller.NewUserController(userService)

	router, apiVersion1 := router.NewRouter()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 24 * time.Hour,
	}))

	tagRouter := apiVersion1.Group("/tags")
	tagRouter.GET("/", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("/", tagController.Create)
	tagRouter.PUT("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	postRouter := apiVersion1.Group("/posts")
	postRouter.GET("/", postController.FindAll)
	postRouter.GET("/:postId", postController.FindById)
	postRouter.POST("/", postController.Create)
	postRouter.PUT("/:postId", postController.Update)
	postRouter.DELETE("/:postId", postController.Delete)

	userRouter := apiVersion1.Group("/users")
	userRouter.GET("/", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("/", userController.Create)
	userRouter.POST("/login", userController.Login)
	userRouter.POST("/logout", userController.Logout)
	userRouter.POST("/authuser", userController.AuthUser)
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
