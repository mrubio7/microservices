package managers

import (
	"ibercs/internal/model"
	"ibercs/internal/repositories"
	"time"

	"gorm.io/gorm"
)

type MatchManager struct {
	repo *repositories.GenericRepository[model.MatchModel]
}

func NewMatchManager(database *gorm.DB) *MatchManager {
	repo := repositories.NewGenericRepository[model.MatchModel](database)

	return &MatchManager{repo: repo}
}

func (m *MatchManager) Create(match *model.MatchModel) (*model.MatchModel, error) {
	return m.repo.Create(match)
}

func (m *MatchManager) Update(match *model.MatchModel) error {
	return m.repo.Update(match)
}

func (m *MatchManager) GetMatchByFaceitId(faceitId string) (*model.MatchModel, error) {
	return m.repo.Get(repositories.Where("faceit_id = ?", faceitId))
}

func (m *MatchManager) GetUpcomingMatches() ([]model.MatchModel, error) {
	return m.repo.Find(repositories.Where("status != ?", "FINISHED"), repositories.Where("status != ?", "CANCELLED"))
}

func (m *MatchManager) GetTodayMatches() ([]model.MatchModel, error) {
	currentDate := time.Now().Format("2006-01-02")
	return m.repo.Find(repositories.Where("DATE(timestamp) = ?", currentDate))
}

func (m *MatchManager) GetYesterdayMatches() ([]model.MatchModel, error) {
	currentDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	return m.repo.Find(repositories.Where("DATE(timestamp) = ?", currentDate))
}

func (m *MatchManager) GetNearbyMatches(days int) ([]model.MatchModel, error) {
	past := time.Now().AddDate(0, 0, -days).Format("2006-01-02")
	future := time.Now().AddDate(0, 0, days).Format("2006-01-02")

	return m.repo.Find(repositories.Where("DATE(timestamp) BETWEEN ? AND ?", past, future))
}

func (m *MatchManager) GetMatchesByTeamId(teamId string) ([]model.MatchModel, error) {
	return m.repo.Find(
		repositories.Where("team_a_faceit_id = ? OR team_b_faceit_id = ?", teamId, teamId),
		repositories.OrderBy("timestamp DESC"),
	)
}

func (m *MatchManager) SetStreamUrl(faceitId, streamUrl string) error {
	existingMatch, err := m.GetMatchByFaceitId(faceitId)
	if err != nil {
		return err
	}

	for _, stream := range existingMatch.Streams {
		if stream == streamUrl {
			return gorm.ErrDuplicatedKey
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
