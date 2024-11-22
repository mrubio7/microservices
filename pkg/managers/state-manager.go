package managers

import (
	"ibercs/internal/model"
	"ibercs/internal/repositories"
	"time"

	"gorm.io/gorm"
)

type StateManager struct {
	repoState *repositories.GenericRepository[model.StateModel]
}

func NewStateManager(database *gorm.DB) *StateManager {
	stateRepo := repositories.NewGenericRepository[model.StateModel](database)

	return &StateManager{
		repoState: stateRepo,
	}
}

func (m *StateManager) Update_PlayersLastUpdate() error {
	state, err := m.repoState.Get()
	if err != nil {
		return err
	}

	state.PlayersLastUpdate.Time = time.Now()
	state.PlayersLastUpdate.Valid = true

	return m.repoState.Update(state)
}
