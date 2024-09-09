package external

import (
	"encoding/json"
	"fmt"
	"go-crud/src/data/response"
	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
)

const baseURL = "https://dummyjson.com"

type ExternalController struct{}

func NewExternalController() *ExternalController {
	return &ExternalController{}
}

func ExternalControllerPrintln(err error) {
	fmt.Println("ExternalControllerPrintln: " + err.Error())
}

// GetCartByUserId godoc
// @ID GetCartByUserId
// @Message Get cart by userId
// @Summary  Get cart by userId
// @Description  Fetches cart by userId from external API
// @Accept json
// @Produce  json
// @Tags cart
// @Success  200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router  /cart/cartbyuser [GET]
func (controller *ExternalController) GetCartByUserId(ctx *gin.Context) {

	userId, userExists := ctx.Get("userId")
	fmt.Println("AuthUserId:", userId)

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

	userIdInt := userId.(int)
	println("userIdInt: ", userIdInt)

	userIdString := strconv.Itoa(userIdInt)
	println("userIdString: ", userIdString)

	url := baseURL + "/carts/user/" + userIdString
	println("url: ", url)

	// var cart CartResponse
	resp, err := http.Get(url)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		webResponse := response.Response{
			Code:    resp.StatusCode,
			Status:  http.StatusText(resp.StatusCode),
			Data:    nil,
			Message: http.StatusText(resp.StatusCode),
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, webResponse)

		return
	}

	var cart CartResponse

	err = json.NewDecoder(resp.Body).Decode(&cart)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	if cart.Total == 0 {
		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}

		// ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)

		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   cart,
	}

	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// AddCardByUserId godoc
// @ID AddCardByUserId
// @Message Create cart by userId
// @Summary Create cart by userId
// @Description Create cart by userId from external API
// @Accept json
// @Produce  json
// @Tags cart
// @Success  200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router  /cart/addusercart [POST]
func (controller *ExternalController) AddCardByUserId(ctx *gin.Context) {

	userId, userExists := ctx.Get("userId")
	fmt.Println("AuthUserId:", userId)

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

	userIdInt := userId.(int)
	println("userIdInt: ", userIdInt)

	userIdString := strconv.Itoa(userIdInt)
	println("userIdString: ", userIdString)

	url := baseURL + "/carts/add"
	println("url: ", url)

	products := make([]CartProduct, 10)
	for i := range products {
		products[i].ID = i + 97
		products[i].Quantity = i + 2
	}

	cartToAdd := AddCartRequest{
		UserID:   userIdInt,
		Products: products,
	}

	print("cartToAdd 0: ", cartToAdd.Products[0].ID, " ", cartToAdd.Products[0].Quantity, "\n")
	println("UserID: ", cartToAdd.UserID)
	print("cartToAdd 1: ", cartToAdd.Products[1].ID, " ", cartToAdd.Products[1].Quantity, "\n")
	println("")

	dataBytes, marshalError := json.Marshal(cartToAdd)

	if marshalError != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: marshalError.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	jsonDataString := string(dataBytes)

	jsonDataReader := strings.NewReader(jsonDataString)

	resp, err := http.Post(url, "application/json", jsonDataReader)

	if err != nil {
		webResponse := response.Response{
			Code:    resp.StatusCode,
			Status:  http.StatusText(resp.StatusCode),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, webResponse)

		return
	}

	defer resp.Body.Close()

	println("resp.StatusCode: ", resp.StatusCode)

	if resp.StatusCode != http.StatusCreated {
		webResponse := response.Response{
			Code:    resp.StatusCode,
			Status:  http.StatusText(resp.StatusCode),
			Data:    nil,
			Message: http.StatusText(resp.StatusCode),
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, webResponse)

		return
	}

	var cart AddCartResponse

	err = json.NewDecoder(resp.Body).Decode(&cart)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	if cart.Total == 0 {
		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}

		// ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)

		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   cart,
	}

	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// UpdateCardByUserId godoc
