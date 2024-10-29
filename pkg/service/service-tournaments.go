package service

import (
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"sync"

	"gorm.io/gorm"
)

type Tournaments struct {
	db    *gorm.DB
	mutex sync.Mutex
}

func NewTournamentsService(database *gorm.DB) *Tournaments {
	return &Tournaments{
		db: database,
	}
}

func (svc *Tournaments) NewOrganizer(organizer *model.OrganizerModel) *model.OrganizerModel {
	if err := svc.db.Model(&model.OrganizerModel{}).Create(&organizer).Error; err != nil {
		return nil
	}

	return organizer
}

func (svc *Tournaments) NewTournament(tournament *model.TournamentModel) *model.TournamentModel {
	if err := svc.db.Model(&model.TournamentModel{}).Create(&tournament).Error; err != nil {
		return nil
	}

	return tournament
}

func (svc *Tournaments) GetAllTournaments() []model.TournamentModel {
	var tournaments []model.TournamentModel

	err := svc.db.Model(&model.TournamentModel{}).Find(&tournaments).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return tournaments
}

func (svc *Tournaments) GetAllOrganizers() []model.OrganizerModel {
	var organizers []model.OrganizerModel

	err := svc.db.Model(&model.OrganizerModel{}).Find(&organizers).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return organizers
}

func (svc *Tournaments) GetOrganizer(faceitId string) *model.OrganizerModel {
	var organizer model.OrganizerModel

	err := svc.db.Model(&model.OrganizerModel{}).Where("faceit_id = ?", faceitId).First(&organizer).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return &organizer
}

func (svc *Tournaments) GetTournament(faceitId string) *model.TournamentModel {
	var tournament model.TournamentModel

	err := svc.db.Model(&model.TournamentModel{}).Where("faceit_id = ?", faceitId).First(&tournament).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return &tournament
}
