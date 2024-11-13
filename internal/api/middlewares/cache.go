package middlewares

import (
	"bytes"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"ibercs/pkg/cache" // Ajusta el import al path correcto

	"github.com/gin-gonic/gin"
)

// ResponseWriter personalizado para capturar la respuesta
type cachedWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *cachedWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // Captura los datos escritos en el body
	return w.ResponseWriter.Write(b)
}

var nonCacheablePatterns = []*regexp.Regexp{
	regexp.MustCompile(`^/api/v1/users/get/.*$`),   // Coincide con /api/v1/users/get?...
	regexp.MustCompile(`^/api/v1/matches/get/.*$`), // Coincide con /api/v1/players/get?...
}

func isCacheableRequest(requestPath string) bool {
	for _, pattern := range nonCacheablePatterns {
		if pattern.MatchString(requestPath) {
			return false
		}
	}
	return true
}

func CacheMiddleware(c *cache.Cache, ttl time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestPath := ctx.Request.URL.Path

		// Verifica si la solicitud debe ser cacheada
		if !isCacheableRequest(requestPath) {
			// Omite el cacheo para esta solicitud
			ctx.Next()
			return
		}

		// Construye la clave de caché
		cacheKey := buildCacheKey(ctx.Request.URL)

		// Intenta obtener la respuesta de la caché
		if cachedData, found := c.Get(cacheKey); found {
			ctx.Data(http.StatusOK, "application/json", cachedData.([]byte))
			ctx.Abort()
			return
		}

		// Captura la respuesta
		writer := &cachedWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: ctx.Writer}
		ctx.Writer = writer

		ctx.Next()

		_, isAuthReq := ctx.Get("identity")

		if ctx.Writer.Status() == http.StatusOK && !isAuthReq {
			responseData := writer.body.Bytes()
			c.Set(cacheKey, responseData, ttl)
		}
	}
}

func buildCacheKey(u *url.URL) string {
	// Tomar el path y los parámetros de la URL
	path := u.Path
	query := u.RawQuery

	// Concatenar el path con los parámetros (si existen) para formar la clave de caché
	if query != "" {
		return path + "?" + query
	}
	return path
}
