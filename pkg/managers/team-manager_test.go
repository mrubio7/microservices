package managers_test

import (
	"ibercs/internal/faker"
	testutil "ibercs/internal/test"
	"ibercs/pkg/managers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamCreate(t *testing.T) {
	db := testutil.GetTestDB("teams")

	managers := managers.NewTeamManager(db)

	fakeTeam := faker.GenerateTeam(6342)

	createdTeam, err := managers.Create(&fakeTeam)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam.FaceitId, createdTeam.FaceitId, "FaceitId should match")

	team, err := managers.GetById(int(createdTeam.Id))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, createdTeam.FaceitId, team.FaceitId, "FaceitId should match")
	assert.Equal(t, createdTeam.Stats.TotalMatches, team.Stats.TotalMatches, "TotalMatches should match")
}

func TestTeamUpdate(t *testing.T) {
	db := testutil.GetTestDB("teams")

	managers := managers.NewTeamManager(db)

	fakeTeam := faker.GenerateTeam(7547)

	createdTeam, err := managers.Create(&fakeTeam)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam.FaceitId, createdTeam.FaceitId, "FaceitId should match")

	createdTeam.Avatar = "https://cdn.faceit.com/avatar/1234567890"
	createdTeam.Stats.TotalMatches = 10

	err = managers.Update(createdTeam)
	assert.Nil(t, err, "Error should be nil")

	team, err := managers.GetByNickname(createdTeam.Nickname)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, createdTeam.Avatar, team.Avatar, "FaceitId should match")
	assert.Equal(t, createdTeam.Stats.TotalMatches, team.Stats.TotalMatches, "TotalMatches should match")
}

func TestTeamActives(t *testing.T) {
	db := testutil.GetTestDB("teams")

	managers := managers.NewTeamManager(db)

	fakeTeam1 := faker.GenerateTeam(1234)

	createdTeam1, err := managers.Create(&fakeTeam1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam1.FaceitId, createdTeam1.FaceitId, "FaceitId should match")

	fakeTeam2 := faker.GenerateTeam(5678)
	fakeTeam2.Active = true

	createdTeam2, err := managers.Create(&fakeTeam2)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam2.FaceitId, createdTeam2.FaceitId, "FaceitId should match")

	teams, err := managers.GetActiveTeams()
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 1, len(teams), "Only one team should be active")
	assert.Equal(t, createdTeam2.FaceitId, teams[0].FaceitId, "FaceitId should match")
}

func TestTeamGetByPlayerId(t *testing.T) {
	db := testutil.GetTestDB("teams")

	managers := managers.NewTeamManager(db)

	fakeTeam1 := faker.GenerateTeam(1234)

	createdTeam1, err := managers.Create(&fakeTeam1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam1.FaceitId, createdTeam1.FaceitId, "FaceitId should match")

	fakeTeam2 := faker.GenerateTeam(5678)

	createdTeam2, err := managers.Create(&fakeTeam2)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam2.FaceitId, createdTeam2.FaceitId, "FaceitId should match")

	teams, err := managers.GetByPlayerId(fakeTeam1.PlayersId[0])
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 1, len(teams), "Only one team should be active")
	assert.Equal(t, createdTeam1.FaceitId, teams[0].FaceitId, "FaceitId should match")
}

func TestTeamGetByFaceitId(t *testing.T) {
	db := testutil.GetTestDB("teams")

	managers := managers.NewTeamManager(db)

	fakeTeam := faker.GenerateTeam(1234)

	createdTeam, err := managers.Create(&fakeTeam)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam.FaceitId, createdTeam.FaceitId, "FaceitId should match")

	team, err := managers.GetByFaceitId(fakeTeam.FaceitId)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam.FaceitId, team.FaceitId, "FaceitId should match")
}

func TestTeamActiveDesactive(t *testing.T) {
	db := testutil.GetTestDB("teams")

	managers := managers.NewTeamManager(db)

	fakeTeam := faker.GenerateTeam(1234)

	createdTeam, err := managers.Create(&fakeTeam)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeTeam.FaceitId, createdTeam.FaceitId, "FaceitId should match")

	err = managers.ActivateTeam(int(createdTeam.Id))
	assert.Nil(t, err, "Error should be nil")

	team, err := managers.GetById(int(createdTeam.Id))
	assert.Nil(t, err, "Error should be nil")
	assert.True(t, team.Active, "Team should be inactive")

	err = managers.DesactivateTeam(int(createdTeam.Id))
	assert.Nil(t, err, "Error should be nil")

	team, err = managers.GetById(int(createdTeam.Id))
	assert.Nil(t, err, "Error should be nil")
	assert.False(t, team.Active, "Team should be active")

	assert.Equal(t, fakeTeam.FaceitId, team.FaceitId, "FaceitId should match")
}
