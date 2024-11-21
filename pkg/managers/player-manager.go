package managers

import (
	"fmt"
	"ibercs/internal/model"
	"ibercs/internal/repositories"
	"time"

	"gorm.io/gorm"
)

type PlayerManager struct {
	repoPlayers          *repositories.PlayersRepository
	repoProminentWeek    *repositories.GenericRepository[model.ProminentWeekModel]
	repoProminentPlayers *repositories.GenericRepository[model.PlayerProminentModel]
	repoLFT              *repositories.GenericRepository[model.LookingForTeamModel]
}

func NewPlayerManager(database *gorm.DB) *PlayerManager {
	players := repositories.NewPlayersRepository(database)
	prominent := repositories.NewGenericRepository[model.ProminentWeekModel](database)
	lft := repositories.NewGenericRepository[model.LookingForTeamModel](database)

	return &PlayerManager{
		repoPlayers:       players,
		repoProminentWeek: prominent,
		repoLFT:           lft,
	}
}

func (m *PlayerManager) Create(player *model.PlayerModel) (*model.PlayerModel, error) {
	return m.repoPlayers.Create(player)
}

func (m *PlayerManager) Update(player *model.PlayerModel) error {
	return m.repoPlayers.Update(player)
}

func (m *PlayerManager) GetByNickname(nickname string) (*model.PlayerModel, error) {
	return m.repoPlayers.Get(repositories.Preload("Stats"), repositories.Where("nickname", nickname))
}

func (m *PlayerManager) GetAll() ([]model.PlayerModel, error) {
	return m.repoPlayers.Find(repositories.Preload("Stats"))
}

func (m *PlayerManager) GetByFaceitId(faceitId string) (*model.PlayerModel, error) {
	return m.repoPlayers.Get(repositories.Preload("Stats"), repositories.Where("faceit_id", faceitId))
}

// Prominent players

func (m *PlayerManager) GetProminentPlayers() (*model.ProminentWeekModel, error) {
	return m.repoProminentWeek.Get(repositories.Preload("Players"), repositories.OrderBy("year DESC, week DESC"))
}

func (m *PlayerManager) CreateProminentPlayers(prominent *model.ProminentWeekModel) (*model.ProminentWeekModel, error) {
	return m.repoProminentWeek.Create(prominent)
}

func (m *PlayerManager) GenerateProminentPlayers() (*model.ProminentWeekModel, error) {
	year, week := time.Now().ISOWeek()

	query := fmt.Sprintf(`
		WITH previous_week AS (
			SELECT ppm.id
			FROM player_prominent_models ppm
			INNER JOIN prominent_week_models pwm ON ppm.prominent_week_id = pwm.id
			WHERE pwm.year = %d AND pwm.week = %d
		)

		SELECT ps.id, p.avatar, p.nickname, 
			((kills_average * 0.35) - (deaths_average * 0.15) + (assist_average * 0.1) + (kr_ratio * 0.2) + (mvp_average * 0.1)) * (1 + (elo / 1000)) AS calc
		FROM player_stats_models ps 
		INNER JOIN player_models p ON p.id = ps.id 
		LEFT JOIN player_prominent_models ppm ON ppm.id = ps.id
		LEFT JOIN prominent_week_models pwm ON pwm.id = ppm.prominent_week_id
		WHERE (pwm.week IS NULL OR pwm.year < %d OR pwm.week < %d)
		AND ps.id NOT IN (SELECT id FROM previous_week)
		ORDER BY calc DESC
		LIMIT 5;
	`, year, week-1, year, week+1)

	results, err := m.repoProminentPlayers.FindRawQuery(query)
	if err != nil {
		return nil, err
	}

	prominentWeek := model.ProminentWeekModel{
		Week:    int16(week),
		Year:    int16(year),
		Players: results,
	}

	return m.CreateProminentPlayers(&prominentWeek)
}

// Looking for team

func (m *PlayerManager) GetLookingForTeamPlayers() ([]model.LookingForTeamModel, error) {
	return m.repoLFT.Find()
}

func (m *PlayerManager) CreateLookingForTeamPlayer(lft *model.LookingForTeamModel) (*model.LookingForTeamModel, error) {
	return m.repoLFT.Create(lft)
}

func (m *PlayerManager) UpdateLookingForTeamPlayer(lft *model.LookingForTeamModel) error {
	return m.repoLFT.Update(lft)
}

func (m *PlayerManager) DeleteLookingForTeamPlayer(faceitId string) error {
	return m.repoLFT.Delete("faceit_id", faceitId)
}
