package update_players

import (
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	"time"

	"github.com/schollz/progressbar/v3"
)

func Start() {
	logger.Info("Initializing players worker [UpdatePlayers]")

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Error loading cfg: %s", err)
		return
	}
	db := database.New(cfg.Database)
	psql, _ := db.DB()
	defer psql.Close()

	svcState := service.NewStateService(db)
	svcPlayers := service.NewPlayersService(db)
	client := faceit.New(cfg.FaceitApiToken)

	state := svcState.GetState()

	if !state.LastPlayerUpdate.Valid {
		return
	}

	svcState.ClearLastUpdatePlayer()

	playersList := svcPlayers.GetPlayers()

	bar := progressbar.NewOptions(len(playersList),
		progressbar.OptionSetDescription("Processing players"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionClearOnFinish(),
	)
	defer bar.Close()

	for _, p := range playersList {

		player := client.GetPlayerAverageDetails(p.FaceitId, consts.LAST_MATCHES_NUMBER)

		if player == nil {
			logger.Warning("User %s doesnt update\n", p.FaceitId)
		}
		err := svcPlayers.UpdatePlayer(*player)
		if err != nil {
			logger.Error(err.Error())
		}
		bar.Add(1)

	}

	svcState.SetLastUpdatePlayer(time.Now())
}
