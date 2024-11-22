package middlewares

import (
	"ibercs/pkg/response"
	"ibercs/proto/users"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(usersServer users.UserServiceClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationToken := ctx.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			ctx.JSON(http.StatusBadRequest, response.BuildError("Id token not available"))
			ctx.Abort()
			return
		}

		session, err := usersServer.GetSessionById(ctx, &users.GetSessionByIdRequest{Token: idToken})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
			ctx.Abort()
			return
		}

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Cache-Control")
		ctx.Set("identity", session.UserId)
		ctx.Set("token", idToken)
		ctx.Next()
	}
}
