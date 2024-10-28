package service

import (
	"database/sql"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"sync"
	"time"

	"gorm.io/gorm"
)

type State struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewStateService(database *gorm.DB) *State {
	return &State{
		db: database,
	}
}

func (svc *State) GetState() *model.StateModel {
	var state *model.StateModel
	err := svc.db.First(&state).Error
	if err != nil {
		logger.Error(err.Error())
		return nil
	}

	return state
}

func (svc *State) ClearLastUpdatePlayer() error {
	state := svc.GetState()

	state.LastPlayerUpdate = sql.NullTime{Valid: false}

	err := svc.db.Where("id = ?", 1).Save(&state).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func (svc *State) SetLastUpdatePlayer(date time.Time) error {
	state := svc.GetState()

	state.LastPlayerUpdate = sql.NullTime{Valid: true, Time: date}

	err := svc.db.Where("id = ?", 1).Save(&state).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
