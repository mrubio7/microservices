package handlers

import (
	"errors"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateEsea(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	faceitClient := faceit.New(cfg.ThirdPartyApiTokens.FaceitApiToken)

	tournamentsDatabase := database.NewDatabase(cfg.TournamentsDb)
	eseaManager := managers.NewEseaManager(tournamentsDatabase.GetDB())

	teamsMap := buildTeamsMap(cfg.TeamsDb)

	err = workerEseaLeagueUpdate(eseaManager, faceitClient, teamsMap)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while update players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Ok", nil))
}

func workerEseaLeagueUpdate(eseaManager *managers.EseaManager, faceitClient *faceit.FaceitClient, teamsMap map[string]bool) error {
	eseaLeague, err := eseaManager.GetEseaLeagueLive()
	if err != nil {
		logger.Error("unable to get ESEA league: %s", err.Error())
		return err
	}

	eseaDivisions := faceitClient.GetESEADivisionBySeasonId_PRODUCTION(eseaLeague.FaceitId, eseaLeague.Name)
	if eseaDivisions == nil {
		err := errors.New("unable to get ESEA divisions")
		logger.Error(err.Error())
		return err
	}

	for i, division := range eseaDivisions {
		division.Id = eseaLeague.Divisions[i].Id

		standings := faceitClient.GetESEADivisionStanding_PRODUCTION(division.FaceitId)
		if standings == nil {
			err := errors.New("unable to get ESEA standings")
			logger.Error(err.Error())
			return err
		}

		for _, standing := range standings {
			if teamsMap[standing.TeamFaceitId] {
				division.Standings = append(division.Standings, standing)
			}
		}

		err := eseaManager.UpdateDivision(division)
		if err != nil {
			logger.Error("unable to update division: %s", err.Error())
			return err
		}
	}

	return nil
}
