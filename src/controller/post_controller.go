package controller

import (
	"fmt"
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/model"
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

func postControllerPrintln(err error) {
	fmt.Println("postControllerPrintln: " + err.Error())
}

// CreatePost godoc
// @Summary  Create post
// @Description  Save post in database
// @Param  post body request.CreatePostRequest true "Create post"
// @Accept json
// @Produce  json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts [POST]
func (controller *PostController) Create(ctx *gin.Context) {

	userId, userExists := ctx.Get("userId")
	fmt.Printf("AuthUserId: %v", userId)

	if !userExists {

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: http.StatusText(http.StatusBadRequest),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	if userId == "" || userId == nil || userId == 0 {

		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusUnauthorized, webResponse)
		// ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return

	}

	createPostRequest := request.CreatePostRequest{}
	err := ctx.ShouldBindJSON(&createPostRequest)

	if err != nil {
		helper.ErrorPanic(err)
	}

	createPost := model.Post{}
	createPost.UserID = userId.(int)
	createPost.Title = createPostRequest.Title
	createPost.Body = createPostRequest.Body

	post, createErr := controller.PostService.Create(createPost)

	if createErr != nil {
		postControllerPrintln(createErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: createErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

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
// @Accept json
// @Produce  json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts/{postId} [PUT]
func (controller *PostController) Update(ctx *gin.Context) {

	userId, userExists := ctx.Get("userId")
	fmt.Printf("AuthUserId: %v", userId)

	if !userExists {

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: http.StatusText(http.StatusBadRequest),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	if userId == "" || userId == nil || userId == 0 {

		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusUnauthorized, webResponse)
		// ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return

	}

	updatePostRequest := request.UpdatePostRequest{}
	err := ctx.ShouldBindJSON(&updatePostRequest)
	helper.ErrorPanic(err)

	postId := ctx.Param("postId")
	id, paramErr := strconv.Atoi(postId)
	// helper.ErrorPanic(paramErr)

	if paramErr != nil {
		postControllerPrintln(paramErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: paramErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return
	}

	found, foundErr := controller.PostService.FindById(id)

	if foundErr != nil {
		postControllerPrintln(foundErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: foundErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return
	}

	if found.UserID != userId {
		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusUnauthorized, webResponse)
		// ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return
	}

	post, updateErr := controller.PostService.Update(id, updatePostRequest)

	if updateErr != nil {
		postControllerPrintln(updateErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: updateErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

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
// @Accept json
// @Produce  json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts/{postId} [DELETE]
func (controller *PostController) Delete(ctx *gin.Context) {

	userId, userExists := ctx.Get("userId")
	fmt.Printf("AuthUserId: %v", userId)

	if !userExists {

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: http.StatusText(http.StatusBadRequest),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	if userId == "" || userId == nil || userId == 0 {

		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusUnauthorized, webResponse)
		// ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return

	}

	postId := ctx.Param("postId")
	id, paramErr := strconv.Atoi(postId)
	// helper.ErrorPanic(paramErr)

	if paramErr != nil {
		postControllerPrintln(paramErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: paramErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	post, findErr := controller.PostService.FindById(id)

	if findErr != nil {
		postControllerPrintln(findErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: findErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	if post.UserID != userId.(int) {
		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusUnauthorized, webResponse)
		// ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return

	}

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
// @Accept json
// @Produce  json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts/{postId} [GET]
func (controller *PostController) FindById(ctx *gin.Context) {
	postId := ctx.Param("postId")
	id, paramErr := strconv.Atoi(postId)
	// helper.ErrorPanic(err)

	if paramErr != nil {
		postControllerPrintln(paramErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: paramErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	post, findErr := controller.PostService.FindById(id)

	if findErr != nil {
		postControllerPrintln(findErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: findErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	if post.ID == 0 {

		webResponse := response.Response{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
			Data:    nil,
			Message: "post not found",
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusNotFound, webResponse)
		// ctx.AbortWithStatus(http.StatusNotFound)
		ctx.AbortWithStatusJSON(http.StatusNotFound, webResponse)

		return
	}

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
// @Description Returns a list of post
// @Accept json
// @Produce  json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts [GET]
func (controller *PostController) FindAll(ctx *gin.Context) {
	posts, listErr := controller.PostService.FindAll()

	if listErr != nil {
		postControllerPrintln(listErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: listErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   posts,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllByUserId godoc
// @Summary  Get all post by userId
// @Description Returns a list of post
// @Accept json
// @Produce  json
// @Post  post
// @Success  200 {object} response.Response{}
// @Router  /posts/userposts [GET]
func (controller *PostController) FindAllByUserId(ctx *gin.Context) {
	userId, userExists := ctx.Get("userId")
	fmt.Printf("AuthUserId: %v", userId)

	if !userExists {

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: http.StatusText(http.StatusBadRequest),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	if userId == "" || userId == nil || userId == 0 {

		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusUnauthorized, webResponse)
		// ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return

	}

	posts, listErr := controller.PostService.FindAllByUserId(userId.(int))

	if listErr != nil {
		postControllerPrintln(listErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: listErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   posts,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
