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

func FindMatches(c *gin.Context) {
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

	teamsDatabase := database.NewDatabase(cfg.TeamsDb)
	teamManager := managers.NewTeamManager(teamsDatabase.GetDB())

	matchesDatabase := database.NewDatabase(cfg.MatchesDb)
	matchManager := managers.NewMatchManager(matchesDatabase.GetDB())

	matchesNumber, err := workerFindMatches(tournamentManager, matchManager, teamManager, faceitClient)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while find matches"))
		return
	}

	if err := stateManager.Update_MatchLastFind(); err != nil {
		logger.Warning(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("State could not be updated"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk(fmt.Sprintf("%d new matches found", matchesNumber), nil))
}

func workerFindMatches(tournamentManager *managers.TournamentManager, matchManager *managers.MatchManager, teamsManager *managers.TeamManager, faceitClient *faceit.FaceitClient) (int, error) {
	var newMatches int

	tournaments, err := tournamentManager.GetAllTournamentsActive()
	if err != nil {
		return 0, err
	}

	teams, err := teamsManager.GetAll()
	if err != nil {
		return 0, err
	}

	teamIds := make(map[string]bool)
	for _, team := range teams {
		teamIds[team.FaceitId] = true
	}

	for _, tournament := range tournaments {
		matches := faceitClient.GetMatchesFromTournamentId(tournament.FaceitId, tournament.Name)

		for _, match := range matches {
			if teamIds[match.TeamAFaceitId] || teamIds[match.TeamBFaceitId] {
				match.IsTeamAKnown = teamIds[match.TeamAFaceitId]
				match.IsTeamBKnown = teamIds[match.TeamBFaceitId]

				_, err := matchManager.Create(&match)
				if err != nil {
					if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
						continue
					}
					logger.Error("Unable to create tournament %s: %s", tournament.Name, err.Error())
				}
				newMatches += 1
			}
		}
	}

	return newMatches, nil
}
