package handlers

import (
	"ibercs/cmd/api_gateway/api_app/requests"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb "ibercs/proto/teams"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Team_Handlers struct {
	teams_client pb.TeamServiceClient
}

func NewTeamsHandlers(client pb.TeamServiceClient) *Team_Handlers {
	return &Team_Handlers{
		teams_client: client,
	}
}

func (h *Team_Handlers) GetAll(c *gin.Context) {
	res, err := h.teams_client.GetAllTeams(c, &pb.Empty{})
	if err != nil {
		logger.Error("Error getting all teams: %v", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting all teams"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}

func (h *Team_Handlers) GetActiveTeams(c *gin.Context) {
	res, err := h.teams_client.GetActiveTeams(c, &pb.Empty{})
	if err != nil {
		logger.Error("Error getting all teams: %v", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting all teams"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}

func (h *Team_Handlers) Get(c *gin.Context) {
	idStr := c.Query("id")
	nickname := c.Query("nickname")

	if idStr == "" && nickname == "" {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			logger.Error("Error converting id to int: %v", err)
			c.JSON(http.StatusBadRequest, response.BuildError("Invalid id"))
			return
		}

		res, err := h.teams_client.GetById(c, &pb.GetTeamByIdRequest{Id: int32(id)})
		if err != nil {
			logger.Error("Error getting team by id: %v", err)
			c.JSON(http.StatusInternalServerError, response.BuildError("Error getting team by id"))
			return
		}

		c.JSON(http.StatusOK, response.BuildOk("ok", res))
		return
	}

	res, err := h.teams_client.GetByNickname(c, &pb.GetTeamByNicknameRequest{Nickname: nickname})
	if err != nil {
		logger.Error("Error getting team by nickname: %v", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error getting team by nickname"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}

func (h *Team_Handlers) CreateFromFaceit(c *gin.Context) {
	var req requests.CreateTeamFromFaceitRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Error binding request: %v", err)
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	pbReq, err := req.ToProto()
	if err != nil {
		logger.Error("Error converting request to proto: %v", err)
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	res, err := h.teams_client.CreateFromFaceit(c, pbReq)
	if err != nil {
		logger.Error("Error creating team from faceit: %v", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error creating team from faceit"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}

func (h *Team_Handlers) FindTeamByPlayerId(c *gin.Context) {
	playerId := c.Param("player_id")
	if playerId == "" {
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	res, err := h.teams_client.FindTeamsByPlayerId(c, &pb.GetTeamByPlayerIdRequest{PlayerId: playerId})
	if err != nil {
		logger.Error("Error finding team by player id: %v", err)
		c.JSON(http.StatusInternalServerError, response.BuildError("Error finding team by player id"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}
