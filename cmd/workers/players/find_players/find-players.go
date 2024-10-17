package find_players

import (
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	"time"

	"github.com/schollz/progressbar/v3"
)

func Start(size int) {
	logger.Info("Initializing players worker [FindPlayers]")
	startTime := time.Now()

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Error loading cfg: %s", err)
		return
	}
	db := database.New(cfg.Database)
	sqldb, _ := db.DB()
	defer sqldb.Close()

	svcPlayers := service.NewPlayersService(db)
	client := faceit.New(cfg.FaceitApiToken)

	chPlayers := make(chan model.PlayerModel, 200)

	go client.GetAllPlayers(chPlayers, size)

	bar := progressbar.NewOptions(size,
		progressbar.OptionSetDescription("Processing players"),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionClearOnFinish(),
	)
	defer bar.Close()

	for p := range chPlayers {
		err := svcPlayers.UpdatePlayer(p)
		if err != nil {
			logger.Error("Error saving the player %s: %s", p.Nickname, err)
		}

		bar.Add(1)
	}

	logger.Info("All the players were found and registered successfully")
	logger.Info("[Find players] ended in %s", time.Since(startTime).String())
}
