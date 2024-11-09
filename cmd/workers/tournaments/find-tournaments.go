package tournaments

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	"slices"
)

func Find() {
	logger.Info("Initializing worker [FindTournaments]")

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Error loading cfg: %s", err)
		return
	}
	db := database.New(cfg.Database)
	sqldb, _ := db.DB()
	defer sqldb.Close()

	faceit := faceit.New(cfg.FaceitApiToken)
	svcTournaments := service.NewTournamentsService(db)
	svcTeams := service.NewTeamsService(db)
	organizers := svcTournaments.GetAllOrganizers()

	teams := svcTeams.GetAll(true)

	teamsId := make(map[string]bool)
	for _, team := range teams {
		teamsId[team.FaceitId] = true
	}

	for _, org := range organizers {
		switch org.Type {
		case "ESEA":
			save_ESEA_Tournament(org.Type, faceit, svcTournaments, teamsId)

		case "ORGANIZER":
			save_ORGANIZER_Tournament(org.FaceitId, org.Type, faceit, svcTournaments, teamsId)
		}
	}

	logger.Info("Worker [FindTournaments] finished")
}

func save_ESEA_Tournament(orgType string, faceit *faceit.FaceitClient, svcTournaments *service.Tournaments, teams map[string]bool) {
	tournaments := faceit.GetESEASeasons_PRODUCTION()
	for _, t := range tournaments {
		t.Type = orgType

		err := svcTournaments.UpdateTournament(&t)
		if err != nil {
			logger.Error("Unable to create tournament: %s", t.Name)
		}

		eseaDivisions := faceit.GetESEADivisionBySeasonId_PRODUCTION(t.FaceitId, t.Name)
		for _, division := range eseaDivisions {
			tournamentTeams := faceit.GetTeamsInTournament(division.FaceitId)
			for _, team := range tournamentTeams {
				if _, ok := teams[team.FaceitId]; ok {
					division.TeamsId = append(division.TeamsId, team.FaceitId)
				}
			}

			div := svcTournaments.UpdateEseaDivision(division)
			if div == nil {
				logger.Warning("cannot save esea division %s", division.Name)
			}
		}
	}
}

func save_ORGANIZER_Tournament(organizerFaceitID, orgType string, faceit *faceit.FaceitClient, svcTournaments *service.Tournaments, teams map[string]bool) {
	tournaments := faceit.GetAllChampionshipFromOrganizer(organizerFaceitID, 0, 40)
	for _, t := range tournaments {
		var countries []string = t.GeoCountries

		if !slices.Contains(countries, "ES") {
			continue
		}

		if t.JoinPolicy != "public" {
			continue
		}

		tournamentTeams := faceit.GetTeamsInTournament(t.FaceitId)
		for _, team := range tournamentTeams {
			if _, ok := teams[team.FaceitId]; ok {
				t.TeamsId = append(t.TeamsId, team.FaceitId)
			}
		}

		t.Type = orgType
		err := svcTournaments.UpdateTournament(&t)
		if err != nil {
			logger.Error("Unable to create tournament: %s", t.Name)
		}
	}
}
