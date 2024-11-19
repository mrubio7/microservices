package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_players "ibercs/proto/players"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Players_Handlers struct {
	matches_client pb_players.PlayerServiceClient
}

func NewPlayersHandlers(client pb_players.PlayerServiceClient) *Players_Handlers {
	return &Players_Handlers{
		matches_client: client,
	}
}

func (h *Players_Handlers) GetAll(c *gin.Context) {
	players, err := h.matches_client.GetPlayers(c, &pb_players.Empty{})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting players"))
		return
	}

	c.JSON(http.StatusOK, players)
}
