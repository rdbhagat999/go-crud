package middlewares

import (
	"fmt"
	"go-crud/src/constants"
	"go-crud/src/controller"
	"go-crud/src/data/response"
	"go-crud/src/helper"
	"go-crud/src/model"
	"net/http"
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

		jwtString, errNoCookie := ctx.Cookie("JWT")
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

		token, parseErr := jwt.ParseWithClaims(tokenFields[1], &model.MyCustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(helper.GetEnvVariable(constants.TOKEN_SECRET)), nil
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

		if claims, ok := token.Claims.(*model.MyCustomJWTClaims); ok {

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
			fmt.Println("claims.user_id: ", claims.UserID)
			fmt.Println("claims.role_id: ", claims.RoleID)
			intId, intIdErr := strconv.Atoi(claims.UserID)
			intRole, roleIdErr := strconv.Atoi(claims.RoleID)
			fmt.Println("intId: ", intId)
			fmt.Println("intRole: ", intRole)

			userId = intId

			if intId == 0 || intRole == 0 {
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

			if intIdErr != nil {
				webResponse = response.Response{
					Code:    http.StatusBadRequest,
					Status:  http.StatusText(http.StatusBadRequest),
					Data:    nil,
					Message: intIdErr.Error(),
				}

				// ctx.Header("Content-Type", "application/json")
				// ctx.JSON(http.StatusBadRequest, webResponse)
				// ctx.AbortWithStatus(http.StatusUnauthorized)
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)

				return
			}

			if roleIdErr != nil {
				webResponse = response.Response{
					Code:    http.StatusBadRequest,
					Status:  http.StatusText(http.StatusBadRequest),
					Data:    nil,
					Message: roleIdErr.Error(),
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
		ctx.Header("X-ROLE-ID", strconv.Itoa(user.RoleID))
		ctx.Set("userId", user.ID)
		ctx.Set("roleId", user.RoleID)

		ctx.Next()
	}
}
