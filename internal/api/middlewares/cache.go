package middlewares

import (
	"bytes"
	"net/http"
	"net/url"
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

func CacheMiddleware(c *cache.Cache, ttl time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cacheKey := buildCacheKey(ctx.Request.URL)

		if cachedData, found := c.Get(cacheKey); found {
			ctx.Data(http.StatusOK, "application/json", cachedData.([]byte))
			ctx.Abort()
			return
		}

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
