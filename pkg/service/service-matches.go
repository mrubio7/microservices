package service

import (
	"errors"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"sync"

	"gorm.io/gorm"
)

type Matches struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewMatchesService(database *gorm.DB) *Matches {
	return &Matches{
		db: database,
	}
}

func (svc *Matches) SaveMatch(match model.MatchModel) *model.MatchModel {
	var existingMatch model.MatchModel

	err := svc.db.Model(&model.MatchModel{}).Where("faceit_id = ?", match.FaceitId).First(&existingMatch).Error
	if err == nil {
		logger.Warning("Match %s already exist", match.FaceitId)
		return &existingMatch
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := svc.db.Model(&model.MatchModel{}).Create(&match).Error; err != nil {
			logger.Error("Error saving match: %s", err.Error())
			return nil
		}
		return &match
	}

	return nil
}

func (svc *Matches) GetMatchByFaceitId(faceitId string) *model.MatchModel {
	var existingMatch model.MatchModel

	err := svc.db.Model(&model.MatchModel{}).Where("faceit_id = ?", faceitId).First(&existingMatch).Error
	if err != nil {
		logger.Error("Error getting match: %s", err)
		return nil
	}

	return &existingMatch
}

func (svc *Matches) GetAllMatches() []model.MatchModel {
	var matches []model.MatchModel

	err := svc.db.Model(&model.MatchModel{}).Find(&matches).Error
	if err != nil {
		logger.Error("Error getting match: %s", err)
		return nil
	}

	return matches
}
