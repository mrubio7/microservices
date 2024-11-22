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

	tournamentsNumber, err := workerFindTournaments(tournamentManager, faceitClient)
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

func workerFindTournaments(tournamentManager *managers.TournamentManager, faceitClient *faceit.FaceitClient) (int, error) {
	var newTournaments int

	organizers, err := tournamentManager.GetAllOrganizers()
	if err != nil {
		return 0, err
	}

	for _, organizer := range organizers {
		tournaments := faceitClient.GetAllChampionshipFromOrganizer(organizer.FaceitId)

		for _, tournament := range tournaments {
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
