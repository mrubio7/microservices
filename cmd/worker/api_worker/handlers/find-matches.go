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
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	eseaManager := managers.NewEseaManager(tournamentsDatabase.GetDB())

	teamsDatabase := database.NewDatabase(cfg.TeamsDb)
	teamManager := managers.NewTeamManager(teamsDatabase.GetDB())

	matchesDatabase := database.NewDatabase(cfg.MatchesDb)
	matchManager := managers.NewMatchManager(matchesDatabase.GetDB())

	dbtournament, _ := tournamentsDatabase.GetDB().DB()
	dbteams, _ := teamsDatabase.GetDB().DB()
	dbmatches, _ := matchesDatabase.GetDB().DB()
	dbstates, _ := stateDatabase.GetDB().DB()
	defer dbtournament.Close()
	defer dbteams.Close()
	defer dbmatches.Close()
	defer dbstates.Close()

	matchesNumber, err := workerFindMatches(tournamentManager, eseaManager, matchManager, teamManager, faceitClient)
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

func workerFindMatches(tournamentManager *managers.TournamentManager, eseaManager *managers.EseaManager, matchManager *managers.MatchManager, teamsManager *managers.TeamManager, faceitClient *faceit.FaceitClient) (int, error) {
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

				matchExist, err := matchManager.GetMatchByFaceitId(match.FaceitId)
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						_, err := matchManager.Create(&match)
						if err != nil {
							logger.Info("Unable to create match %s: %s", match.FaceitId, err.Error())
							continue
						}
					}
					logger.Info("Unable to get match %s: %s", match.FaceitId, err.Error())
					continue
				}

				match.ID = matchExist.ID // For compare
				if !reflect.DeepEqual(match, *matchExist) {
					err := matchManager.Update(&match)
					if err != nil {
						logger.Error("Unable to update match %s: %s", match.FaceitId, err.Error())
						continue
					}
				}

				newMatches += 1
			}
		}
	}

	eseaLeague, err := eseaManager.GetEseaLeagueLive()
	if err != nil {
		return 0, err
	}

	for _, division := range eseaLeague.Divisions {
		matches := faceitClient.GetMatchesFromTournamentId(division.TournamentId, fmt.Sprintf("%s %s", division.DivisionName, division.StageName))

		for _, match := range matches {
			if teamIds[match.TeamAFaceitId] || teamIds[match.TeamBFaceitId] {
				match.IsTeamAKnown = teamIds[match.TeamAFaceitId]
				match.IsTeamBKnown = teamIds[match.TeamBFaceitId]

				matchExist, err := matchManager.GetMatchByFaceitId(match.FaceitId)
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						_, err := matchManager.Create(&match)
						if err != nil {
							logger.Error("Unable to create match %s: %s", match.FaceitId, err.Error())
							continue
						}
					}
					logger.Error("Unable to get match %s: %s", match.FaceitId, err.Error())
					continue
				}

				match.ID = matchExist.ID // For compare
				if !reflect.DeepEqual(match, *matchExist) {
					err := matchManager.Update(&match)
					if err != nil {
						logger.Error("Unable to update match %s: %s", match.FaceitId, err.Error())
						continue
					}
				}

				newMatches += 1
			}
		}
	}

	return newMatches, nil
}
