package controller

import (
	"go-crud/data/request"
	"go-crud/data/response"
	"go-crud/helper"
	"go-crud/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	TagService service.TagService
}

func NewTagController(service service.TagService) *TagController {
	return &TagController{
		TagService: service,
	}
}

// Create Controller
func (controller *TagController) Create(ctx *gin.Context) {
	createTagRequest := request.CreateTagRequest{}
	err := ctx.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.TagService.Create(createTagRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// Update Controller
func (controller *TagController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, errr := strconv.Atoi(tagId)
	helper.ErrorPanic(errr)

	updateTagRequest.Id = id

	controller.TagService.Update(updateTagRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// Delete Controller
func (controller *TagController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, errr := strconv.Atoi(tagId)
	helper.ErrorPanic(errr)

	controller.TagService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindById Controller
func (controller *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tag := controller.TagService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tag,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindAll Controller
func (controller *TagController) FindAll(ctx *gin.Context) {
	tags := controller.TagService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   tags,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
