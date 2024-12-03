package handlers

import (
	"fmt"
	"ibercs/internal/model"
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

func UpdateNearbyMatches(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	faceitClient := faceit.New(cfg.ThirdPartyApiTokens.FaceitApiToken)

	matchesDatabase := database.NewDatabase(cfg.MatchesDb)
	matchManager := managers.NewMatchManager(matchesDatabase.GetDB())

	dbmatches, _ := matchesDatabase.GetDB().DB()
	defer dbmatches.Close()

	matchesNumber, err := workerUpdateNearMatches(matchManager, faceitClient)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, response.BuildOk("There are no matches to update", nil))
			return
		}
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while find matches"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk(fmt.Sprintf("%d matches updated", matchesNumber), nil))
}

func workerUpdateNearMatches(matchManager *managers.MatchManager, faceitClient *faceit.FaceitClient) (int, error) {
	var matchesUpdated int

	matches := getYesterdayAndTodayMatches(matchManager)
	if len(matches) == 0 {
		return 0, gorm.ErrRecordNotFound
	}

	for _, match := range matches {
		matchDetails := faceitClient.GetMatchDetails(match.FaceitId)
		if matchDetails == nil {
			logger.Warning("Unable to get details of match %s", match.FaceitId)
			continue
		}

		if !reflect.DeepEqual(*matchDetails, match) {
			matchDetails.ID = match.ID
			err := matchManager.Update(matchDetails)
			if err != nil {
				logger.Warning("Unable to update match %s", match.FaceitId)
				continue
			}
			matchesUpdated += 1
		}
	}

	return matchesUpdated, nil
}

func getYesterdayAndTodayMatches(matchManager *managers.MatchManager) map[string]model.MatchModel {
	yesterdayMatches, _ := matchManager.GetYesterdayMatches()
	todayMatches, _ := matchManager.GetUpcomingMatches()

	matches := make(map[string]model.MatchModel)

	for _, m := range yesterdayMatches {
		matches[m.FaceitId] = m
	}
	for _, m := range todayMatches {
		matches[m.FaceitId] = m
	}

	return matches
}
