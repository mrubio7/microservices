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

func (m *StateManager) Get() (*model.StateModel, error) {
	return m.repoState.Get(repositories.Where("id = ?", 1))
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

func (m *StateManager) Update_TournamentLastFind() error {
	state, err := m.repoState.Get()
	if err != nil {
		return err
	}

	state.TournamentsLastFind.Time = time.Now()
	state.TournamentsLastFind.Valid = true

	return m.repoState.Update(state)
}

func (m *StateManager) Update_MatchLastFind() error {
	state, err := m.repoState.Get()
	if err != nil {
		return err
	}

	state.MatchesLastFind.Time = time.Now()
	state.MatchesLastFind.Valid = true

	return m.repoState.Update(state)
}

func (m *StateManager) Update_TeamLastUpdate() error {
	state, err := m.repoState.Get()
	if err != nil {
		return err
	}

	state.TeamsLastUpdate.Time = time.Now()
	state.TeamsLastUpdate.Valid = true

	return m.repoState.Update(state)
}
