package service

import (
	"errors"
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"sync"

	"gorm.io/gorm"
)

type Teams struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewTeamsService(database *gorm.DB) *Teams {
	return &Teams{
		db: database,
	}
}

func (s *Teams) GetAll(active bool) []model.TeamModel {
	var teams []model.TeamModel

	if active {
		err := s.db.Model(&model.TeamModel{}).Preload("Stats").Where("active = ?", active).Find(&teams).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil
			}
			logger.Error(err.Error())
			return nil
		}
	} else {
		err := s.db.Model(&model.TeamModel{}).Preload("Stats").Find(&teams).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil
			}
			logger.Error(err.Error())
			return nil
		}
	}

	return teams
}

func (s *Teams) GetTeam(id string) *model.TeamModel {
	var team *model.TeamModel

	err := s.db.Model(&model.TeamModel{}).Preload("Stats").First(&team, "faceit_id = ?", id).Error
	if err != nil {
		logger.Error("Team not found: %s", err.Error())
		return nil
	}

	return team
}

func (s *Teams) GetTeamByNickname(nickname string) *model.TeamModel {
	var team *model.TeamModel

	err := s.db.Model(&model.TeamModel{}).Preload("Stats").First(&team, "nickname = ?", nickname).Error
	if err != nil {
		logger.Error("Team not found: %s", err.Error())
		return nil
	}

	return team
}

func (s *Teams) NewTeam(team model.TeamModel) *model.TeamModel {
	var existingTeam model.TeamModel

	err := s.db.Where("faceit_id = ?", team.FaceitId).First(&existingTeam).Error
	if err == nil {
		logger.Warning("Team %s already exist", team.Name)
		return &existingTeam
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := s.db.Create(&team).Error; err != nil {
			logger.Error("Error saving team: %s", err.Error())
			return nil
		}
		return &team
	}

	return nil
}

func (s *Teams) UpdateTeam(team model.TeamModel) *model.TeamModel {
	var existingTeam model.TeamModel

	err := s.db.Where("faceit_id = ?", team.FaceitId).First(&existingTeam).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warning("Team %s does not exist", team.Name)
			return nil
		}
		logger.Error("Error finding team: %s", err.Error())
		return nil
	}

	if err := s.db.Model(&existingTeam).Updates(team).Error; err != nil {
		logger.Error("Error updating team: %s", err.Error())
		return nil
	}

	return &existingTeam
}
