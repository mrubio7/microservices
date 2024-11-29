package managers

import (
	"fmt"
	"ibercs/internal/model"
	"ibercs/internal/repositories"

	"gorm.io/gorm"
)

type TeamManager struct {
	repoTeams     *repositories.GenericRepository[model.TeamModel]
	repoTeamsRank *repositories.GenericRepository[model.TeamRankModel]
}

func NewTeamManager(database *gorm.DB) *TeamManager {
	teams := repositories.NewGenericRepository[model.TeamModel](database)
	teamRanks := repositories.NewGenericRepository[model.TeamRankModel](database)

	return &TeamManager{
		repoTeams:     teams,
		repoTeamsRank: teamRanks,
	}
}

func (m *TeamManager) Create(team *model.TeamModel) (*model.TeamModel, error) {
	return m.repoTeams.Create(team)
}

func (m *TeamManager) GetAll() ([]model.TeamModel, error) {
	return m.repoTeams.Find()
}

func (m *TeamManager) GetActiveTeams() ([]model.TeamModel, error) {
	return m.repoTeams.Find(repositories.Preload("Stats"), repositories.Where("active = ?", true))
}

func (m *TeamManager) GetById(id int) (*model.TeamModel, error) {
	return m.repoTeams.Get(repositories.Preload("Stats"), repositories.Where("id = ?", id))
}

func (m *TeamManager) GetByFaceitId(faceitId string) (*model.TeamModel, error) {
	return m.repoTeams.Get(repositories.Preload("Stats"), repositories.Where("faceit_id = ?", faceitId))
}

func (m *TeamManager) GetByNickname(nickname string) (*model.TeamModel, error) {
	return m.repoTeams.Get(repositories.Preload("Stats"), repositories.Where("nickname = ?", nickname))
}

func (m *TeamManager) GetByPlayerId(playerId string) ([]model.TeamModel, error) {
	return m.repoTeams.Find(repositories.Preload("Stats"), repositories.Where("players_id @> ?", fmt.Sprintf(`["%s"]`, playerId)))
}

func (m *TeamManager) Update(team *model.TeamModel) error {
	return m.repoTeams.Update(team)
}

func (m *TeamManager) DesactivateTeam(id int) error {
	team, err := m.GetById(id)
	if err != nil {
		return err
	}

	team.Active = false
	return m.repoTeams.Update(team)
}

func (m *TeamManager) ActivateTeam(id int) error {
	team, err := m.GetById(id)
	if err != nil {
		return err
	}

	team.Active = true
	return m.repoTeams.Update(team)
}

// TeamRank
func (m *TeamManager) GetTeamRankByFaceitId(faceitId string) (*model.TeamRankModel, error) {
	return m.repoTeamsRank.Get(repositories.Where("faceit_id = ?", faceitId))
}

func (m *TeamManager) GetAllTeamRank() ([]model.TeamRankModel, error) {
	return m.repoTeamsRank.Find()
}

func (m *TeamManager) UpdateTeamRank(teamRank *model.TeamRankModel) error {
	return m.repoTeamsRank.Update(teamRank)
}

func (m *TeamManager) CreateTeamRank(teamRank *model.TeamRankModel) (*model.TeamRankModel, error) {
	return m.repoTeamsRank.Create(teamRank)
}
