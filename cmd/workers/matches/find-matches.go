package matches

import (
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
)

func Find() {
	logger.Info("Initializing matches worker [FindMatches]")

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Error loading cfg: %s", err)
		return
	}
	db := database.New(cfg.Database)
	psql, _ := db.DB()
	defer psql.Close()

	//svcState := service.NewStateService(db)
	svcTeams := service.NewTeamsService(db)
	svcTournaments := service.NewTournamentsService(db)
	svcMatches := service.NewMatchesService(db)
	faceitClient := faceit.New(cfg.FaceitApiToken)

	tournaments := svcTournaments.GetAllTournaments()
	if tournaments == nil {
		logger.Error("tournaments empty")
		return
	}

	teams := svcTeams.GetAll(true)

	for _, t := range tournaments {
		var matches []model.MatchModel

		if t.Type == "ESEA" {
			divs := svcTournaments.GetEseaDivisions(t.FaceitId)
			for _, d := range divs {
				matches = append(matches, faceitClient.GetMatchesFromTournamentId(d.FaceitId)...)
			}
		} else {
			matches = faceitClient.GetMatchesFromTournamentId(t.FaceitId)
		}

		teamIds := make(map[string]bool)
		for _, team := range teams {
			teamIds[team.FaceitId] = true
		}

		var filteredMatches []model.MatchModel
		for _, match := range matches {
			if teamIds[match.TeamAFaceitId] || teamIds[match.TeamBFaceitId] {
				filteredMatches = append(filteredMatches, match)
			}
			if teamIds[match.TeamAFaceitId] {
				match.IsTeamAKnown = true
			}
			if teamIds[match.TeamBFaceitId] {
				match.IsTeamBKnown = true
			}
		}

		for _, m := range filteredMatches {
			mm := svcMatches.SaveMatch(m)
			if mm == nil {
				logger.Warning("cannot save match %s", m.FaceitId)
			}
		}
	}

}
