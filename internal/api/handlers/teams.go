package handlers

import (
	"ibercs/pkg/logger"
	"ibercs/pkg/response"
	pb_teams "ibercs/proto/teams"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Team_Handlers struct {
	teams_client pb_teams.TeamServiceClient
}

func NewTeamsHandlers(teamsClient pb_teams.TeamServiceClient) *Team_Handlers {
	return &Team_Handlers{
		teams_client: teamsClient,
	}
}

func (th *Team_Handlers) GetAll(c *gin.Context) {
	var res *pb_teams.TeamList
	var err error

	active := c.Query("active")
	if active == "" {
		res, err = th.teams_client.GetTeams(c, &pb_teams.GetTeamsRequest{Active: false})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting all teams"))
			return
		}
	} else {
		res, err = th.teams_client.GetTeams(c, &pb_teams.GetTeamsRequest{Active: true})
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, response.BuildError("Error getting all teams"))
			return
		}
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (th *Team_Handlers) Get(c *gin.Context) {
	teamId := c.Query("id")
	if teamId == "" {
		logger.Error("tried to get an empty id")
		c.JSON(http.StatusBadRequest, response.BuildError("Invalid ID"))
		return
	}

	res, err := th.teams_client.GetTeam(c, &pb_teams.NewTeamRequest{FaceitId: teamId})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting team"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}

func (th *Team_Handlers) New(c *gin.Context) {
	var input struct {
		ID string `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, response.BuildError("Error creating the team"))
		return
	}

	res, err := th.teams_client.NewTeam(c, &pb_teams.NewTeamRequest{FaceitId: input.ID})
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, response.BuildError("Error getting team"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", res))
}
