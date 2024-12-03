package tournaments_mapper_test

import (
	"ibercs/internal/model"
	"ibercs/pkg/mapper"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	pb "ibercs/proto/tournaments"
)

func TestTournamentMapper(t *testing.T) {
	var mockTournament model.TournamentModel
	gofakeit.Struct(&mockTournament)

	teamProto := mapper.Convert[model.TournamentModel, *pb.Tournament](mockTournament)
	teamModelAgain := mapper.Convert[*pb.Tournament, model.TournamentModel](teamProto)

	assert.Equal(t, teamProto.FaceitId, teamModelAgain.FaceitId, "FaceitId should match")
	assert.Equal(t, teamProto.Id, int32(teamModelAgain.Id), "ID should match")
	assert.Equal(t, teamProto.Name, teamModelAgain.Name, "Name should match")
}
