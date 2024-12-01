package handlers

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/managers"
)

func buildTeamsMap(cfgTeamDb config.DatabaseConfig) map[string]bool {
	teamDatabase := database.NewDatabase(cfgTeamDb)
	teamManager := managers.NewTeamManager(teamDatabase.GetDB())
	dbteams, _ := teamDatabase.GetDB().DB()
	defer dbteams.Close()

	teams, err := teamManager.GetAll()
	if err != nil {
		return nil
	}

	teamsId := make(map[string]bool)
	for _, team := range teams {
		teamsId[team.FaceitId] = true
	}

	return teamsId
}
