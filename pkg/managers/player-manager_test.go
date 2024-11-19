package managers_test

import (
	"ibercs/internal/faker"
	testutil "ibercs/internal/test"
	"ibercs/pkg/managers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPlayerCreate(t *testing.T) {
	db := testutil.GetTestDB("players")

	manager := managers.NewPlayerManager(db)

	fakePlayer := faker.GeneratePlayer(time.Now().Unix())

	createdPlayer, err := manager.Create(&fakePlayer)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakePlayer.FaceitId, createdPlayer.FaceitId, "FaceitId should match")
}

func TestPlayerUpdate(t *testing.T) {
	db := testutil.GetTestDB("players")

	manager := managers.NewPlayerManager(db)

	fakePlayer := faker.GeneratePlayer(time.Now().Unix())

	createdPlayer, err := manager.Create(&fakePlayer)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakePlayer.FaceitId, createdPlayer.FaceitId, "FaceitId should match")

	createdPlayer.Nickname = "New Nickname"

	err = manager.Update(createdPlayer)
	assert.Nil(t, err, "Error should be nil")

	updatedPlayer, err := manager.GetByFaceitId(createdPlayer.FaceitId)
	assert.Nil(t, err, "Error should be nil")

	assert.Equal(t, createdPlayer.FaceitId, updatedPlayer.FaceitId, "FaceitId should match")
}

func TestGetByNickname(t *testing.T) {
	db := testutil.GetTestDB("players")

	manager := managers.NewPlayerManager(db)

	fakePlayer := faker.GeneratePlayer(time.Now().Unix())

	createdPlayer, err := manager.Create(&fakePlayer)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakePlayer.FaceitId, createdPlayer.FaceitId, "FaceitId should match")

	player, err := manager.GetByNickname(createdPlayer.Nickname)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, createdPlayer.FaceitId, player.FaceitId, "FaceitId should match")
}

func TestGetAll(t *testing.T) {
	db := testutil.GetTestDB("players")

	manager := managers.NewPlayerManager(db)

	fakePlayer1 := faker.GeneratePlayer(345623)
	fakePlayer2 := faker.GeneratePlayer(6457)

	createdPlayer1, err := manager.Create(&fakePlayer1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakePlayer1.FaceitId, createdPlayer1.FaceitId, "FaceitId should match")

	createdPlayer2, err := manager.Create(&fakePlayer2)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakePlayer2.FaceitId, createdPlayer2.FaceitId, "FaceitId should match")

	players, err := manager.GetAll()
	assert.Nil(t, err, "Error should be nil")
	assert.Len(t, players, 2, "Players should not be empty")
}

func TestGetProminentPlayers(t *testing.T) {
	db := testutil.GetTestDB("players")

	manager := managers.NewPlayerManager(db)

	fakeProminent := faker.GenerateProminentWeek(time.Now().Unix())

	createdProminent, err := manager.CreateProminentPlayers(&fakeProminent)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeProminent.Year, createdProminent.Year, "Year should match")
	assert.Equal(t, fakeProminent.Week, createdProminent.Week, "Week should match")

	prominent, err := manager.GetProminentPlayers()
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, createdProminent.Year, prominent.Year, "Year should match")
	assert.Equal(t, createdProminent.Week, prominent.Week, "Week should match")
}
