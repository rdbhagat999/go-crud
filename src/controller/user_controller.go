package controller

import (
	"fmt"
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/service"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	UserService service.UserService
}

func userControllerPrintln(err error) {
	fmt.Println("userControllerPrintln: " + err.Error())
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
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users [POST]
func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	user, createErr := controller.UserService.Create(createUserRequest)

	if createErr != nil {

		userControllerPrintln(createErr)

		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: createErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Data:    user,
		Message: "User registered successfully",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// AuthUser godoc
// @Summary  Get authenticated user
// @Description  Get user in database
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/user [POST]
func (controller *UserController) AuthUser(ctx *gin.Context) {
	userId, userExists := ctx.Get("userId")
	fmt.Printf("AuthUserId: %v", userId)

	if !userExists {

		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusUnauthorized, webResponse)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return

	}

	user, userErr := controller.UserService.FindById(userId.(int))

	if userErr != nil {
		userControllerPrintln(userErr)
		// helper.ErrorPanic(userErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: userErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// Logout godoc
// @Summary  Logout authenticated user
// @Description  Logout user in database
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/Logout [POST]
func (controller *UserController) Logout(ctx *gin.Context) {

	ctx.SetCookie("jwt", "", int(-1), "/", "localhost", false, true)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Data:    nil,
		Message: "logout successfull",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// LoginUser godoc
// @Summary  Login user
// @Description  Login user in database
// @Param  user body request.LoginUserRequest true "Login user"
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/login [POST]
func (controller *UserController) Login(ctx *gin.Context) {
	loginUserRequest := request.LoginUserRequest{}
	err := ctx.ShouldBindJSON(&loginUserRequest)
	helper.ErrorPanic(err)

	user, userErr := controller.UserService.Login(loginUserRequest)

	if userErr != nil {
		userControllerPrintln(userErr)
		// helper.ErrorPanic(userErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: userErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		// Also fixed dates can be used for the NumericDate
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		Issuer:    strconv.Itoa(user.ID),
		// ID:        strconv.Itoa(int(user.ID)),
	})

	println("TOKEN_SECRET", os.Getenv("TOKEN_SECRET"))
	tokenString, tokenErr := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	// helper.ErrorPanic(tokenErr)

	if tokenErr != nil {
		userControllerPrintln(tokenErr)
		// helper.ErrorPanic(tokenErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: tokenErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	maxAge, _ := strconv.Atoi(os.Getenv("TOKEN_MAX_AGE"))

	ctx.SetCookie("jwt", tokenString, maxAge, "/", "localhost", false, true)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Data:    user,
		Message: "Login successful",
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// UpdateUser godoc
// @Summary  Update user
// @Description  Update and save user in database
// @Param  userId path string true "Update user by id"
// @Param  user body request.UpdateUserRequest true "Update user"
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/{userId} [PUT]
func (controller *UserController) Update(ctx *gin.Context) {

	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userId := ctx.Param("userId")
	id, paramErr := strconv.Atoi(userId)
	// helper.ErrorPanic(paramErr)

	if paramErr != nil {
		userControllerPrintln(paramErr)
		// helper.ErrorPanic(paramErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: paramErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	// updateUserRequest.ID = id

	user, userErr := controller.UserService.Update(id, updateUserRequest)

	if userErr != nil {
		userControllerPrintln(userErr)
		// helper.ErrorPanic(userErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: userErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// DeleteUser godoc
// @Summary  Delete user
// @Description  Delete user from database
// @Param  userId path string true "Delete user by id"
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/{userId} [DELETE]
func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, paramErr := strconv.Atoi(userId)
	helper.ErrorPanic(paramErr)

	if paramErr != nil {
		userControllerPrintln(paramErr)
		// helper.ErrorPanic(paramErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: paramErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

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
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/{userId} [GET]
func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, paramErr := strconv.Atoi(userId)
	// helper.ErrorPanic(paramErr)

	if paramErr != nil {
		userControllerPrintln(paramErr)
		// helper.ErrorPanic(paramErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: paramErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	user, userErr := controller.UserService.FindById(id)

	if userErr != nil {
		userControllerPrintln(userErr)
		// helper.ErrorPanic(userErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: userErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	if user.ID == 0 {

		webResponse := response.Response{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
			Data:    nil,
			Message: "user not found",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusNotFound, webResponse)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

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
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users [GET]
func (controller *UserController) FindAll(ctx *gin.Context) {
	users, usersErr := controller.UserService.FindAll()

	if usersErr != nil {
		userControllerPrintln(usersErr)
		// helper.ErrorPanic(usersErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: usersErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return

	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   users,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
