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
		if org.Type == "ESEA" {
			continue
		}

		tournaments := faceit.GetAllChampionshipFromOrganizer(org.FaceitId, 0, 40)
		for _, t := range tournaments {
			var countries []string = t.GeoCountries

			if !slices.Contains(countries, "ES") {
				continue
			}

			if t.JoinPolicy != "public" {
				continue
			}

			t.Type = org.Type
			name := t.Name
			t := svcTournaments.NewTournament(&t)
			if t == nil {
				logger.Error("Unable to create tournament: %s", name)
			}
		}
	}
}
