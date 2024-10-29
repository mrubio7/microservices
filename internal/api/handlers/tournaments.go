package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_tournaments "ibercs/proto/tournaments"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Tournament_Handlers struct {
	tournaments_client pb_tournaments.TournamentServiceClient
}

func NewTournamentsHandlers(tournamentsClient pb_tournaments.TournamentServiceClient) *Tournament_Handlers {
	return &Tournament_Handlers{
		tournaments_client: tournamentsClient,
	}
}

func (h *Tournament_Handlers) NewOrganizer(c *gin.Context) {
	var req struct {
		FaceitId string `json:"faceit_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error parsing request"))
		return
	}

	res, err := h.tournaments_client.NewOrganizer(c, &pb_tournaments.NewOrganizerRequest{FaceitId: req.FaceitId})
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.AlreadyExists {
			c.JSON(http.StatusOK, response.BuildOk("Already exist", nil))
			return
		}
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error creating organizer"))
		return
	}

	c.JSON(http.StatusCreated, response.BuildOk("Ok", res))
}

func (h *Tournament_Handlers) NewTournament(c *gin.Context) {
	var req struct {
		FaceitId string `json:"faceit_id"`
		Type     string `json:"type"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error parsing request"))
		return
	}

	res, err := h.tournaments_client.NewTournament(c, &pb_tournaments.NewTournamentRequest{FaceitId: req.FaceitId, Type: req.Type})
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.AlreadyExists {
			c.JSON(http.StatusOK, response.BuildOk("Already exist", nil))
			return
		}
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error creating tournament"))
		return
	}

	c.JSON(http.StatusCreated, response.BuildOk("Ok", res))
}

func (h *Tournament_Handlers) GetAllTournaments(c *gin.Context) {
	res, err := h.tournaments_client.GetAllTorunaments(c, &pb_tournaments.Empty{})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error creating tournament"))
		return
	}

	c.JSON(http.StatusCreated, response.BuildOk("Ok", res))
}
