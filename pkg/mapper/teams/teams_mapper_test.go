package teams_mapper_test

import (
	"ibercs/internal/model"
	"ibercs/pkg/mapper"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"

	pb "ibercs/proto/teams"
)

func TestTeamMapper(t *testing.T) {
	var mockTeam model.TeamModel
	gofakeit.Struct(&mockTeam)

	teamProto := mapper.Convert[model.TeamModel, *pb.Team](mockTeam)
	teamModelAgain := mapper.Convert[*pb.Team, model.TeamModel](teamProto)

	assert.Equal(t, teamProto.FaceitId, teamModelAgain.FaceitId, "FaceitId should match")
	assert.Equal(t, teamProto.Id, int32(teamModelAgain.Id), "ID should match")
	assert.Equal(t, teamProto.Stats.Wins, teamModelAgain.Stats.Wins, "MapStats should match")
}
