package handlers

import (
	"fmt"
	"ibercs/pkg/config"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Workers_Handlers struct {
	playersWorkerHost string
}

func NewWorkersHandlers(cfg config.WorkersConfig) *Workers_Handlers {
	return &Workers_Handlers{
		playersWorkerHost: cfg.PlayersHost,
	}
}

func (w *Workers_Handlers) Update(c *gin.Context) {
	// Establecer los headers para SSE en la respuesta
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	// Realizar la solicitud al servicio de actualización
	resp, err := http.Get(fmt.Sprintf("%s/update", w.playersWorkerHost))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call update service"})
		return
	}
	defer resp.Body.Close()

	// Leer y retransmitir los eventos SSE
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read update response"})
			return
		}

		// Escribir los datos leídos directamente en la respuesta del cliente
		_, writeErr := c.Writer.Write(buf[:n])
		if writeErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write SSE to client"})
			return
		}

		// Hacer flush de los datos para enviarlos inmediatamente
		c.Writer.Flush()
	}
}
