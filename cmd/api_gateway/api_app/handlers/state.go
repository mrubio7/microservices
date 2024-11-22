package handlers

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type State_Handlers struct {
	StateManager *managers.StateManager
}

func NewStateHandlers() *State_Handlers {
	config, err := config.Load()
	if err != nil {
		panic("cannot read config")
	}

	database := database.NewDatabase(config.StateDb)

	return &State_Handlers{
		StateManager: managers.NewStateManager(database.GetDB()),
	}
}

func (h *State_Handlers) GetState(c *gin.Context) {
	state, err := h.StateManager.Get()
	if err != nil {
		logger.Error("Error getting state: %v", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting state"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", state))
}

func (h *State_Handlers) GetUpdatePlayersState(c *gin.Context) {
	state, err := h.StateManager.Get()
	if err != nil {
		logger.Error("Error getting state: %v", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting state"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", state.PlayersLastUpdate.Time))
}
