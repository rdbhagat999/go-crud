package external

import (
	"encoding/json"
	"fmt"
	"go-crud/src/data/response"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

const baseURL = "https://dummyjson.com/carts/user"

type ExternalController struct{}

func NewExternalController() *ExternalController {
	return &ExternalController{}
}

func ExternalControllerPrintln(err error) {
	fmt.Println("ExternalControllerPrintln: " + err.Error())
}

// GetCartByUserId godoc
// @Summary  Get cart by userId
// @Description  Fetches cart by userId from external API
// @Accept json
// @Produce  json
// @Success  200 {object} response.Response{}
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

	url := baseURL + "/" + userIdString
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
