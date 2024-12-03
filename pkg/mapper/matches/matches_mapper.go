package matches_mapper

import (
	"ibercs/internal/model"
	teams_mapper "ibercs/pkg/mapper/teams"
	pb "ibercs/proto/matches"
	"time"
)

type MatchMapper struct{}

func (MatchMapper) Proto(model model.MatchModel, _ ...interface{}) *pb.Match {
	return &pb.Match{
		ID:                 int32(model.ID),
		FaceitId:           model.FaceitId,
		TeamAName:          model.TeamAName,
		TeamBName:          model.TeamBName,
		ScoreTeamA:         model.ScoreTeamA,
		ScoreTeamB:         model.ScoreTeamB,
		TeamAFaceitId:      model.TeamAFaceitId,
		TeamBFaceitId:      model.TeamBFaceitId,
		IsTeamAKnown:       model.IsTeamAKnown,
		IsTeamBKnown:       model.IsTeamBKnown,
		BestOf:             model.BestOf,
		Timestamp:          model.Timestamp.Unix(),
		Streams:            model.Streams,
		TournamentName:     model.TournamentName,
		TournamentFaceitId: model.TournamentFaceitId,
		TeamA:              teams_mapper.TeamMapper{}.Proto(model.TeamA),
		TeamB:              teams_mapper.TeamMapper{}.Proto(model.TeamB),
	}
}

func (MatchMapper) Model(proto *pb.Match, _ ...interface{}) model.MatchModel {
	return model.MatchModel{
		ID:                 int(proto.ID),
		FaceitId:           proto.FaceitId,
		TeamAName:          proto.TeamAName,
		TeamBName:          proto.TeamBName,
		ScoreTeamA:         proto.ScoreTeamA,
		ScoreTeamB:         proto.ScoreTeamB,
		TeamAFaceitId:      proto.TeamAFaceitId,
		TeamBFaceitId:      proto.TeamBFaceitId,
		IsTeamAKnown:       proto.IsTeamAKnown,
		IsTeamBKnown:       proto.IsTeamBKnown,
		BestOf:             proto.BestOf,
		Timestamp:          time.Unix(proto.Timestamp, 0),
		Streams:            proto.Streams,
		TournamentName:     proto.TournamentName,
		TournamentFaceitId: proto.TournamentFaceitId,
		TeamA:              teams_mapper.TeamMapper{}.Model(proto.TeamA),
		TeamB:              teams_mapper.TeamMapper{}.Model(proto.TeamB),
	}
}
