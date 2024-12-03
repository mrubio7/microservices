package managers_test

import (
	"ibercs/internal/faker"
	testutil "ibercs/internal/test"
	"ibercs/pkg/managers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrganizer(t *testing.T) {
	db := testutil.GetTestDB("tournaments")

	managers := managers.NewTournamentManager(db)

	organizer := faker.GenerateOrganizer(5214)

	_, err := managers.CreateOrganizer(&organizer)
	assert.Nil(t, err, "Error should be nil")

	organizerFound, err := managers.GetOrganizerByFaceitId(organizer.FaceitId)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, organizer.FaceitId, organizerFound.FaceitId, "FaceitId should match")
}

func TestCreateTournament(t *testing.T) {
	db := testutil.GetTestDB("tournaments")

	managers := managers.NewTournamentManager(db)

	tournament := faker.GenerateTournament(489)

	_, err := managers.CreateTournament(&tournament)
	assert.Nil(t, err, "Error should be nil")

	tournamentFound, err := managers.GetTournamentByFaceitId(tournament.FaceitId)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, tournament.FaceitId, tournamentFound.FaceitId, "FaceitId should match")
}

func TestUpdateTournament(t *testing.T) {
	db := testutil.GetTestDB("tournaments")

	managers := managers.NewTournamentManager(db)

	tournament := faker.GenerateTournament(9842)

	_, err := managers.CreateTournament(&tournament)
	assert.Nil(t, err, "Error should be nil")

	tournament.Name = "New Tournament Name"
	err = managers.UpdateTournament(&tournament)
	assert.Nil(t, err, "Error should be nil")

	tournamentFound, err := managers.GetTournamentByFaceitId(tournament.FaceitId)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, tournament.FaceitId, tournamentFound.FaceitId, "FaceitId should match")
	assert.Equal(t, tournament.Name, tournamentFound.Name, "Name should match")
}
