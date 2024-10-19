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
	resp, err := http.Get(fmt.Sprintf("%s/update", w.playersWorkerHost))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call update service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read update response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}

func (w *Workers_Handlers) Find(c *gin.Context) {
	size := c.DefaultQuery("size", "5000")

	url := fmt.Sprintf("%s/find?size=%s", w.playersWorkerHost, size)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call find service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read find response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}
