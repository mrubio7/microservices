package teams

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	"net/http"
	"time"
)

func Update(w http.ResponseWriter) {
	logger.Info("Initializing players worker [UpdatePlayers]")

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Error loading cfg: %s", err)
		return
	}
	db := database.New(cfg.Database)
	psql, _ := db.DB()
	defer psql.Close()

	svcState := service.NewStateService(db)
	svcTeams := service.NewTeamsService(db)
	client := faceit.New(cfg.FaceitApiToken)

	teams := svcTeams.GetAll(true)
	svcState.ClearLastUpdateTeams()
	for _, t := range teams {
		teamData := client.GetTeamById(t.FaceitId)
		teamData.Tournaments = t.Tournaments
		teamData.Twitter = t.Twitter
		teamData.Web = t.Web

		teamUpdated := svcTeams.UpdateTeam(*teamData)
		if teamUpdated == nil {
			logger.Warning("team %s cannot be updated", teamData.Name)
		}
	}
	svcState.SetLastUpdateTeams(time.Now())
}
