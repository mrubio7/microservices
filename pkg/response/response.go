package response

import (
	"ibercs/pkg/logger"

	"github.com/gin-gonic/gin"
)

func BuildError(err string) gin.H {
	logger.Error(err)
	return gin.H{
		"error":   err,
		"message": nil,
		"data":    nil,
	}
}

func BuildOk(msg string, data any) gin.H {
	return gin.H{
		"error":   nil,
		"message": msg,
		"data":    data,
	}
}
