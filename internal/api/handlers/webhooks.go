package handlers

import (
	"ibercs/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Webhooks_Handlers struct {
	db *gorm.DB
}

func NewWebhooksHandlers(db *gorm.DB) *Webhooks_Handlers {

	return &Webhooks_Handlers{
		db: db,
	}
}

func (h *Webhooks_Handlers) AllstarClipProcessed(c *gin.Context) {
	// Print the request body as string it comes in json
	body, err := c.GetRawData()
	if err != nil {
		logger.Error("Error reading request body")
		return
	}
	logger.Info(string(body))
}

/*
var clip webhooks.AllstarClipProcessed
	if err := c.BindJSON(&clip); err != nil {
		logger.Error("Error binding json", err)
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid json"))
		return
	}

	if err := h.db.Model(&webhooks.AllstarClipProcessed{}).Create(&clip).Error; err != nil {
		logger.Error("Error creating clip", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error creating clip"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Clip processed", nil))
*/
