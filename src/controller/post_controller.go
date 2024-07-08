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

type PostController struct {
	PostService service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{
		PostService: service,
	}
}

// CreatePost godoc
// @Summary  Create post
// @Description  Save post in database
// @Param  post body request.CreatePostRequest true "Create post"
// @Produce  application/json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts [POST]
func (controller *PostController) Create(ctx *gin.Context) {
	createPostRequest := request.CreatePostRequest{}
	err := ctx.ShouldBindJSON(&createPostRequest)
	helper.ErrorPanic(err)

	post := controller.PostService.Create(createPostRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   post,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// UpdatePost godoc
// @Summary  Update post
// @Description  Update and save post in database
// @Param  postId path string true "Update post by id"
// @Param  post body request.UpdatePostRequest true "Update post"
// @Produce  application/json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts/{postId} [PUT]
func (controller *PostController) Update(ctx *gin.Context) {
	updatePostRequest := request.UpdatePostRequest{}
	err := ctx.ShouldBindJSON(&updatePostRequest)
	helper.ErrorPanic(err)

	postId := ctx.Param("postId")
	id, errr := strconv.Atoi(postId)
	helper.ErrorPanic(errr)

	updatePostRequest.ID = id

	post := controller.PostService.Update(updatePostRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   post,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// DeletePost godoc
// @Summary  Delete post
// @Description  Delete post from database
// @Param  postId path string true "Delete post by id"
// @Produce  application/json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts/{postId} [DELETE]
func (controller *PostController) Delete(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, errr := strconv.Atoi(postId)
	helper.ErrorPanic(errr)

	controller.PostService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindByIdPost godoc
// @Summary  Get a single post by its id
// @Description  Returns a single post when postId maches id
// @Param  postId path string true "Find post by id"
// @Produce  application/json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts/{postId} [GET]
func (controller *PostController) FindById(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, err := strconv.Atoi(postId)
	helper.ErrorPanic(err)

	post := controller.PostService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   post,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindAllPost godoc
// @Summary  Get all post
// @Description  Returns a list of post
// @Produce  application/json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts [GET]
func (controller *PostController) FindAll(ctx *gin.Context) {
	posts := controller.PostService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   posts,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