// @ID UpdateCardByUserId
// @Message Update cart by userId
// @Summary Update cart by cartId
// @Description Update cart by cartId from external API
// @Accept json
// @Produce  json
// @Tags cart
// @Success  200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router  /cart/updateusercart [PUT]
func (controller *ExternalController) UpdateCardByUserId(ctx *gin.Context) {

	userId, userExists := ctx.Get("userId")
	fmt.Println("AuthUserId:", userId)

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

	userIdInt := userId.(int)
	println("userIdInt: ", userIdInt)

	userIdString := strconv.Itoa(userIdInt)
	println("userIdString: ", userIdString)

	url := baseURL + "/carts/" + userIdString
	println("url: ", url)

	products := make([]CartProduct, 10)
	for i := range products {
		products[i].ID = i + 97
		products[i].Quantity = i + 2
	}

	cartToAdd := AddCartRequest{
		UserID:   userIdInt,
		Products: products,
	}

	print("cartToAdd 0: ", cartToAdd.Products[0].ID, " ", cartToAdd.Products[0].Quantity, "\n")
	println("UserID: ", cartToAdd.UserID)
	print("cartToAdd 1: ", cartToAdd.Products[1].ID, " ", cartToAdd.Products[1].Quantity, "\n")
	println("")

	dataBytes, marshalError := json.Marshal(cartToAdd)

	if marshalError != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: marshalError.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	jsonDataString := string(dataBytes)

	jsonDataReader := strings.NewReader(jsonDataString)

	updateCartReq, updateCartReqErr := http.NewRequest(http.MethodPut, url, jsonDataReader)
	updateCartReq.Header.Set("Content-Type", "application/json")

	if updateCartReqErr != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: updateCartReqErr.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	client := http.Client{}
	resp, err := client.Do(updateCartReq)

	if err != nil {
		webResponse := response.Response{
			Code:    resp.StatusCode,
			Status:  http.StatusText(resp.StatusCode),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, webResponse)

		return
	}

	defer resp.Body.Close()

	println("resp.StatusCode: ", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		webResponse := response.Response{
			Code:    resp.StatusCode,
			Status:  http.StatusText(resp.StatusCode),
			Data:    nil,
			Message: http.StatusText(resp.StatusCode),
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, webResponse)

		return
	}

	var cart UpdateCartResponse

	err = json.NewDecoder(resp.Body).Decode(&cart)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	if cart.Total == 0 {
		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: http.StatusText(http.StatusOK),
			Data:   nil,
		}

		// ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)

		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   cart,
	}

	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

// DeleteCardById godoc
// @ID DeleteCardById
// @Message Delete cart by userId
// @Summary Delete cart by cartId
// @Description  Delete cart by cartId from external API
// @Accept json
// @Produce  json
// @Tags cart
// @Success  200 {object} response.Response{}
// @Failure 400 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router  /cart/deletecart [DELETE]
func (controller *ExternalController) DeleteCardById(ctx *gin.Context) {

	userId, userExists := ctx.Get("userId")
	fmt.Println("AuthUserId:", userId)

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

	userIdInt := userId.(int)
	println("userIdInt: ", userIdInt)

	userIdString := strconv.Itoa(userIdInt)
	println("userIdString: ", userIdString)

	url := baseURL + "/carts/" + userIdString
	println("url: ", url)

	deleteCartReq, deleteCartReqErr := http.NewRequest(http.MethodPut, url, nil)
	deleteCartReq.Header.Set("Content-Type", "application/json")

	if deleteCartReqErr != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: deleteCartReqErr.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	client := http.Client{}
	resp, err := client.Do(deleteCartReq)

	if err != nil {
		webResponse := response.Response{
			Code:    resp.StatusCode,
			Status:  http.StatusText(resp.StatusCode),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, webResponse)

		return
	}

	defer resp.Body.Close()

	println("resp.StatusCode: ", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		webResponse := response.Response{
			Code:    resp.StatusCode,
			Status:  http.StatusText(resp.StatusCode),
			Data:    nil,
			Message: http.StatusText(resp.StatusCode),
		}

		ctx.AbortWithStatusJSON(resp.StatusCode, webResponse)

		return
	}

	var cart DeleteCartResponse

	err = json.NewDecoder(resp.Body).Decode(&cart)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Data:    nil,
			Message: err.Error(),
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    cart,
		Message: "Cart deletd successfully",
	}

	// ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
