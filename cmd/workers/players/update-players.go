package players

import (
	"encoding/json"
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	"net/http"
	"sync"
	"time"
)

func Update(w http.ResponseWriter) {
	logger.Info("Initializing players worker [UpdatePlayers]")

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Error loading cfg: %s", err)
		return
	}
	db := database.New(cfg.Database)
	psql, _ := db.DB()
	defer psql.Close()

	// Configura el límite de conexiones a la base de datos para optimizar la concurrencia
	psql.SetMaxOpenConns(10)
	psql.SetMaxIdleConns(10)
	psql.SetConnMaxLifetime(time.Hour)

	svcState := service.NewStateService(db)
	svcPlayers := service.NewPlayersService(db)
	client := faceit.New(cfg.FaceitApiToken)

	state := svcState.GetState()

	if !state.LastPlayerUpdate.Time.Before(time.Now().Add(-1 * time.Hour)) {
		fmt.Fprintf(w, "data: %s\n\n", `{"error": "Less than 1h"}`)
		w.(http.Flusher).Flush()
		return
	}

	svcState.ClearLastUpdatePlayer()

	playersList := svcPlayers.GetPlayers()
	totalPlayers := len(playersList)

	// Canal y WaitGroup para manejar la concurrencia
	playerChan := make(chan model.PlayerModel, 5)
	var wg sync.WaitGroup
	var curr = 1

	// Lanza un grupo de workers limitados para hacer la actualización en paralelo
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range playerChan {
				player := client.GetPlayerAverageDetails(p.FaceitId, consts.LAST_MATCHES_NUMBER)
				if player == nil {
					logger.Warning("User %s doesn't update\n", p.FaceitId)
					continue
				}

				err := svcPlayers.UpdatePlayer(*player)
				if err != nil {
					logger.Error(err.Error())
				}

				// Actualizamos el progreso en el SSE
				progressData := struct {
					Current int    `json:"current"`
					Total   int    `json:"total"`
					Player  string `json:"player"`
				}{
					Current: curr,
					Total:   totalPlayers,
					Player:  p.Nickname,
				}

				data, err := json.Marshal(progressData)
				if err != nil {
					logger.Error("Error marshalling progress data: %s", err)
					continue
				}

				fmt.Fprintf(w, "data: %s\n\n", data)
				w.(http.Flusher).Flush()
				curr++
			}
		}()
	}

	// Envía los jugadores al canal para procesarlos en paralelo
	for _, p := range playersList {
		playerChan <- p
	}
	close(playerChan)

	// Espera a que todas las goroutines finalicen
	wg.Wait()

	svcState.SetLastUpdatePlayer(time.Now().UTC())
	logger.Info("Players update worker [UpdatePlayers] finished")
}
