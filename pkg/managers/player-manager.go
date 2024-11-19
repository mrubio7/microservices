package managers

import (
	"ibercs/internal/model"
	"ibercs/internal/repositories"

	"gorm.io/gorm"
)

type PlayerManager struct {
	repoPlayers   *repositories.PlayersRepository
	repoProminent *repositories.GenericRepository[model.ProminentWeekModel]
}

func NewPlayerManager(database *gorm.DB) *PlayerManager {
	players := repositories.NewPlayersRepository(database)
	prominent := repositories.NewGenericRepository[model.ProminentWeekModel](database)

	return &PlayerManager{
		repoPlayers:   players,
		repoProminent: prominent,
	}
}

func (m *PlayerManager) Create(player *model.PlayerModel) (*model.PlayerModel, error) {
	return m.repoPlayers.Create(player)
}

func (m *PlayerManager) Update(player *model.PlayerModel) error {
	return m.repoPlayers.Update(player, "faceit_id", player.FaceitId)
}

func (m *PlayerManager) GetByNickname(nickname string) (*model.PlayerModel, error) {
	return m.repoPlayers.Get(repositories.Preload("Stats"), repositories.Where("nickname", nickname))
}

func (m *PlayerManager) GetAll() ([]model.PlayerModel, error) {
	return m.repoPlayers.Find()
}

func (m *PlayerManager) GetProminentPlayers() (*model.ProminentWeekModel, error) {
	return m.repoProminent.Get(repositories.Preload("Players"), repositories.OrderBy("year DESC, week DESC"))
}

func (m *PlayerManager) CreateProminentPlayers(prominent *model.ProminentWeekModel) (*model.ProminentWeekModel, error) {
	return m.repoProminent.Create(prominent)
}
