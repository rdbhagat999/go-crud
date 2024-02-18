package controller

import (
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/service"
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

// CreateTag godoc
// @Summary  Create tag
// @Description  Save tag in database
// @Param  tag body request.CreateTagRequest true "Create tag"
// @Produce  application/json
// @Tag  tag
// @Success  200 {object} response.Response{}
// @Router  /tags [POST]
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

// UpdateTag godoc
// @Summary  Update tag
// @Description  Update and save tag in database
// @Param  tagId path string true "Update tag by id"
// @Param  tag body request.UpdateTagRequest true "Update tag"
// @Produce  application/json
// @Tag  tag
// @Success  200 {object} response.Response{}
// @Router  /tags/{tagId} [PUT]
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

// DeleteTag godoc
// @Summary  Delete tag
// @Description  Delete tag from database
// @Param  tagId path string true "Delete tag by id"
// @Produce  application/json
// @Tag  tag
// @Success  200 {object} response.Response{}
// @Router  /tags/{tagId} [DELETE]
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

// FindByIdTag godoc
// @Summary  Get a single tag by its id
// @Description  Returns a single tag when tagId maches id
// @Param  tagId path string true "Find tag by id"
// @Produce  application/json
// @Tag  tag
// @Success  200 {object} response.Response{}
// @Router  /tags/{tagId} [GET]
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

// FindAllTag godoc
// @Summary  Get all tags
// @Description  Returns a list of tags
// @Produce  application/json
// @Tag  tags
// @Success  200 {object} response.Response{}
// @Router  /tags [GET]
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
