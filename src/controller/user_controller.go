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
// @Produce  application/json
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

// AuthUser godoc
// @Summary  Get authenticated user
// @Description  Get user in database
// @Param
// @Produce  application/json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/user [POST]
func (controller *UserController) AuthUser(ctx *gin.Context) {
	var webResponse response.Response
	var username string
	var expiresAt *jwt.NumericDate

	jwtString, cookieError := ctx.Cookie("jwt")
	// helper.ErrorPanic(cookieError)

	if cookieError != nil {
		userControllerPrintln(cookieError)

		webResponse = response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: cookieError.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)

		return
	}

	token, parseErr := jwt.ParseWithClaims(jwtString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if parseErr != nil {
		userControllerPrintln(parseErr)
		// helper.ErrorPanic(parseErr)

		webResponse = response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: parseErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return

	} else if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		fmt.Println(claims.Issuer)
		username = claims.Issuer
		expiresAt = claims.ExpiresAt
	} else {
		fmt.Println("unknown claims type, cannot proceed")

		webResponse = response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: "unknown claims type, cannot proceed",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	if username == "" {
		webResponse = response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: "invalid token",
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	if jwt.NewNumericDate(time.Now()).UnixNano() > expiresAt.UnixNano() {
		webResponse = response.Response{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   nil,
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return

	}

	user, userErr := controller.UserService.FindByUsername(username)

	if userErr != nil {
		userControllerPrintln(userErr)
		// helper.ErrorPanic(userErr)

		webResponse = response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: userErr.Error(),
		}

		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusBadRequest, webResponse)
		return

	}

	webResponse = response.Response{
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
// @Param
// @Produce  application/json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/Logout [POST]
func (controller *UserController) Logout(ctx *gin.Context) {

	ctx.SetCookie("jwt", "", time.Now().Add(-time.Hour*24).Second(), "/", "", false, true)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// LoginUser godoc
// @Summary  Login user
// @Description  Login user in database
// @Param  user body request.LoginUserRequest true "Login user"
// @Produce  application/json
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
		return

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		// Also fixed dates can be used for the NumericDate
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		Issuer:    user.Username,
		ID:        strconv.Itoa(int(user.ID)),
	})

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
		return

	}

	ctx.SetCookie("jwt", tokenString, time.Now().Add(time.Hour*24).Second(), "/", "", false, true)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
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
// @Produce  application/json
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
// @Produce  application/json
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
// @Produce  application/json
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
