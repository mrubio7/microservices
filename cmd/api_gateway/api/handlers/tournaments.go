package handlers

import (
	"ibercs/cmd/api_gateway/api/requests"
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb "ibercs/proto/tournaments"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tournament_Handlers struct {
	tournament_client pb.TournamentServiceClient
}

func NewTournamentsHandlers(client pb.TournamentServiceClient) *Tournament_Handlers {
	return &Tournament_Handlers{
		tournament_client: client,
	}
}

func (h *Tournament_Handlers) CreateOrganizer(c *gin.Context) {
	var req requests.CreateOrganizerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	payload, err := req.ToProto()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid request"))
		return
	}

	res, err := h.tournament_client.CreateOrganizer(c, payload)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error creating organizer"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}

func (h *Tournament_Handlers) GetAll(c *gin.Context) {
	res, err := h.tournament_client.GetAllTournaments(c, nil)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting tournaments"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}

// Esea
func (h *Tournament_Handlers) GetEseaLeagues(c *gin.Context) {
	seasonStr := c.Query("season")
	if seasonStr == "" {
		logger.Error("tried to get an empty id")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID"))
		return
	}

	var res *pb.Esea
	var err error
	if seasonStr == "current" {
		res, err = h.tournament_client.GetLiveEseaDetails(c, nil)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting tournament"))
			return
		}
	} else {
		season, err := strconv.Atoi(seasonStr)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Invalid season"))
			return
		}

		res, err = h.tournament_client.GetEseaDetailsBySeason(c, &pb.GetEseaLeagueBySeasonNumberRequest{Season: int32(season)})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting tournament"))
			return
		}
	}

	c.JSON(http.StatusOK, response.BuildOk("ok", res))
}
