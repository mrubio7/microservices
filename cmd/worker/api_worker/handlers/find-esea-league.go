package handlers

import (
	"errors"
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

func FindEseaLeague(c *gin.Context) {
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

	eseaLeaguesNumber, err := workerFindEseaLeague(eseaManager, faceitClient, teamsMap)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while find esea league"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk(fmt.Sprintf("%d esea leagues found", eseaLeaguesNumber), nil))
}

func workerFindEseaLeague(eseaManager *managers.EseaManager, faceitClient *faceit.FaceitClient, teamsMap map[string]bool) (int, error) {
	var eseaLeaguesNumber int

	eseaLeagues := faceitClient.GetESEASeasons_PRODUCTION()
	if eseaLeagues == nil {
		err := errors.New("unable to get ESEA seasons")
		logger.Error(err.Error())
		return 0, err
	}

	for i := range eseaLeagues {
		eseaDivisions := faceitClient.GetESEADivisionBySeasonId_PRODUCTION(eseaLeagues[i].FaceitId, eseaLeagues[i].Name)
		if eseaDivisions == nil {
			err := errors.New("unable to get ESEA divisions")
			logger.Error(err.Error())
			return 0, err
		}

		if eseaLeagues[i].Status == "live" {
			for j := range eseaDivisions {

				standings := faceitClient.GetESEADivisionStanding_PRODUCTION(eseaDivisions[j].FaceitId)
				if standings == nil {
					err := errors.New("unable to get ESEA standings")
					logger.Error(err.Error())
					return 0, err
				}

				for _, standing := range standings {
					if teamsMap[standing.TeamFaceitId] {
						eseaDivisions[j].Standings = append(eseaDivisions[j].Standings, standing)
					}
				}

				eseaLeagues[i].Divisions = append(eseaLeagues[i].Divisions, eseaDivisions[j])
			}
		}

		_, err := eseaManager.CreateEseaLeague(&eseaLeagues[i])
		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
				continue
			} else {
				logger.Error("unable to create ESEA league: %s", err.Error())
				return 0, err
			}
		}
		eseaLeaguesNumber += 1
	}

	return eseaLeaguesNumber, nil
}
