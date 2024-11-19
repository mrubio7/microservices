package managers

import (
	"errors"
	"ibercs/internal/model"
	"ibercs/internal/repositories"

	"gorm.io/gorm"
)

type MatchManager struct {
	repo *repositories.MatchRepository
}

func NewMatchManager(database *gorm.DB) *MatchManager {
	repo := repositories.NewMatchRepository(database)

	return &MatchManager{repo: repo}
}

func (m *MatchManager) UpdateOrCreateMatch(match model.MatchModel) (*model.MatchModel, error) {
	existingMatch, err := m.repo.GetByFaceitId(match.FaceitId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingMatch == nil {
		if err := m.repo.Create(&match); err != nil {
			return nil, err
		}
		return &match, nil
	}

	if err := m.repo.Update(&match); err != nil {
		return nil, err
	}

	updatedMatch, err := m.repo.GetByFaceitId(match.FaceitId)
	if err != nil {
		return nil, err
	}

	return updatedMatch, nil
}

func (m *MatchManager) GetMatchByFaceitId(faceitId string) (*model.MatchModel, error) {
	return m.repo.GetByFaceitId(faceitId)
}

func (m *MatchManager) GetMatchesByTeamId(teamId string) ([]model.MatchModel, error) {
	return m.repo.GetWhere("team_a_faceit_id = ? OR team_b_faceit_id = ?", teamId, teamId)
}

func (m *MatchManager) SetStreamUrl(faceitId, streamUrl string) error {
	existingMatch, err := m.repo.GetByFaceitId(faceitId)
	if err != nil {
		return err
	}

	for _, stream := range existingMatch.Streams {
		if stream == streamUrl {
			return errors.New("stream already exists")
		}
	}
	existingMatch.Streams = append(existingMatch.Streams, streamUrl)

	if err := m.repo.Update(existingMatch); err != nil {
		return err
	}

	return nil
}

func (m *MatchManager) GetAll() ([]model.MatchModel, error) {
	return m.repo.GetAll()
}
