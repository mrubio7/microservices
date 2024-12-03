package players_mapper_test

import (
	"ibercs/internal/model"
	"ibercs/pkg/mapper"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	pb "ibercs/proto/players"
)

func TestPlayerMapper(t *testing.T) {
	var mockPlayer model.PlayerModel
	gofakeit.Struct(&mockPlayer)

	playerProto := mapper.Convert[model.PlayerModel, *pb.Player](mockPlayer)
	playerModelAgain := mapper.Convert[*pb.Player, model.PlayerModel](playerProto)

	assert.Equal(t, playerProto.FaceitId, playerModelAgain.FaceitId, "FaceitId should match")
	assert.Equal(t, playerProto.Id, int32(playerModelAgain.Id), "ID should match")
}
