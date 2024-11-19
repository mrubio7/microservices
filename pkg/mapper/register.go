package mapper

import (
	"ibercs/internal/model"
	matches_mapper "ibercs/pkg/mapper/matches"
	pb "ibercs/proto/matches"
)

func RegisterMappers() {
	// Mapper para MatchModel -> Proto
	Register(Mapper[model.MatchModel, *pb.Match]{
		From: matches_mapper.MatchMapper{}.Proto,
		To:   matches_mapper.MatchMapper{}.Model,
	})

	// Mapper para Proto -> MatchModel
	Register(Mapper[*pb.Match, model.MatchModel]{
		From: matches_mapper.MatchMapper{}.Model,
		To:   matches_mapper.MatchMapper{}.Proto,
	})

}
