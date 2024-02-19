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

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

// CreateUser godoc
// @Summary  Create user
// @Description  Save user in database
// @Param  user body request.CreateUserRequest true "Create user"
// @Produce  application/json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users [POST]
func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	controller.UserService.Create(createUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// UpdateUser godoc
// @Summary  Update user
// @Description  Update and save user in database
// @Param  userId path string true "Update user by id"
// @Param  user body request.UpdateUserRequest true "Update user"
// @Produce  application/json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/{userId} [PUT]
func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, errr := strconv.Atoi(userId)
	helper.ErrorPanic(errr)

	updateUserRequest.ID = id

	controller.UserService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// DeleteUser godoc
// @Summary  Delete user
// @Description  Delete user from database
// @Param  userId path string true "Delete user by id"
// @Produce  application/json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/{userId} [DELETE]
func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, errr := strconv.Atoi(userId)
	helper.ErrorPanic(errr)

	controller.UserService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindByIdUser godoc
// @Summary  Get a single user by its id
// @Description  Returns a single user when userId maches id
// @Param  userId path string true "Find user by id"
// @Produce  application/json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/{userId} [GET]
func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ErrorPanic(err)

	user := controller.UserService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// FindAllUser godoc
// @Summary  Get all users
// @Description  Returns a list of users
// @Produce  application/json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users [GET]
func (controller *UserController) FindAll(ctx *gin.Context) {
	users := controller.UserService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   users,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
