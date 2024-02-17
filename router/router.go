package router

import (
	"go-crud/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(tagController *controller.TagController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home!")
	})

	baseRouter := router.Group("/api")
	tagRouter := baseRouter.Group("/tags")

	tagRouter.GET("", tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PUT("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	return router

}
