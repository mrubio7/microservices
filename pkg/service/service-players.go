package service

import (
	"errors"
	"fmt"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Players struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewPlayersService(database *gorm.DB) *Players {
	return &Players{
		db: database,
	}
}

func (svc *Players) UpdatePlayer(player model.PlayerModel) error {
	var existingPlayer model.PlayerModel

	svc.mutex.Lock()
	defer svc.mutex.Unlock()

	// Busca al jugador existente por FaceitId
	err := svc.db.Preload("Stats").First(&existingPlayer, "faceit_id = ?", player.FaceitId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Inserta tanto el jugador como las estadísticas si no existe
			return svc.db.Create(&player).Error
		}
		return err
	}

	// Actualiza el jugador si hay cambios en el Nickname, SteamId o Avatar
	if existingPlayer.Nickname != player.Nickname || existingPlayer.SteamId != player.SteamId || existingPlayer.Avatar != player.Avatar {
		if err := svc.db.Model(&existingPlayer).Updates(player).Error; err != nil {
			return err
		}
	}

	// Verifica si las estadísticas necesitan ser actualizadas o insertadas
	if existingPlayer.Stats.ID == 0 {
		// Inserta las estadísticas si no existen
		player.Stats.ID = existingPlayer.ID // Asegura que la clave foránea sea la correcta
		if err := svc.db.Create(&player.Stats).Error; err != nil {
			return err
		}
	} else {
		// Actualiza las estadísticas si ya existen y han cambiado
		if err := svc.db.Model(&existingPlayer.Stats).Updates(player.Stats).Error; err != nil {
			return err
		}
	}

	return nil
}

func (svc *Players) GetPlayers() []model.PlayerModel {
	var players []model.PlayerModel

	err := svc.db.Preload("Stats").Find(&players).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		fmt.Println(err)
		return nil
	}

	return players
}

func (svc *Players) GetNewProminentPlayers() *model.ProminentWeekModel {
	// Definir la consulta SQL para obtener los jugadores prominentes
	query := `
		SELECT ps.id, p.nickname, p.faceit_id, p.steam_id, p.avatar, 
			((ps.kills_average - ps.deaths_average + (ps.assist_average * 0.3)) * ps.kr_ratio * ps.mvp_average) as Score
		FROM player_stats_models ps
		INNER JOIN player_models p ON p.id = ps.id
		ORDER BY Score DESC
		LIMIT 5;
	`

	var results []model.PlayerProminentModel

	err := svc.db.Raw(query).Scan(&results).Error
	if err != nil {
		logger.Error("Error fetching prominent players:", err)
		return nil
	}

	year, week := time.Now().ISOWeek()
	prominentWeek := model.ProminentWeekModel{
		Week:    int16(week),
		Year:    int16(year),
		Players: results,
	}

	err = svc.db.Create(&prominentWeek).Error
	if err != nil {
		logger.Error("Error saving new prominent week:", err)
		return nil
	}

	return &prominentWeek
}

func (svc *Players) GetProminentPlayers() *model.ProminentWeekModel {
	var week model.ProminentWeekModel

	err := svc.db.Preload("Players").Order("year DESC, week DESC").First(&week).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logger.Error("No prominent week found.")
			return nil
		}
		logger.Error("Error fetching prominent week:", err)
		return nil
	}

	return &week
}
