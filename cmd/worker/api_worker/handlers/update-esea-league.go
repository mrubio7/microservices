package handlers

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while update esea leagues"))
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
		logger.Error("unable to get ESEA divisions")
		return err
	}

	for _, division := range eseaDivisions {
		d, err := eseaManager.GetDivisionByFaceitId(division.ConferenceId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				division.EseaLeagueId = eseaLeague.Id
				divCreated, err := eseaManager.CreateDivision(&division)
				if err != nil {
					return err
				}
				division = *divCreated
			} else {
				return err
			}
		}

		if d != nil {
			division.Id = d.Id
		} else {
			continue
		}

		standings := faceitClient.GetESEADivisionStanding_PRODUCTION(division.ConferenceId)
		if standings == nil {
			logger.Warning("Error getting ESEA %s %s: %s", division.DivisionName, division.StageName, err.Error())
			continue
		}

		for _, standing := range standings {
			if teamsMap[standing.TeamFaceitId] {
				s, err := eseaManager.GetStandingByTeamFaceitIdAndDivisionId(standing.TeamFaceitId, int(division.Id))
				if err != nil {
					logger.Warning("Esea standing for team %s does not exist", standing.TeamFaceitId)
					division.Standings = append(division.Standings, standing)
					continue
				}
				standing.Id = s.Id
				division.Standings = append(division.Standings, standing)
			}
		}

		err = eseaManager.UpdateDivision(division)
		if err != nil {
			logger.Error("unable to update division: %s", err.Error())
			return err
		}
	}

	return nil
}
