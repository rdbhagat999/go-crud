package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	// router.Use(gin.Logger())

	// adds Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	baseRouter := router.Group("/api")
	apiVersion1 := baseRouter.Group("/v1")

	return router, apiVersion1

}
