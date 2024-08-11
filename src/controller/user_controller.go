package controller

import (
	"fmt"
	"go-crud/src/constants"
	"go-crud/src/data/request"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/model"
	"go-crud/src/service"
	"net/http"
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
// @Summary  Register user
// @Description  Save user in database
// @Param  user body request.CreateUserRequest true "Create user"
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /auth/register [POST]
func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.CreateUserRequest{}
	createUserRequest.RoleID = 1
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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Data:    user,
		Message: "User registered successfuly",
	}

	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// AuthUser godoc
// @Summary  Get authenticated user
// @Description  Get user in database
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /users/authuser [GET]
func (controller *UserController) AuthUser(ctx *gin.Context) {
	userId, userExists := ctx.Get("userId")
	roleId, roleExists := ctx.Get("roleId")
	fmt.Printf("AuthUserId: %v", userId)
	fmt.Printf("RoleId: %v", roleId)

	if !userExists || !roleExists {

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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   user,
	}

	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// Logout godoc
// @Summary  Logout user
// @Description  Logout user in database
// @Accept json
// @Produce  json
// @Tag  user
// @Success  200 {object} response.Response{}
// @Router  /auth/logout [POST]
func (controller *UserController) Logout(ctx *gin.Context) {

	ctx.SetCookie(constants.AUTH_COOKIE_NAME, "", int(-1), "/", "localhost", false, true)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Data:    nil,
		Message: "logout successful",
	}

	// ctx.Header("Content-Type", "application/json")
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
// @Router  /auth/login [POST]
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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	claims := model.MyCustomJWTClaims{
		UserID: strconv.Itoa(user.ID),
		RoleID: strconv.Itoa(user.RoleID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "go-crud",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, tokenErr := token.SignedString([]byte(helper.GetEnvVariable(constants.TOKEN_SECRET)))
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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	maxAge, _ := strconv.Atoi(helper.GetEnvVariable(constants.TOKEN_MAX_AGE))

	ctx.SetCookie(constants.AUTH_COOKIE_NAME, "bearer "+tokenString, maxAge, "/", "localhost", false, true)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Data:    user,
		Message: "Login successful",
	}

	// ctx.Header("Content-Type", "application/json")
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
	ctx_user_Id, userIdExists := ctx.Get("userId")
	roleId, roleExists := ctx.Get("roleId")
	user_id, userParamErr := strconv.Atoi(userId)
	// helper.ErrorPanic(paramErr)

	fmt.Printf("user_id: %v", user_id)
	fmt.Printf("ctx_user_Id: %v", ctx_user_Id)
	fmt.Printf("RoleId: %v", roleId)

	if !userIdExists || !roleExists {
		// helper.ErrorPanic(userIdExists)
		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return

	}

	if userParamErr != nil {
		userControllerPrintln(userParamErr)
		// helper.ErrorPanic(userParamErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: userParamErr.Error(),
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	// updateUserRequest.ID = id

	user, userErr := controller.UserService.Update(user_id, updateUserRequest)

	if userErr != nil {
		userControllerPrintln(userErr)
		// helper.ErrorPanic(userErr)
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: userErr.Error(),
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
		Data:   user,
	}

	// ctx.Header("Content-Type", "application/json")
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
	ctx_user_Id, userIdExists := ctx.Get("userId")
	roleId, roleExists := ctx.Get("roleId")
	userId := ctx.Param("userId")
	id, paramErr := strconv.Atoi(userId)

	fmt.Printf("id: %v", id)
	fmt.Printf("ctx_user_Id: %v", ctx_user_Id)
	fmt.Printf("RoleId: %v", roleId)

	if !userIdExists || !roleExists {
		// helper.ErrorPanic(userIdExists)
		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: http.StatusText(http.StatusUnauthorized),
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return

	}
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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	controller.UserService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   nil,
	}

	// ctx.Header("Content-Type", "application/json")
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
	user_id, userExists := ctx.Get("userId")
	roleId, roleExists := ctx.Get("roleId")
	fmt.Printf("AuthUserId: %v", user_id)
	fmt.Printf("RoleId: %v", roleId)

	if !userExists || !roleExists {

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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	if user.ID == 0 {

		webResponse := response.Response{
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
			Data:    nil,
			Message: "user not found",
		}

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusNotFound, webResponse)
		// ctx.AbortWithStatus(http.StatusNotFound)
		ctx.AbortWithStatusJSON(http.StatusNotFound, webResponse)

		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	}

	// ctx.Header("Content-Type", "application/json")
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

	user_id, userExists := ctx.Get("userId")
	roleId, roleExists := ctx.Get("roleId")
	fmt.Printf("AuthUserId: %v", user_id)
	fmt.Printf("RoleId: %v", roleId)

	if !userExists || !roleExists {

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

	intRole := roleId.(int)

	if intRole == 0 {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: http.StatusText(http.StatusBadRequest),
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return
	}

	if intRole != 2 {
		webResponse := response.Response{
			Code:    http.StatusUnauthorized,
			Status:  http.StatusText(http.StatusUnauthorized),
			Data:    nil,
			Message: "Only admins can access user list",
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

		return
	}

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

		// ctx.Header("Content-Type", "application/json")
		// ctx.JSON(http.StatusBadRequest, webResponse)
		// ctx.AbortWithStatus(http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return

	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   users,
	}

	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
