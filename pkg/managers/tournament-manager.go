package managers

import (
	"ibercs/internal/model"
	"ibercs/internal/repositories"

	"gorm.io/gorm"
)

type TournamentManager struct {
	repoOrganizers  *repositories.GenericRepository[model.OrganizerModel]
	repoTournaments *repositories.GenericRepository[model.TournamentModel]
}

func NewTournamentManager(database *gorm.DB) *TournamentManager {
	teams := repositories.NewGenericRepository[model.TournamentModel](database)
	organizers := repositories.NewGenericRepository[model.OrganizerModel](database)

	return &TournamentManager{
		repoTournaments: teams,
		repoOrganizers:  organizers,
	}
}

func (m *TournamentManager) GetTournamentByFaceitId(faceitId string) (*model.TournamentModel, error) {
	return m.repoTournaments.Get(repositories.Where("faceit_id = ?", faceitId))
}

func (m *TournamentManager) GetAllTournaments() ([]model.TournamentModel, error) {
	return m.repoTournaments.Find()
}

func (m *TournamentManager) CreateTournament(tournament *model.TournamentModel) (*model.TournamentModel, error) {
	return m.repoTournaments.Create(tournament)
}

func (m *TournamentManager) UpdateTournament(tournament *model.TournamentModel) error {
	return m.repoTournaments.Update(tournament)
}

// Organizers
func (m *TournamentManager) GetOrganizerByFaceitId(faceitId string) (*model.OrganizerModel, error) {
	return m.repoOrganizers.Get(repositories.Where("faceit_id = ?", faceitId))
}

func (m *TournamentManager) CreateOrganizer(organizer *model.OrganizerModel) (*model.OrganizerModel, error) {
	return m.repoOrganizers.Create(organizer)
}
