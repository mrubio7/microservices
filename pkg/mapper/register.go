package mapper

import (
	"ibercs/internal/model"
	matches_mapper "ibercs/pkg/mapper/matches"
	players_mapper "ibercs/pkg/mapper/players"
	teams_mapper "ibercs/pkg/mapper/teams"
	users_mapper "ibercs/pkg/mapper/users"
	pb_matches "ibercs/proto/matches"
	pb_players "ibercs/proto/players"
	pb_teams "ibercs/proto/teams"
	pb_users "ibercs/proto/users"
)

func RegisterMappers() {
	// Match
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

	// Player
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

	// Mapper para LookingForTeamModel -> Proto
	Register(Mapper[model.LookingForTeamModel, *pb_players.PlayerLookingForTeam]{
		From: players_mapper.PlayerLookingForTeamMapper{}.Proto,
		To:   players_mapper.PlayerLookingForTeamMapper{}.Model,
	})
	// Mapper para Proto -> LookingForTeamModel
	Register(Mapper[*pb_players.PlayerLookingForTeam, model.LookingForTeamModel]{
		From: players_mapper.PlayerLookingForTeamMapper{}.Model,
		To:   players_mapper.PlayerLookingForTeamMapper{}.Proto,
	})

	// Mapper para CreatePlayerLookingForTeamRequest -> LookingForTeamModel
	Register(Mapper[*pb_players.CreatePlayerLookingForTeamRequest, model.LookingForTeamModel]{
		From: players_mapper.CreatePlayerLookingForTeamMapper{}.Model,
		To:   players_mapper.CreatePlayerLookingForTeamMapper{}.Proto,
	})

	// Team
	// Mapper para TeamModel -> Proto
	Register(Mapper[model.TeamModel, *pb_teams.Team]{
		From: teams_mapper.TeamMapper{}.Proto,
		To:   teams_mapper.TeamMapper{}.Model,
	})
	// Mapper para Proto -> TeamModel
	Register(Mapper[*pb_teams.Team, model.TeamModel]{
		From: teams_mapper.TeamMapper{}.Model,
		To:   teams_mapper.TeamMapper{}.Proto,
	})

	// Mapper para User -> Proto
	Register(Mapper[model.UserModel, *pb_users.User]{
		From: users_mapper.UserMapper{}.Proto,
		To:   users_mapper.UserMapper{}.Model,
	})
	// Mapper para Proto -> User
	Register(Mapper[*pb_users.User, model.UserModel]{
		From: users_mapper.UserMapper{}.Model,
		To:   users_mapper.UserMapper{}.Proto,
	})

}
