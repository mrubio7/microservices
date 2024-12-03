package matches_mapper_test

import (
	"ibercs/internal/model"
	"ibercs/pkg/mapper"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	pb "ibercs/proto/matches"
)

func TestUserMapper(t *testing.T) {
	t.Log("Mappers registered successfully")

	var mockMatch model.MatchModel
	gofakeit.Struct(&mockMatch)

	userProto := mapper.Convert[model.MatchModel, *pb.Match](mockMatch)
	userModelAgain := mapper.Convert[*pb.Match, model.MatchModel](userProto)

	assert.Equal(t, userProto.BestOf, userModelAgain.BestOf, "BestOf should match")
	assert.Equal(t, userProto.FaceitId, userModelAgain.FaceitId, "FaceitId should match")
	assert.Equal(t, userProto.ID, int32(userModelAgain.ID), "ID should match")
}
