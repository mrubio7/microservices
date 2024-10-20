package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_players "ibercs/proto/players"
	"net/http"
	"strings"

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

func (ph *Player_Handlers) GetPlayers(c *gin.Context) {
	// Recibe varios IDs separados por comas
	playerIds := c.Query("ids")
	if playerIds == "" {
		logger.Error("tried to get empty ids")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid IDs"))
		return
	}

	// Divide los IDs por comas
	ids := strings.Split(playerIds, ",")
	if len(ids) == 0 {
		logger.Error("no valid ids provided")
		c.JSON(http.StatusBadRequest, response.BuildError("No valid IDs provided"))
		return
	}

	// Llama al servicio con la lista de IDs
	res, err := ph.players_client.GetPlayer(c, &pb_players.GetPlayerRequest{FaceitId: ids})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
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
