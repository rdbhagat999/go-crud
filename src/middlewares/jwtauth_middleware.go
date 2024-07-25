package middlewares

import (
	"fmt"
	"go-crud/src/controller"
	"go-crud/src/data/response"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func middlewarePrintln(err error) {
	fmt.Println("middlewarePrintln: " + err.Error())
}

func JWTAuthMiddleware(controller *controller.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("JWTAuthMiddleware: " + ctx.Request.RequestURI)

		var webResponse response.Response
		var expiresAt *jwt.NumericDate
		var userId int

		jwtString, errNoCookie := ctx.Cookie("jwt")
		// helper.ErrorPanic(errNoCookie)

		if errNoCookie != nil {
			middlewarePrintln(errNoCookie)

			webResponse = response.Response{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   nil,
				// Message: errNoCookie.Error(),
				Message: "cookie not found",
			}

			// ctx.Header("Content-Type", "application/json")
			// ctx.JSON(http.StatusUnauthorized, webResponse)
			// ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

			return
		}

		token, parseErr := jwt.ParseWithClaims(jwtString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})

		if parseErr != nil {
			middlewarePrintln(parseErr)
			// helper.ErrorPanic(parseErr)

			webResponse = response.Response{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Data:   nil,
				// Message: parseErr.Error(),
				Message: "parseErr",
			}

			// ctx.Header("Content-Type", "application/json")
			// ctx.JSON(http.StatusBadRequest, webResponse)
			// ctx.AbortWithStatus(http.StatusBadRequest)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

			return

		}

		tokenFields := strings.Split(jwtString, " ")

		if len(tokenFields) < 2 {
			webResponse = response.Response{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   nil,
				// Message: parseErr.Error(),
				Message: "Invalid token",
			}

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

			return
		}

		authType := strings.ToLower(tokenFields[0])

		if authType != "bearer" {
			webResponse = response.Response{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Data:   nil,
				// Message: parseErr.Error(),
				Message: "Invalid token",
			}

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

			return
		}

		if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {

			expiresAt = claims.ExpiresAt

			if time.Now().After(expiresAt.Time) {
				webResponse = response.Response{
					Code:    http.StatusUnauthorized,
					Status:  http.StatusText(http.StatusUnauthorized),
					Data:    nil,
					Message: "token expired",
				}

				// ctx.Header("Content-Type", "application/json")
				// ctx.JSON(http.StatusUnauthorized, webResponse)
				// ctx.AbortWithStatus(http.StatusUnauthorized)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

				return

			}

			fmt.Println("claims.Issuer: ", claims.Issuer)
			intId, err := strconv.Atoi(claims.Issuer)
			fmt.Println("intId: ", intId)

			userId = intId

			if intId == 0 {
				webResponse = response.Response{
					Code:    http.StatusBadRequest,
					Status:  http.StatusText(http.StatusBadRequest),
					Data:    nil,
					Message: "invalid token",
				}

				// ctx.Header("Content-Type", "application/json")
				// ctx.JSON(http.StatusBadRequest, webResponse)
				// ctx.AbortWithStatus(http.StatusBadRequest)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

				return
			}

			if err != nil {
				webResponse = response.Response{
					Code:    http.StatusBadRequest,
					Status:  http.StatusText(http.StatusBadRequest),
					Data:    nil,
					Message: err.Error(),
				}

				// ctx.Header("Content-Type", "application/json")
				// ctx.JSON(http.StatusBadRequest, webResponse)
				// ctx.AbortWithStatus(http.StatusUnauthorized)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

				return
			}

		} else {
			fmt.Println("unknown claims type, cannot proceed")

			webResponse = response.Response{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Data:    nil,
				Message: "unknown claims type, cannot proceed",
			}

			// ctx.Header("Content-Type", "application/json")
			// ctx.JSON(http.StatusBadRequest, webResponse)
			// ctx.AbortWithStatus(http.StatusBadRequest)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

			return
		}

		user, userErr := controller.UserService.FindById(userId)

		if userErr != nil {
			middlewarePrintln(userErr)
			// helper.ErrorPanic(userErr)

			webResponse = response.Response{
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
				Data:    nil,
				Message: "userErr",
			}

			// ctx.Header("Content-Type", "application/json")
			// ctx.JSON(http.StatusBadRequest, webResponse)
			// ctx.AbortWithStatus(http.StatusBadRequest)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, webResponse)

			return

		}

		ctx.Header("X-USER-ID", strconv.Itoa(user.ID))
		ctx.Set("userId", user.ID)

		ctx.Next()
	}
}
