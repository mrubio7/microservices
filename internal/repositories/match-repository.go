package repositories

import (
	"ibercs/internal/model"

	"gorm.io/gorm"
)

type MatchRepository struct {
	*GenericRepository[model.MatchModel]
}

func NewMatchRepository(database *gorm.DB) *MatchRepository {
	return &MatchRepository{
		GenericRepository: NewGenericRepository[model.MatchModel](database),
	}
}
