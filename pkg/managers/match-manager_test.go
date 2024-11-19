package managers_test

import (
	"ibercs/internal/faker"
	testutil "ibercs/internal/test"
	"ibercs/pkg/managers"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	schemas := []string{"matches"}

	testutil.SetupTestDBs(schemas)

	code := m.Run()

	testutil.CleanupTestDBs()

	os.Exit(code)
}

func TestMatchCreateWithFaker(t *testing.T) {
	db := testutil.GetTestDB("matches")

	// Create the MatchManager
	manager := managers.NewMatchManager(db)

	// Generate a fake match
	fakeMatch := faker.GenerateMatch(time.Now().Unix())

	// Create the match
	createdMatch, err := manager.UpdateOrCreateMatch(fakeMatch)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch.FaceitId, createdMatch.FaceitId, "FaceitId should match")
	assert.Equal(t, fakeMatch.TeamAName, createdMatch.TeamAName, "TeamAName should match")
}

func TestMatchUpdateWithFaker(t *testing.T) {
	db := testutil.GetTestDB("matches")

	// Create the MatchManager
	manager := managers.NewMatchManager(db)

	// Generate a fake match
	fakeMatch := faker.GenerateMatch(time.Now().Unix())

	// Create the match
	createdMatch, err := manager.UpdateOrCreateMatch(fakeMatch)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch.FaceitId, createdMatch.FaceitId, "FaceitId should match")
	assert.Equal(t, fakeMatch.TeamAName, createdMatch.TeamAName, "TeamAName should match")

	createdMatch.TeamAName = "New Team A Name"
	// Update the match
	updatedMatch, err := manager.UpdateOrCreateMatch(*createdMatch)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, createdMatch.FaceitId, updatedMatch.FaceitId, "FaceitId should match")
	assert.Equal(t, createdMatch.TeamAName, updatedMatch.TeamAName, "TeamAName should match")
}

func TestMatchGetByFaceitId(t *testing.T) {
	db := testutil.GetTestDB("matches")

	// Create the MatchManager
	manager := managers.NewMatchManager(db)

	// Generate a fake match
	fakeMatch := faker.GenerateMatch(time.Now().Unix())

	// Create the match
	createdMatch, err := manager.UpdateOrCreateMatch(fakeMatch)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch.FaceitId, createdMatch.FaceitId, "FaceitId should match")

	// Get the match by FaceitId
	gotMatch, err := manager.GetMatchByFaceitId(fakeMatch.FaceitId)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch.FaceitId, gotMatch.FaceitId, "FaceitId should match")
}

func TestSetStream(t *testing.T) {
	db := testutil.GetTestDB("matches")

	// Create the MatchManager
	manager := managers.NewMatchManager(db)

	// Generate a fake match
	fakeMatch := faker.GenerateMatch(time.Now().Unix())

	// Create the match
	createdMatch, err := manager.UpdateOrCreateMatch(fakeMatch)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch.FaceitId, createdMatch.FaceitId, "FaceitId should match")

	// Set a stream
	err = manager.SetStreamUrl(fakeMatch.FaceitId, "https://twitch.tv/stream")
	assert.Nil(t, err, "Error should be nil")

	// Get the match by FaceitId
	gotMatch, err := manager.GetMatchByFaceitId(fakeMatch.FaceitId)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch.FaceitId, gotMatch.FaceitId, "FaceitId should match")
	assert.Contains(t, gotMatch.Streams, "https://twitch.tv/stream", "Stream should match")
}

func TestGetAllMatches(t *testing.T) {
	db := testutil.GetTestDB("matches")

	// Create the MatchManager
	manager := managers.NewMatchManager(db)

	// Generate a fake match
	fakeMatch1 := faker.GenerateMatch(1)
	fakeMatch2 := faker.GenerateMatch(22)
	fakeMatch3 := faker.GenerateMatch(333)
	fakeMatch4 := faker.GenerateMatch(4444)

	// Create the matches
	createdMatch1, err := manager.UpdateOrCreateMatch(fakeMatch1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch1.FaceitId, createdMatch1.FaceitId, "FaceitId should match")

	createdMatch2, err := manager.UpdateOrCreateMatch(fakeMatch2)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch2.FaceitId, createdMatch2.FaceitId, "FaceitId should match")

	createdMatch3, err := manager.UpdateOrCreateMatch(fakeMatch3)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch3.FaceitId, createdMatch3.FaceitId, "FaceitId should match")

	createdMatch4, err := manager.UpdateOrCreateMatch(fakeMatch4)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch4.FaceitId, createdMatch4.FaceitId, "FaceitId should match")

	// Get all matches
	gotMatches, err := manager.GetAll()
	assert.Nil(t, err, "Error should be nil")
	assert.Len(t, gotMatches, 4, "Length should be 4")
}

func TestGetMatchesByTeam(t *testing.T) {
	db := testutil.GetTestDB("matches")

	// Create the MatchManager
	manager := managers.NewMatchManager(db)

	// Generate a fake match
	fakeMatch1 := faker.GenerateMatch(1)
	fakeMatch2 := faker.GenerateMatch(22)
	fakeMatch3 := faker.GenerateMatch(333)

	teamFaceitID := gofakeit.UUID()
	fakeMatch1.TeamAFaceitId = teamFaceitID
	fakeMatch2.TeamBFaceitId = teamFaceitID
	fakeMatch3.TeamBFaceitId = teamFaceitID

	// Create the matches
	createdMatch1, err := manager.UpdateOrCreateMatch(fakeMatch1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch1.FaceitId, createdMatch1.FaceitId, "FaceitId should match")

	createdMatch2, err := manager.UpdateOrCreateMatch(fakeMatch2)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch2.FaceitId, createdMatch2.FaceitId, "FaceitId should match")

	createdMatch3, err := manager.UpdateOrCreateMatch(fakeMatch3)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeMatch3.FaceitId, createdMatch3.FaceitId, "FaceitId should match")

	// Get the match by TeamId
	gotMatches, err := manager.GetMatchesByTeamId(teamFaceitID)
	assert.Nil(t, err, "Error should be nil")
	assert.Len(t, gotMatches, 3, "Length should be 3")
}
