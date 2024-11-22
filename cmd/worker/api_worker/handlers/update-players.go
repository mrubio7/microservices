package handlers

import (
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/managers"
	"ibercs/pkg/response"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

func UpdatePlayers(c *gin.Context) {
	cfg, err := config.LoadWorker()
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Unable to load config"))
		return
	}

	faceitClient := faceit.New(cfg.ThirdPartyApiTokens.FaceitApiToken)

	stateDatabase := database.NewDatabase(cfg.StateDb)
	stateManager := managers.NewStateManager(stateDatabase.GetDB())

	playersDatabase := database.NewDatabase(cfg.PlayersDb)
	playerManager := managers.NewPlayerManager(playersDatabase.GetDB())

	err = worker(playerManager, faceitClient)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error while update players"))
		return
	}

	if err := stateManager.Update_PlayersLastUpdate(); err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, response.BuildError("Error updating PlayersLastUpdate state"))
		return
	}

	c.JSON(http.StatusOK, response.BuildOk("Players updated", nil))
}

func worker(playerManager *managers.PlayerManager, faceitClient *faceit.FaceitClient) error {
	players, err := playerManager.GetAll()
	if err != nil {
		return err
	}

	semaphore := make(chan struct{}, 5)
	errorChan := make(chan error, len(players))
	defer close(errorChan)

	totalPlayers := len(players)
	logger.Info("Starting worker to process %d players...", totalPlayers)
	processPlayers(players, playerManager, faceitClient, semaphore, errorChan)

	return handleWorkerErrors(errorChan)
}

func processPlayers(players []model.PlayerModel, playerManager *managers.PlayerManager, faceitClient *faceit.FaceitClient, semaphore chan struct{}, errorChan chan error) {
	var wg sync.WaitGroup
	var processedCount int32

	for i, player := range players {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(p model.PlayerModel, index int) {
			defer wg.Done()
			defer func() { <-semaphore }()
			updatePlayer(p, playerManager, faceitClient, errorChan, &processedCount, len(players))
		}(player, i)
	}

	wg.Wait()
}

func updatePlayer(player model.PlayerModel, playerManager *managers.PlayerManager, faceitClient *faceit.FaceitClient, errorChan chan error, processedCount *int32, totalPlayers int) {
	playerRefreshed := faceitClient.GetPlayerAverageDetails(player.FaceitId, consts.LAST_MATCHES_NUMBER)
	if playerRefreshed == nil {
		logger.Warning("Player %s cannot be refreshed", player.Nickname)
		return
	}

	playerRefreshed.Id = player.Id
	playerRefreshed.Stats.Id = player.Id

	if err := playerManager.Update(playerRefreshed); err != nil {
		logger.Error("Player %s cannot be updated: %s", playerRefreshed.Nickname, err.Error())
		errorChan <- fmt.Errorf("player %s update error: %w", playerRefreshed.Nickname, err)
	}

	processed := atomic.AddInt32(processedCount, 1)
	logger.Info("Processed player (%d/%d) -> Nickname: %s", processed, totalPlayers, player.Nickname)
}

func handleWorkerErrors(errorChan chan error) error {
	var allErrors []error

	for err := range errorChan {
		allErrors = append(allErrors, err)
	}

	if len(allErrors) > 0 {
		for _, err := range allErrors {
			logger.Error("Worker error: %s", err.Error())
		}
		return fmt.Errorf("worker encountered %d errors", len(allErrors))
	}

	logger.Info("Worker completed successfully without errors.")
	return nil
}
