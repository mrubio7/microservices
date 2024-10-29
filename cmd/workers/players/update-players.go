package players

import (
	"encoding/json"
	"fmt"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	"net/http"
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

	for i, p := range playersList {
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
			Current: i + 1,
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
	}

	svcState.SetLastUpdatePlayer(time.Now())

	// Notificamos que el proceso ha terminado
	fmt.Fprintf(w, "data: %s\n\n", `{"message": "Update completed"}`)
	w.(http.Flusher).Flush()
}
