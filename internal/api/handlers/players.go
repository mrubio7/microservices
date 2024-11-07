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
	playerIds := c.Query("ids")
	nickname := c.Query("nickname")

	if playerIds == "" && nickname == "" {
		logger.Error("tried to get an empty id and nickname")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID or Nickname"))
		return
	}

	if playerIds != "" {
		ids := strings.Split(playerIds, ",")
		if len(ids) == 0 {
			logger.Error("no valid ids provided")
			c.JSON(http.StatusBadRequest, response.BuildError("No valid IDs provided"))
			return
		}

		res, err := ph.players_client.GetPlayer(c, &pb_players.GetPlayerRequest{FaceitId: ids})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, response.BuildError("Error getting players"))
			return
		}

		c.JSON(http.StatusOK, response.BuildOk("Ok", res))
	} else {
		res, err := ph.players_client.GetPlayerByNickname(c, &pb_players.GetPlayerByNicknameRequest{Nickname: nickname})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, response.BuildError("Error getting players"))
			return
		}

		c.JSON(http.StatusOK, response.BuildOk("Ok", res))
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

func (ph *Player_Handlers) NewLookingForTeam(c *gin.Context) {
	var payload struct {
		PlayerId     string   `json:"player_id"`
		InGameRole   []string `json:"in_game_role"`
		TimeTable    string   `json:"time_table"`
		OldTeams     string   `json:"old_teams"`
		PlayingYears int32    `json:"playing_years"`
		Description  string   `json:"description"`
	}

	identity, identityExist := c.Get("identity")
	if !identityExist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting NewLookingForTeam body"))
		return
	}

	req := &pb_players.NewPlayerLookingForTeam{
		InGameRole:   payload.InGameRole,
		TimeTable:    payload.TimeTable,
		OldTeams:     payload.OldTeams,
		PlayingYears: payload.PlayingYears,
		Description:  payload.Description,
		PlayerId:     payload.PlayerId,
		UserId:       int32(identity.(int)),
	}

	res, err := ph.players_client.UpdateLookingForTeam(c, req)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Internal error"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (ph *Player_Handlers) GetAllLookingForTeam(c *gin.Context) {
	res, err := ph.players_client.GetAllLookingForTeam(c, nil)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting all looking for team"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (ph *Player_Handlers) DeleteLookingForTeam(c *gin.Context) {
	var payload struct {
		PlayerId string `json:"player_id"`
	}

	identity, identityExist := c.Get("identity")
	if !identityExist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting DeleteLookingForTeam body"))
		return
	}

	res, err := ph.players_client.DeleteLookingForTeam(c, &pb_players.DeleteLookingForTeamRequest{PlayerId: payload.PlayerId, UserId: int32(identity.(int))})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Internal error"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}
