package service

import (
	"ibercs/internal/model"
	"ibercs/pkg/logger"
	"reflect"
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
	tx := svc.db.Begin()

	if err := tx.Model(&model.OrganizerModel{}).Create(&organizer).Error; err != nil {
		tx.Rollback()
		logger.Error("error creating organizer %s", organizer.Name)
		return nil
	}
	tx.Commit()
	return organizer
}

func (svc *Tournaments) NewTournament(tournament *model.TournamentModel) *model.TournamentModel {
	tx := svc.db.Begin()

	if err := tx.Model(&model.TournamentModel{}).Create(&tournament).Error; err != nil {
		tx.Rollback()
		logger.Error(err.Error())
		return nil
	}

	tx.Commit()
	return tournament
}

func (svc *Tournaments) GetAllTournaments(active bool) []model.TournamentModel {
	var tournaments []model.TournamentModel

	if active {
		err := svc.db.Model(&model.TournamentModel{}).Where("status != ? || status != ?", "finished", "cancelled").Find(&tournaments).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil
			}
			logger.Error(err.Error())
			return nil
		}
	} else {
		err := svc.db.Model(&model.TournamentModel{}).Find(&tournaments).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil
			}
			logger.Error(err.Error())
			return nil
		}
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

func (svc *Tournaments) UpdateTournament(tournament *model.TournamentModel) error {
	var existingTournament model.TournamentModel

	svc.mutex.Lock()
	defer svc.mutex.Unlock()

	err := svc.db.First(&existingTournament, "faceit_id = ?", tournament.FaceitId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			t := svc.NewTournament(tournament)
			if t == nil {
				logger.Error("Error cannot create tournament")
			}
			return nil
		}
		logger.Error(err.Error())
		return err
	}

	if !reflect.DeepEqual(existingTournament, tournament) {
		if err := svc.db.Model(&existingTournament).Updates(tournament).Error; err != nil {
			return err
		}
	}

	return nil
}

func (svc *Tournaments) NewEseaDivision(division model.EseaDivisionModel) *model.EseaDivisionModel {
	tx := svc.db.Begin()

	if err := tx.Model(&model.EseaDivisionModel{}).Create(&division).Error; err != nil {
		tx.Rollback()
		logger.Error(err.Error())
		return nil
	}
	tx.Commit()

	return &division
}

func (svc *Tournaments) UpdateEseaDivision(tournament model.EseaDivisionModel) error {
	var existingDivision model.EseaDivisionModel

	svc.mutex.Lock()
	defer svc.mutex.Unlock()

	err := svc.db.First(&existingDivision, "faceit_id = ?", tournament.FaceitId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			t := svc.NewEseaDivision(tournament)
			if t == nil {
				logger.Error("Error cannot create tournament")
			}
			return nil
		}
		logger.Error(err.Error())
		return err
	}

	if !reflect.DeepEqual(existingDivision, tournament) {
		if err := svc.db.Model(&existingDivision).Updates(tournament).Error; err != nil {
			return err
		}
	}

	return nil
}

func (svc *Tournaments) GetEseaDivisions(faceitId string) []model.EseaDivisionModel {
	var tournament []model.EseaDivisionModel

	err := svc.db.Model(&model.EseaDivisionModel{}).Where("tournament_id = ?", faceitId).Find(&tournament).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil
		}
		logger.Error(err.Error())
		return nil
	}

	return tournament
}

func (svc *Tournaments) UpdateEseaDivisionStanding(Standings model.EseaStandingModel) error {
	var existingStanding model.EseaStandingModel

	svc.mutex.Lock()
	defer svc.mutex.Unlock()

	err := svc.db.First(&existingStanding, "faceit_id = ? and tournament_id = ?", Standings.FaceitId, Standings.TournamentId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			tx := svc.db.Begin()

			if err := tx.Model(&model.EseaStandingModel{}).Create(&Standings).Error; err != nil {
				tx.Rollback()
				logger.Error(err.Error())
				return err
			}
			tx.Commit()
			return nil
		}
		logger.Error(err.Error())
		return err
	}

	if !reflect.DeepEqual(existingStanding, Standings) {
		if err := svc.db.Model(&existingStanding).Updates(Standings).Error; err != nil {
			return err
		}
	}

	return nil
}
