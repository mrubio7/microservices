package managers_test

import (
	"ibercs/internal/faker"
	testutil "ibercs/internal/test"
	"ibercs/pkg/managers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	db := testutil.GetTestDB("users")

	manager := managers.NewUserManager(db)

	fakeUser := faker.GenerateUser(time.Now().Unix())

	createdUser, err := manager.Create(&fakeUser)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeUser.FaceitId, createdUser.FaceitId, "FaceitId should match")
}

func TestUserUpdate(t *testing.T) {
	db := testutil.GetTestDB("users")

	manager := managers.NewUserManager(db)

	fakeUser := faker.GenerateUser(time.Now().Unix())

	createdUser, err := manager.Create(&fakeUser)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeUser.FaceitId, createdUser.FaceitId, "FaceitId should match")

	createdUser.Name = "New Name"
	_, err = manager.Update(createdUser)
	assert.Nil(t, err, "Error should be nil")

	updatedUser, err := manager.GetByFaceitId(createdUser.FaceitId)
	assert.Nil(t, err, "Error should be nil")

	assert.Equal(t, createdUser.FaceitId, updatedUser.FaceitId, "FaceitId should match")
}

func TestSessions(t *testing.T) {
	db := testutil.GetTestDB("users")

	manager := managers.NewUserManager(db)

	fakeUser := faker.GenerateUser(time.Now().Unix())

	createdUser, err := manager.Create(&fakeUser)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeUser.FaceitId, createdUser.FaceitId, "FaceitId should match")

	session, err := manager.CreateNewSession(int(createdUser.ID))
	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, session.SessionID, "SessionID should not be empty")

	session2, err := manager.GetSessionByUserId(int(createdUser.ID))
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, session.SessionID, session2.SessionID, "SessionID should match")

	err = manager.DeleteSession(int(createdUser.ID))
	assert.Nil(t, err, "Error should be nil")

	session3, err := manager.GetSessionByUserId(int(createdUser.ID))
	assert.NotNil(t, err, "Error should not be nil")
	assert.Empty(t, session3, "SessionID should be empty")
}

func TestStreams(t *testing.T) {
	db := testutil.GetTestDB("users")

	manager := managers.NewUserManager(db)

	fakeUser1 := faker.GenerateUser(23423)
	fakeUser2 := faker.GenerateUser(74265)

	createdUser1, err := manager.Create(&fakeUser1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeUser1.FaceitId, createdUser1.FaceitId, "FaceitId should match")

	createdUser2, err := manager.Create(&fakeUser2)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, fakeUser2.FaceitId, createdUser2.FaceitId, "FaceitId should match")

	streams, err := manager.GetAllStreams()
	assert.Nil(t, err, "Error should be nil")

	assert.Len(t, streams, 2, "FaceitId should match")
}
