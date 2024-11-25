package client

import (
	"fmt"
	"github.com/mconnat/go-faceit/pkg/models"
)

func (c *FaceITClient) GetTeamByID(teamID string, params map[string]interface{}) (models.Team, error) {
	team, err := Get(models.Team{}, c, fmt.Sprintf("/teams/%s", teamID), params)
	if err != nil {
		return models.Team{}, err
	}
	return team, nil
}

func (c *FaceITClient) GetTeamStats(teamID string, gameID string, params map[string]interface{}) (models.TeamStatsResponse, error) {
	stats, err := Get(models.TeamStatsResponse{}, c, fmt.Sprintf("/teams/%s/stats/%s", teamID, gameID), params)
	if err != nil {
		return models.TeamStatsResponse{}, err
	}
	return stats, nil
}

func (c *FaceITClient) GetTeamTournaments(teamID string, params map[string]interface{}) (models.Tournaments, error) {
	tournaments, err := Get(models.Tournaments{}, c, fmt.Sprintf("/teams/%s/tournaments", teamID), params)
	if err != nil {
		return models.Tournaments{}, err
	}
	return tournaments, nil
}
