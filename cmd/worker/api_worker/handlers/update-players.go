package handlers

import (
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePlayers(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	playersDatabase := database.NewDatabase(cfg.PlayersDb)
	playerManager := managers.NewPlayerManager(playersDatabase.GetDB())

	err = worker(playerManager)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while update players"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Players updated", nil))
}

func worker(playerManager *managers.PlayerManager) error {
	return nil
}
