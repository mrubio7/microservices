package mapper

import (
	"ibercs/internal/model"
	matches_mapper "ibercs/pkg/mapper/matches"
	players_mapper "ibercs/pkg/mapper/players"
	pb_matches "ibercs/proto/matches"
	pb_players "ibercs/proto/players"
)

func RegisterMappers() {
	// Mapper para MatchModel -> Proto
	Register(Mapper[model.MatchModel, *pb_matches.Match]{
		From: matches_mapper.MatchMapper{}.Proto,
		To:   matches_mapper.MatchMapper{}.Model,
	})

	// Mapper para Proto -> MatchModel
	Register(Mapper[*pb_matches.Match, model.MatchModel]{
		From: matches_mapper.MatchMapper{}.Model,
		To:   matches_mapper.MatchMapper{}.Proto,
	})

	// Mapper para PlayerModel -> Proto
	Register(Mapper[model.PlayerModel, *pb_players.Player]{
		From: players_mapper.PlayerMapper{}.Proto,
		To:   players_mapper.PlayerMapper{}.Model,
	})

	// Mapper para Proto -> PlayerModel
	Register(Mapper[*pb_players.Player, model.PlayerModel]{
		From: players_mapper.PlayerMapper{}.Model,
		To:   players_mapper.PlayerMapper{}.Proto,
	})

	// Mapper para ProminentPlayerModel -> Proto
	Register(Mapper[model.PlayerProminentModel, *pb_players.ProminentPlayer]{
		From: players_mapper.PlayerProminentMapper{}.Proto,
		To:   players_mapper.PlayerProminentMapper{}.Model,
	})

	// Mapper para Proto -> ProminentPlayerModel
	Register(Mapper[*pb_players.ProminentPlayer, model.PlayerProminentModel]{
		From: players_mapper.PlayerProminentMapper{}.Model,
		To:   players_mapper.PlayerProminentMapper{}.Proto,
	})

}
