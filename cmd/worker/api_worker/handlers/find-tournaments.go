package handlers

import (
	"fmt"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func FindTournaments(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	faceitClient := faceit.New(cfg.ThirdPartyApiTokens.FaceitApiToken)

	stateDatabase := database.NewDatabase(cfg.StateDb)
	stateManager := managers.NewStateManager(stateDatabase.GetDB())

	tournamentsDatabase := database.NewDatabase(cfg.TournamentsDb)
	tournamentManager := managers.NewTournamentManager(tournamentsDatabase.GetDB())

	dbtournament, _ := tournamentsDatabase.GetDB().DB()
	dbstates, _ := stateDatabase.GetDB().DB()
	defer dbtournament.Close()
	defer dbstates.Close()

	teamsMap := buildTeamsMap(cfg.TeamsDb)

	tournamentsNumber, err := workerFindTournaments(tournamentManager, faceitClient, teamsMap)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while find tournaments"))
		return
	}

	if err := stateManager.Update_TournamentLastFind(); err != nil {
		logger.Warning(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("State could not be updated"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk(fmt.Sprintf("%d new tournaments found", tournamentsNumber), nil))
}

func workerFindTournaments(tournamentManager *managers.TournamentManager, faceitClient *faceit.FaceitClient, teamsMap map[string]bool) (int, error) {
	var newTournaments int

	organizers, err := tournamentManager.GetAllOrganizers()
	if err != nil {
		return 0, err
	}

	for _, organizer := range organizers {
		tournaments := faceitClient.GetAllChampionshipFromOrganizer(organizer.FaceitId)

		for _, tournament := range tournaments {
			if !slices.Contains(tournament.GeoCountries, "ES") || tournament.JoinPolicy != "public" {
				continue
			}

			teams := faceitClient.GetTeamsInTournament(tournament.FaceitId)
			for _, team := range teams {
				if teamsMap[team.FaceitId] {
					tournament.TeamsId = append(tournament.TeamsId, team.FaceitId)
				}
			}

			_, err := tournamentManager.CreateTournament(&tournament)
			if err != nil {
				if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
					continue
				}
				logger.Error("Unable to create tournament %s: %s", tournament.Name, err.Error())
			}
			newTournaments += 1
		}
	}

	return newTournaments, nil
}
