package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	"ibercs/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type State_Handlers struct {
	serviceState *service.State
}

func NewStateHandlers(db *gorm.DB) *State_Handlers {
	svc := service.NewStateService(db)
	return &State_Handlers{
		serviceState: svc,
	}
}

func (sh *State_Handlers) GetState(c *gin.Context) {
	state := sh.serviceState.GetState()
	if state == nil {
		logger.Error("Unable to get state")
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to get state"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", state))
}

func (sh *State_Handlers) GetLastPlayerUpdate(c *gin.Context) {
	state := sh.serviceState.GetState()
	if state == nil {
		logger.Error("Unable to get state")
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to get last players update"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", state.LastPlayerUpdate.Time))
}
