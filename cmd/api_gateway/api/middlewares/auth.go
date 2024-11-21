package middlewares

import (
	"ibercs/internal/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationToken := ctx.GetHeader("Authorization")
		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

		if idToken == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
			ctx.Abort()
			return
		}

		var session *model.UserSessionModel
		if err := db.Model(&model.UserSessionModel{}).Where("session_id = ?", idToken).First(&session).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Cache-Control")
		ctx.Set("identity", session.UserID)
		ctx.Set("token", idToken)
		ctx.Next()
	}
}
