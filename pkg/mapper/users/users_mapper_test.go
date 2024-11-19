package users_mapper_test

import (
	"psonder/internal/models"
	"psonder/pkg/mapper"
	pb "psonder/proto/users"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserMapper(t *testing.T) {
	t.Log("Mappers registered successfully")

	// Generar datos simulados para UserModel
	var mockUser models.UserModel
	err := faker.FakeData(&mockUser)
	if err != nil {
		t.Fatalf("Error generating mock data: %v", err)
	}

	userProto := mapper.Convert[models.UserModel, *pb.User](mockUser)
	userModelAgain := mapper.Convert[*pb.User, models.UserModel](userProto)

	assert.Equal(t, mockUser, userModelAgain)
}
