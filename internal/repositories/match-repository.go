package repositories

import (
	"errors"
	"ibercs/internal/model"

	"gorm.io/gorm"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(database *gorm.DB) *MatchRepository {
	return &MatchRepository{db: database}
}

func (r *MatchRepository) Create(match interface{}) error {
	return r.db.Model(&model.MatchModel{}).Create(match).Error
}

func (r *MatchRepository) Update(match *model.MatchModel) error {
	if match.ID == 0 {
		return errors.New("cannot update match: missing primary key (ID)")
	}

	// Actualizar los campos del registro basado en el ID
	return r.db.Model(&match).Where("id = ?", match.ID).Updates(match).Error
}

func (r *MatchRepository) GetByFaceitId(faceitId string) (*model.MatchModel, error) {
	var existingMatch model.MatchModel

	err := r.db.Model(&model.MatchModel{}).Where("faceit_id = ?", faceitId).First(&existingMatch).Error
	if err != nil {
		return nil, err
	}

	return &existingMatch, nil
}

func (r *MatchRepository) GetAll() ([]model.MatchModel, error) {
	var matches []model.MatchModel

	err := r.db.Model(&model.MatchModel{}).Find(&matches).Error
	if err != nil {
		return nil, err
	}

	return matches, nil
}

func (r *MatchRepository) DeleteByFaceitId(faceitId string) error {
	return r.db.Model(&model.MatchModel{}).Where("faceit_id = ?", faceitId).Delete(&model.MatchModel{}).Error
}

func (r *MatchRepository) GetWhere(field string, value ...interface{}) ([]model.MatchModel, error) {
	var matches []model.MatchModel

	err := r.db.Where(field, value...).Find(&matches).Error
	if err != nil {
		return nil, err
	}

	return matches, nil
}
