package main

import (
	"fmt"
	_ "go-crud/docs"
	"go-crud/src/config"
	"go-crud/src/constants"
	"go-crud/src/controller"
	"go-crud/src/dsa"
	"go-crud/src/helper"
	"go-crud/src/middlewares"
	"go-crud/src/model"
	"go-crud/src/repository"
	"go-crud/src/router"
	"go-crud/src/service"
	"net/http"
	"time"

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

// @title Demo Service API
// @version 1.0.0
// @description A demo Service API in Golang using Gin Framework

// @host localhost:8888
// @BasePath /api/v1
func main() {

	executeGreetFunction()

	log.Info().Msg("Server started")

	loadConfig, err := config.LoadConfig()

	if err != nil {
		evt := log.Fatal()
		evt.Msg("🚀 Could not load environment variables")
	}

	db := config.DatabaseConnection(&loadConfig)

	db.Table(constants.PostTable).AutoMigrate(&model.Post{})
	db.Table(constants.UserTable).AutoMigrate(&model.User{})

	validate := validator.New()

	//Init Repository
	postRepository := repository.NewPostRepositoryImpl(db)
	userRepository := repository.NewUserRepositoryImpl(db)

	//Init Service
	postService := service.NewPostServiceImpl(postRepository, validate)
	userService := service.NewUserServiceImpl(userRepository, validate)

	//Init Controller
	postController := controller.NewPostController(postService)
	userController := controller.NewUserController(userService)

	router, apiVersion1 := router.NewRouter()

	authRouter := apiVersion1.Group(constants.AuthGroup)
	authRouter.POST(constants.RegisterRoute, userController.Create)
	authRouter.POST(constants.LoginRoute, userController.Login)
	authRouter.POST(constants.LogoutRoute, userController.Logout)

	uploadRouter := apiVersion1.Group(constants.UploadGroup)
	uploadRouter.POST(constants.UploadFileRoute, controller.UploadFile)

	userRouter := apiVersion1.Group(constants.UserGroup)
	userRouter.Use(middlewares.JWTAuthMiddleware(userController))

	userRouter.GET(constants.GetAllUsersRoute, userController.FindAll)
	userRouter.GET(constants.GetUserByIdRoute, userController.FindById)
	userRouter.GET(constants.GetAuthUserRoute, userController.AuthUser)
	userRouter.PUT(constants.UpdateUserRoute, userController.Update)
	userRouter.DELETE(constants.DeleteUserRoute, userController.Delete)

	postRouter := apiVersion1.Group(constants.PostGroup)
	postRouter.Use(middlewares.JWTAuthMiddleware(userController))

	postRouter.GET(constants.GetAllPostsRoute, postController.FindAll)
	postRouter.GET(constants.GetPostsByUserRoute, postController.FindAllByUserId)
	postRouter.GET(constants.GetPostByIdRoute, postController.FindById)
	postRouter.POST(constants.CreatePostRoute, postController.Create)
	postRouter.PUT(constants.UpdatePostRoute, postController.Update)
	postRouter.DELETE(constants.DeletePostRoute, postController.Delete)

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
