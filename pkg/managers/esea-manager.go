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

func (m *EseaManager) GetDivisionsByTournamentId(tournamentId string) ([]model.EseaDivisionModel, error) {
	return m.repoEseaDivisions.Find(repositories.Where("tournament_id = ?", tournamentId))
}

func (m *EseaManager) GetStandingsByConferenceId(conferenceId string) ([]model.EseaStandingModel, error) {
	return m.repoEseaStandings.Find(repositories.Where("tournament_id = ?", conferenceId))
}

func (m *EseaManager) UpdateStanding(standing model.EseaStandingModel) error {
	err := m.repoEseaStandings.Update(&standing)
	if err != nil {
		logger.Error("Error updating standing: %s", err.Error())
		return err
	}

	return nil
}
