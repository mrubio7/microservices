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

	organizers := svcTournaments.GetAllOrganizers()

	for _, org := range organizers {
		switch org.Type {
		case "ESEA":
			save_ESEA_Tournament(org.Type, faceit, svcTournaments)

		case "ORGANIZER":
			save_ORGANIZER_Tournament(org.FaceitId, org.Type, faceit, svcTournaments)
		}

	}
}

func save_ESEA_Tournament(orgType string, faceit *faceit.FaceitClient, svcTournaments *service.Tournaments) {
	tournaments := faceit.GetESEASeasons_PRODUCTION()
	for _, t := range tournaments {
		t.Type = orgType
		err := svcTournaments.UpdateTournament(&t)
		if err != nil {
			logger.Error("Unable to create tournament: %s", t.Name)
		}

		eseaDivisions := faceit.GetESEADivisionBySeasonId_PRODUCTION(t.FaceitId, t.Name)
		for _, division := range eseaDivisions {
			div := svcTournaments.NewEseaDivision(division)
			if div == nil {
				logger.Warning("cannot save esea division %s", division.Name)
			}
		}
	}
}

func save_ORGANIZER_Tournament(organizerFaceitID, orgType string, faceit *faceit.FaceitClient, svcTournaments *service.Tournaments) {
	tournaments := faceit.GetAllChampionshipFromOrganizer(organizerFaceitID, 0, 40)
	for _, t := range tournaments {
		var countries []string = t.GeoCountries

		if !slices.Contains(countries, "ES") {
			continue
		}

		if t.JoinPolicy != "public" {
			continue
		}

		t.Type = orgType
		err := svcTournaments.UpdateTournament(&t)
		if err != nil {
			logger.Error("Unable to create tournament: %s", t.Name)
		}
	}
}
