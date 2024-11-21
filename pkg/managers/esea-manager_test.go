package managers_test

import (
	"ibercs/internal/faker"
	testutil "ibercs/internal/test"
	"ibercs/pkg/managers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEseaLeague(t *testing.T) {
	db := testutil.GetTestDB("tournaments")

	managers := managers.NewEseaManager(db)

	league := faker.GenerateEseaLeague(98767)

	leagueCreated, err := managers.CreateEseaLeague(&league)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, league.FaceitId, leagueCreated.FaceitId, "FaceitId should match")

	leagueFound, err := managers.GetEseaLeagueBySeasonNumber(league.Season)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, league.FaceitId, leagueFound.FaceitId, "FaceitId should match")
	assert.Equal(t, league.Divisions[1].FaceitId, leagueFound.Divisions[1].FaceitId, "Name should match")
	assert.Equal(t, league.Divisions[1].Standings[1].TeamFaceitId, leagueFound.Divisions[1].Standings[1].TeamFaceitId, "Name should match")
}
