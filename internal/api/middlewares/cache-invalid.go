package middlewares

import (
	"ibercs/pkg/cache"

	"github.com/gin-gonic/gin"
)

func getCacheKeysToInvalidate(endpoint string) []string {
	switch endpoint {
	case "/api/v1/teams/new":
		return []string{"/api/v1/teams/get-all"}
	case "/api/v1/players/looking-for-team":
		return []string{"/api/v1/players/get-looking-for-team"}

	default:
		return []string{}
	}
}

func CacheInvalidationMiddleware(c *cache.Cache) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Guarda el path del request actual
		requestPath := ctx.Request.URL.Path

		// Continúa al siguiente middleware/controlador
		ctx.Next()

		// Después de procesar la petición, invalidamos la caché si es necesario
		keysToInvalidate := getCacheKeysToInvalidate(requestPath)
		for _, key := range keysToInvalidate {
			c.Delete(key)
		}
	}
}
