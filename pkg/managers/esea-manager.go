package managers

import (
	"ibercs/internal/model"
	"ibercs/internal/repositories"
	"ibercs/pkg/logger"

	"gorm.io/gorm"
)

type EseaManager struct {
	repoEseaDivisions *repositories.GenericRepository[model.EseaDivisionModel]
	repoEseaStandings *repositories.GenericRepository[model.EseaStandingModel]
	repoEseaLeague    *repositories.GenericRepository[model.EseaLeagueModel]
	repoTournament    *repositories.GenericRepository[model.TournamentModel]
}

func NewEseaManager(database *gorm.DB) *EseaManager {
	eseaDivisions := repositories.NewGenericRepository[model.EseaDivisionModel](database)
	eseaStandings := repositories.NewGenericRepository[model.EseaStandingModel](database)
	eseaLeague := repositories.NewGenericRepository[model.EseaLeagueModel](database)
	tournamentLeague := repositories.NewGenericRepository[model.TournamentModel](database)

	return &EseaManager{
		repoEseaDivisions: eseaDivisions,
		repoEseaStandings: eseaStandings,
		repoEseaLeague:    eseaLeague,
		repoTournament:    tournamentLeague,
	}
}

func (m *EseaManager) GetEseaLeagueLive() (*model.EseaLeagueModel, error) {
	return m.repoEseaLeague.Get(repositories.Preload("Divisions"), repositories.Preload("Divisions.Standings"), repositories.Where("status = ?", "live"))
}

func (m *EseaManager) GetEseaLeagueBySeasonNumber(seasonNumber int32) (*model.EseaLeagueModel, error) {
	return m.repoEseaLeague.Get(repositories.Preload("Divisions"), repositories.Preload("Divisions.Standings"), repositories.Where("season = ?", seasonNumber))
}

func (m *EseaManager) CreateEseaLeague(league *model.EseaLeagueModel) (*model.EseaLeagueModel, error) {
	return m.repoEseaLeague.Create(league)
}

func (m *EseaManager) CreateDivision(division *model.EseaDivisionModel) (*model.EseaDivisionModel, error) {
	return m.repoEseaDivisions.Create(division)
}

func (m *EseaManager) GetDivisionsByEseaLeagueFaceitId(tournamentId string) ([]model.EseaDivisionModel, error) {
	return m.repoEseaDivisions.Find(repositories.Where("tournament_id = ?", tournamentId))
}

func (m *EseaManager) GetDivisionByFaceitId(divisionId string) (*model.EseaDivisionModel, error) {
	return m.repoEseaDivisions.Get(repositories.Where("conference_id = ?", divisionId))
}

func (m *EseaManager) GetStandingsByDivisionId(divisionId int) ([]model.EseaStandingModel, error) {
	return m.repoEseaStandings.Find(repositories.Where("division_id = ?", divisionId))
}

func (m *EseaManager) GetStandingByTeamFaceitIdAndDivisionId(teamFaceitId string, divisionId int) (*model.EseaStandingModel, error) {
	return m.repoEseaStandings.Get(repositories.Where("team_faceit_id = ?", teamFaceitId), repositories.Where("division_id = ?", divisionId))
}

func (m *EseaManager) UpdateStanding(standing model.EseaStandingModel) error {
	err := m.repoEseaStandings.Update(&standing)
	if err != nil {
		logger.Error("Error updating standing: %s", err.Error())
		return err
	}

	return nil
}

func (m *EseaManager) UpdateDivision(division model.EseaDivisionModel) error {
	return m.repoEseaDivisions.Update(&division)
}
