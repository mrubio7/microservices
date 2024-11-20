package teams_mapper_test

import (
	"ibercs/internal/model"
	"ibercs/pkg/mapper"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	pb "ibercs/proto/users"
)

func TestUserMapper(t *testing.T) {
	var mockUser model.UserModel
	gofakeit.Struct(&mockUser)

	userProto := mapper.Convert[model.UserModel, *pb.User](mockUser)
	userModelAgain := mapper.Convert[*pb.User, model.UserModel](userProto)

	assert.Equal(t, userProto.PlayerID, userModelAgain.FaceitId, "FaceitId should match")
	assert.Equal(t, userProto.ID, int32(userModelAgain.ID), "ID should match")
	assert.Equal(t, userProto.Name, userModelAgain.Name, "Name should match")
}
