package service

import (
	"errors"
	"fmt"
	"ibercs/internal/model"
	"sync"

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

	// Actualiza el jugador si hay cambios
	if existingPlayer.Nickname != player.Nickname || existingPlayer.SteamId != player.SteamId {
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
