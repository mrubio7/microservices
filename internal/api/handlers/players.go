package handlers

import (
	"ibercs/internal/api/requests"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_players "ibercs/proto/players"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Players_Handlers struct {
	players_client pb_players.PlayerServiceClient
}

func NewPlayersHandlers(client pb_players.PlayerServiceClient) *Players_Handlers {
	return &Players_Handlers{
		players_client: client,
	}
}

func (h *Players_Handlers) GetAll(c *gin.Context) {
	players, err := h.players_client.GetAllPlayers(c, &pb_players.Empty{})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", players))
}

func (h *Players_Handlers) Get(c *gin.Context) {
	nickname := c.Query("nickname")
	playerIds := c.Query("ids")

	if nickname == "" && playerIds == "" {
		logger.Error("tried to get an empty nickname")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid nickname"))
		return
	}

	if nickname != "" {
		res, err := h.players_client.GetPlayerByNickname(c, &pb_players.GetPlayerByNicknameRequest{Nickname: nickname})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting player by nickname"))
			return
		}

		c.JSON(http.StatusOK, response.BuildOk("Ok", res))
	} else {
		ids := strings.Split(playerIds, ",")
		if len(ids) == 0 {
			logger.Error("no valid ids provided")
			c.JSON(http.StatusBadRequest, response.BuildError("No valid IDs provided"))
			return
		}
		res, err := h.players_client.GetPlayersByFaceitId(c, &pb_players.GetPlayerRequest{FaceitId: ids})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting players by ids"))
			return
		}

		c.JSON(http.StatusOK, response.BuildOk("Ok", res))
	}
}

// Prominent players
func (h *Players_Handlers) GetProminentPlayers(c *gin.Context) {
	players, err := h.players_client.GetProminentPlayers(c, &pb_players.Empty{})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting prominent players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", players))
}

// Looking for team
func (h *Players_Handlers) GetLookingForTeamPlayers(c *gin.Context) {
	players, err := h.players_client.GetAllLookingForTeam(c, &pb_players.Empty{})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting looking for team players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", players))
}

func (h *Players_Handlers) CreateLookingForTeam(c *gin.Context) {
	var req requests.CreateLookingForTeam

	identity, identityExist := c.Get("identity")
	if !identityExist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting NewLookingForTeam body"))
		return
	}

	pbReq, err := req.ToProto(int32(identity.(int)))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error converting NewLookingForTeam to proto"))
		return
	}

	res, err := h.players_client.CreateLookingForTeam(c, pbReq)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error creating looking for team"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Players_Handlers) UpdateLookingForTeam(c *gin.Context) {
	var req requests.CreateLookingForTeam

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting PlayerLookingForTeam body"))
		return
	}

	identity, identityExist := c.Get("identity")
	if !identityExist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	pbReq, err := req.ToProto(int32(identity.(int)))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error converting NewLookingForTeam to proto"))
		return
	}

	res, err := h.players_client.UpdateLookingForTeam(c, pbReq)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error updating looking for team"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Players_Handlers) DeleteLookingForTeam(c *gin.Context) {
	var req requests.DeleteLookingForTeam

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting PlayerLookingForTeam body"))
		return
	}

	identity, identityExist := c.Get("identity")
	if !identityExist {
		c.JSON(http.StatusUnauthorized, response.BuildError("Unauthorized"))
		return
	}

	pb, err := req.ToProto(identity.(int))
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error converting NewLookingForTeam to proto"))
		return
	}

	_, err = h.players_client.DeleteLookingForTeam(c, pb)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error deleting looking for team"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", nil))
}
