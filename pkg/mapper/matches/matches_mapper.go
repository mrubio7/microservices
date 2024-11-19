package matches_mapper

import (
	"ibercs/internal/model"
	pb "ibercs/proto/matches"
)

type MatchMapper struct{}

func (MatchMapper) Proto(model model.MatchModel) *pb.Match {
	return &pb.Match{
		ID:            int32(model.ID),
		FaceitId:      model.FaceitId,
		TeamAName:     model.TeamAName,
		TeamBName:     model.TeamBName,
		ScoreTeamA:    model.ScoreTeamA,
		ScoreTeamB:    model.ScoreTeamB,
		TeamAFaceitId: model.TeamAFaceitId,
		TeamBFaceitId: model.TeamBFaceitId,
		IsTeamAKnown:  model.IsTeamAKnown,
		IsTeamBKnown:  model.IsTeamBKnown,
		BestOf:        model.BestOf,
	}
}

func (MatchMapper) Model(proto *pb.Match) model.MatchModel {
	return model.MatchModel{
		ID:            int(proto.ID),
		FaceitId:      proto.FaceitId,
		TeamAName:     proto.TeamAName,
		TeamBName:     proto.TeamBName,
		ScoreTeamA:    proto.ScoreTeamA,
		ScoreTeamB:    proto.ScoreTeamB,
		TeamAFaceitId: proto.TeamAFaceitId,
		TeamBFaceitId: proto.TeamBFaceitId,
		IsTeamAKnown:  proto.IsTeamAKnown,
		IsTeamBKnown:  proto.IsTeamBKnown,
		BestOf:        proto.BestOf,
	}
}
