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

func (m *MatchManager) Create(match *model.MatchModel) (*model.MatchModel, error) {
	return m.repo.Create(match)
}

func (m *MatchManager) Update(match *model.MatchModel) error {
	return m.repo.Update(match, "faceit_id", match.FaceitId)
}

func (m *MatchManager) GetMatchByFaceitId(faceitId string) (*model.MatchModel, error) {
	return m.repo.Get(repositories.Where("faceit_id", faceitId))
}

func (m *MatchManager) GetMatchesByTeamId(teamId string) ([]model.MatchModel, error) {
	return m.repo.Find(repositories.Where("team_a_faceit_id = ? OR team_b_faceit_id = ?", teamId, teamId))
}

func (m *MatchManager) SetStreamUrl(faceitId, streamUrl string) error {
	existingMatch, err := m.GetMatchByFaceitId(faceitId)
	if err != nil {
		return err
	}

	for _, stream := range existingMatch.Streams {
		if stream == streamUrl {
			return errors.New("stream already exists")
		}
	}
	existingMatch.Streams = append(existingMatch.Streams, streamUrl)

	if err := m.Update(existingMatch); err != nil {
		return err
	}

	return nil
}

func (m *MatchManager) GetAll() ([]model.MatchModel, error) {
	return m.repo.Find()
}
