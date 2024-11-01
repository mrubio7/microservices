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

func (svc *Players) GetPlayer(id string) *model.PlayerModel {
	var player *model.PlayerModel

	err := svc.db.Model(&model.PlayerModel{}).Preload("Stats").First(&player, "faceit_id = ?", id).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return player
}

func (svc *Players) GetPlayerByNickname(nickname string) *model.PlayerModel {
	var player *model.PlayerModel

	err := svc.db.Model(&model.PlayerModel{}).Preload("Stats").First(&player, "nickname = ?", nickname).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return player
}

func (svc *Players) GetPlayers() []model.PlayerModel {
	var players []model.PlayerModel

	err := svc.db.Preload("Stats").Find(&players).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return players
}

func (svc *Players) GetNewProminentPlayers() *model.ProminentWeekModel {
	year, week := time.Now().ISOWeek()

	var existingWeek model.ProminentWeekModel
	err := svc.db.Where("year = ? AND week = ?", int16(year), int16(week)).First(&existingWeek).Error
	if err == nil {
		logger.Info("Prominent week already exists for the current week and year.")
		return &existingWeek
	}

	query := fmt.Sprintf(`
		WITH previous_week AS (
			SELECT ppm.id
			FROM player_prominent_models ppm
			INNER JOIN prominent_week_models pwm ON ppm.prominent_week_id = pwm.id
			WHERE pwm.year = %d AND pwm.week = %d
		)

		SELECT ps.id, p.avatar, p.nickname, 
			((kills_average * 0.35) - (deaths_average * 0.15) + (assist_average * 0.1) + (kr_ratio * 0.2) + (mvp_average * 0.1)) * (1 + (elo / 1000)) AS calc
		FROM player_stats_models ps 
		INNER JOIN player_models p ON p.id = ps.id 
		LEFT JOIN player_prominent_models ppm ON ppm.id = ps.id
		LEFT JOIN prominent_week_models pwm ON pwm.id = ppm.prominent_week_id
		WHERE (pwm.week IS NULL OR pwm.year < %d OR pwm.week < %d)
		AND ps.id NOT IN (SELECT id FROM previous_week)
		ORDER BY calc DESC
		LIMIT 5;
	`, year, week-1, year, week+1)

	var results []model.PlayerProminentModel

	err = svc.db.Raw(query).Scan(&results).Error
	if err != nil {
		logger.Error("Error fetching prominent players:", err)
		return nil
	}

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

func (svc *Players) BatchUpdatePlayers(players []model.PlayerModel) error {
	svc.mutex.Lock()
	defer svc.mutex.Unlock()
	return svc.db.Transaction(func(tx *gorm.DB) error {
		for _, player := range players {
			var existingPlayer model.PlayerModel
			err := tx.Preload("Stats").First(&existingPlayer, "faceit_id = ?", player.FaceitId).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					if err := tx.Create(&player).Error; err != nil {
						return err
					}
					continue
				}
				return err
			}

			// Actualiza los datos del jugador si han cambiado
			if existingPlayer.Nickname != player.Nickname || existingPlayer.SteamId != player.SteamId || existingPlayer.Avatar != player.Avatar {
				if err := tx.Model(&existingPlayer).Updates(player).Error; err != nil {
					return err
				}
			}

			// Inserta o actualiza estadísticas
			if existingPlayer.Stats.ID == 0 {
				player.Stats.ID = existingPlayer.ID
				if err := tx.Create(&player.Stats).Error; err != nil {
					return err
				}
			} else if err := tx.Model(&existingPlayer.Stats).Updates(player.Stats).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
