package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_players "ibercs/proto/players"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Player_Handlers struct {
	players_client pb_players.PlayerServiceClient
}

func NewPlayersHandlers(playersClient pb_players.PlayerServiceClient) *Player_Handlers {
	return &Player_Handlers{
		players_client: playersClient,
	}
}

func (ph *Player_Handlers) GetAllPlayers(c *gin.Context) {
	res, err := ph.players_client.GetPlayers(c, nil)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting all players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (ph *Player_Handlers) GetProminentPlayers(c *gin.Context) {
	res, err := ph.players_client.GetProminentPlayers(c, nil)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting prominent players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}
