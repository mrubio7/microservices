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
		// Define los endpoints que no deben ser cacheados
		noCachePaths := map[string]bool{
			"/api/v1/auth/callback": true,
		}

		// Verifica si el endpoint actual debe omitirse de la cache
		if _, shouldSkip := noCachePaths[ctx.Request.URL.Path]; shouldSkip {
			ctx.Next()
			return
		}

		// Construye la clave de cache para el endpoint actual
		cacheKey := buildCacheKey(ctx.Request.URL)

		// Intenta obtener datos en cache
		if cachedData, found := c.Get(cacheKey); found {
			ctx.Data(http.StatusOK, "application/json", cachedData.([]byte))
			ctx.Abort()
			return
		}

		// Intercepta la respuesta para almacenarla en cache si corresponde
		writer := &cachedWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: ctx.Writer}
		ctx.Writer = writer

		ctx.Next()

		// Verifica si la respuesta fue exitosa y si no es una solicitud autenticada
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
