package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_matches "ibercs/proto/matches"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Matches_Handlers struct {
	matches_client pb_matches.MatchesServiceClient
}

func NewMatchesHandlers(client pb_matches.MatchesServiceClient) *Matches_Handlers {
	return &Matches_Handlers{
		matches_client: client,
	}
}

func (h *Matches_Handlers) GetAll(c *gin.Context) {
	res, err := h.matches_client.GetAllMatches(c, nil)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting all matches"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Matches_Handlers) GetById(c *gin.Context) {
	matchId := c.Query("id")

	if matchId == "" {
		logger.Error("tried to get an empty id")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID"))
		return
	}

	res, err := h.matches_client.GetMatchByFaceitId(c, &pb_matches.GetMatchRequest{FaceitId: matchId})
	if err != nil {
		logger.Error("tried to get an empty id")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}
