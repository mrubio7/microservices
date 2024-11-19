package repositories

import (
	"ibercs/internal/model"

	"gorm.io/gorm"
)

type PlayersRepository struct {
	*GenericRepository[model.PlayerModel]
}

func NewPlayersRepository(database *gorm.DB) *PlayersRepository {
	return &PlayersRepository{
		GenericRepository: NewGenericRepository[model.PlayerModel](database),
	}
}
