package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_matches "ibercs/proto/matches"
	"net/http"
	"strconv"

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
	// team_id param is optional
	teamId := c.Query("team_id")

	if teamId != "" {
		res, err := h.matches_client.GetMatchesByTeamId(c, &pb_matches.GetMatchRequest{FaceitId: teamId})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting matches by team id"))
			return
		}
		c.JSON(http.StatusOK, response.BuildOk("Ok", res))
		return
	}

	res, err := h.matches_client.GetAllMatches(c, nil)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting all matches"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Matches_Handlers) Get(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		logger.Error("tried to get an empty id")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID"))
		return
	}

	res, err := h.matches_client.GetMatchByFaceitId(c, &pb_matches.GetMatchRequest{FaceitId: id})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting match"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Matches_Handlers) GetRange(c *gin.Context) {
	days := c.Query("days")

	if days == "" {
		logger.Error("tried to get an empty id")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID"))
		return
	}

	q, err := strconv.Atoi(days)
	if err != nil {
		logger.Error("tried to parse a invalid number %s", days)
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid params"))
		return
	}

	res, err := h.matches_client.GetUpcomingMatches(c, &pb_matches.GetUpcomingRequest{Days: int32(q)})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting matches in the range days"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (h *Matches_Handlers) SetStreamMatch(c *gin.Context) {
	var payload struct {
		FaceitId      string `json:"faceit_id"`
		StreamChannel string `json:"stream_channel"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		logger.Error("body error")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid body"))
		return
	}

	res, err := h.matches_client.SetStreamToMatch(c, &pb_matches.SetStreamRequest{FaceitId: payload.FaceitId, StreamChannel: payload.StreamChannel})
	if err != nil {
		logger.Error("Error saving stream")
		c.JSON(http.StatusInternalServerError, response.BuildError("Internal error"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}
