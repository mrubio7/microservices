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
)

func UpdateTournaments(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	faceitClient := faceit.New(cfg.ThirdPartyApiTokens.FaceitApiToken)

	teamsDatabase := database.NewDatabase(cfg.TeamsDb)
	teamManager := managers.NewTeamManager(teamsDatabase.GetDB())

	tournamentsDatabase := database.NewDatabase(cfg.TournamentsDb)
	tournamentManager := managers.NewTournamentManager(tournamentsDatabase.GetDB())

	tournamentsNumber, err := workerUpdateTournaments(tournamentManager, teamManager, faceitClient)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while find tournaments"))
		return
	}
	c.JSON(http.StatusOK, response.BuildOk(fmt.Sprintf("%d tournaments updated", tournamentsNumber), nil))
}

func workerUpdateTournaments(tournamentManager *managers.TournamentManager, teamsManager *managers.TeamManager, faceitClient *faceit.FaceitClient) (int, error) {
	var tournamentsUpdated int

	teams, err := teamsManager.GetAll()
	if err != nil {
		return 0, err
	}

	teamIds := make(map[string]bool)
	for _, team := range teams {
		teamIds[team.FaceitId] = true
	}

	tournaments, err := tournamentManager.GetAllTournamentsActive()
	if err != nil {
		return 0, err
	}

	for _, tournament := range tournaments {
		t := faceitClient.GetChampionshipById(tournament.FaceitId)
		tournamentTeams := faceitClient.GetTeamsInTournament(tournament.FaceitId)

		t.Id = tournament.Id
		for _, tournamentTeam := range tournamentTeams {
			if teamIds[tournamentTeam.FaceitId] {
				t.TeamsId = append(t.TeamsId, tournamentTeam.FaceitId)
			}
		}

		err := tournamentManager.UpdateTournament(t)
		if err != nil {
			logger.Error("Error updating tournament %s", t.Name)
			return 0, err
		}

		tournamentsUpdated += 1
	}

	return tournamentsUpdated, nil
}
