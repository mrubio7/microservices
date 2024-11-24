package handlers

import (
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateTeams(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	faceitClient := faceit.New(cfg.ThirdPartyApiTokens.FaceitApiToken)

	stateDatabase := database.NewDatabase(cfg.StateDb)
	stateManager := managers.NewStateManager(stateDatabase.GetDB())

	teamDatabase := database.NewDatabase(cfg.TeamsDb)
	teamManager := managers.NewTeamManager(teamDatabase.GetDB())

	tournamentsDatabase := database.NewDatabase(cfg.TournamentsDb)
	tournamentManager := managers.NewTournamentManager(tournamentsDatabase.GetDB())
	eseaManager := managers.NewEseaManager(tournamentsDatabase.GetDB())

	err = workerTeamsUpdate(teamManager, tournamentManager, eseaManager, faceitClient)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while update players"))
		return
	}

	if err := stateManager.Update_TeamLastUpdate(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error updating TeamLastUpdate state"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Teams updated", nil))
}

func workerTeamsUpdate(teamManager *managers.TeamManager, tournamentManager *managers.TournamentManager, eseaManager *managers.EseaManager, faceitClient *faceit.FaceitClient) error {
	teamsActive := make(map[string]bool)

	// Get esea teams
	eseaLeague, err := eseaManager.GetEseaLeagueLive()
	if err != nil {
		logger.Error("Unable to get esea league: %s", err.Error())
		return err
	}

	var standings []model.EseaStandingModel
	for _, division := range eseaLeague.Divisions {
		temp, err := eseaManager.GetStandingsByDivisionId(int(division.Id))
		if err != nil {
			logger.Error("Unable to get standings by division faceit id: %s", err.Error())
			return err
		}

		standings = append(standings, temp...)
	}

	for _, eseaStanding := range standings {
		teamsActive[eseaStanding.TeamFaceitId] = true
	}

	// Get tournaments teams
	tournaments, err := tournamentManager.GetAllTournamentsActive()
	if err != nil {
		logger.Error("Unable to get tournaments: %s", err.Error())
		return err
	}

	for _, tournament := range tournaments {
		for _, tournamentTeam := range tournament.TeamsId {
			teamsActive[tournamentTeam] = true
		}
	}

	err = teamUpdater(teamsActive, teamManager, faceitClient)
	if err != nil {
		logger.Error("Unable to update teams: %s", err.Error())
		return err
	}

	return nil
}

func teamUpdater(teamsActive map[string]bool, teamManager *managers.TeamManager, faceitClient *faceit.FaceitClient) error {
	teams, err := teamManager.GetAll()
	if err != nil {
		return err
	}

	for _, team := range teams {
		if teamsActive[team.FaceitId] {
			teamRefreshed := faceitClient.GetTeamById(team.FaceitId)
			teamRefreshed.Id = team.Id
			teamRefreshed.Active = true
			err := teamManager.Update(teamRefreshed)
			if err != nil {
				logger.Error("unable to update team %s: %s", team.Nickname, err.Error())
			}
		} else {
			err := teamManager.DesactivateTeam(int(team.Id))
			if err != nil {
				logger.Error("unable to desactivate team %s: %s", team.Nickname, err.Error())
			}
		}
	}

	return nil
}
