package mapper

import (
	"ibercs/internal/model"
	esea_mapper "ibercs/pkg/mapper/esea"
	matches_mapper "ibercs/pkg/mapper/matches"
	players_mapper "ibercs/pkg/mapper/players"
	teams_mapper "ibercs/pkg/mapper/teams"
	tournaments_mapper "ibercs/pkg/mapper/tournaments"
	users_mapper "ibercs/pkg/mapper/users"
	pb_matches "ibercs/proto/matches"
	pb_players "ibercs/proto/players"
	pb_teams "ibercs/proto/teams"
	pb_tournaments "ibercs/proto/tournaments"
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

	// Mapper para TeamRankModel -> Proto
	Register(Mapper[model.TeamRankModel, *pb_teams.TeamRank]{
		From: teams_mapper.TeamRankMapper{}.Proto,
		To:   teams_mapper.TeamRankMapper{}.Model,
	})
	// Mapper para Proto -> TeamRankModel
	Register(Mapper[*pb_teams.TeamRank, model.TeamRankModel]{
		From: teams_mapper.TeamRankMapper{}.Model,
		To:   teams_mapper.TeamRankMapper{}.Proto,
	})

	// User
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

	// Session
	// Mapper para Session -> Proto
	Register(Mapper[model.UserSessionModel, *pb_users.SessionResponse]{
		From: users_mapper.SessionMapper{}.Proto,
		To:   users_mapper.SessionMapper{}.Model,
	})
	// Mapper para Proto -> Session
	Register(Mapper[*pb_users.SessionResponse, model.UserSessionModel]{
		From: users_mapper.SessionMapper{}.Model,
		To:   users_mapper.SessionMapper{}.Proto,
	})

	// Tournament
	// Mapper para TournamentModel -> Proto
	Register(Mapper[model.TournamentModel, *pb_tournaments.Tournament]{
		From: tournaments_mapper.TournamentMapper{}.Proto,
		To:   tournaments_mapper.TournamentMapper{}.Model,
	})
	// Mapper para Proto -> TournamentModel
	Register(Mapper[*pb_tournaments.Tournament, model.TournamentModel]{
		From: tournaments_mapper.TournamentMapper{}.Model,
		To:   tournaments_mapper.TournamentMapper{}.Proto,
	})

	// Esea
	// Maper para EseaLeagueModel -> Proto
	Register(Mapper[model.EseaLeagueModel, *pb_tournaments.Esea]{
		From: esea_mapper.EseaMapper{}.Proto,
		To:   esea_mapper.EseaMapper{}.Model,
	})
	// Mapper para Proto -> EseaLeagueModel
	Register(Mapper[*pb_tournaments.Esea, model.EseaLeagueModel]{
		From: esea_mapper.EseaMapper{}.Model,
		To:   esea_mapper.EseaMapper{}.Proto,
	})

	// Mapper para EseaDivisionModel -> Proto
	Register(Mapper[model.EseaDivisionModel, *pb_tournaments.EseaDivision]{
		From: esea_mapper.EseaDivisionMapper{}.Proto,
		To:   esea_mapper.EseaDivisionMapper{}.Model,
	})
	// Mapper para Proto -> EseaDivisionModel
	Register(Mapper[*pb_tournaments.EseaDivision, model.EseaDivisionModel]{
		From: esea_mapper.EseaDivisionMapper{}.Model,
		To:   esea_mapper.EseaDivisionMapper{}.Proto,
	})

	// Mapper para EseaStandingModel -> Proto
	Register(Mapper[model.EseaStandingModel, *pb_tournaments.EseaStanding]{
		From: esea_mapper.EseaStandingMapper{}.Proto,
		To:   esea_mapper.EseaStandingMapper{}.Model,
	})
	// Mapper para Proto -> EseaStandingModel
	Register(Mapper[*pb_tournaments.EseaStanding, model.EseaStandingModel]{
		From: esea_mapper.EseaStandingMapper{}.Model,
		To:   esea_mapper.EseaStandingMapper{}.Proto,
	})

	// Organizer
	// Mapper para OrganizerModel -> Proto
	Register(Mapper[model.OrganizerModel, *pb_tournaments.Organizer]{
		From: tournaments_mapper.OrganizerMapper{}.Proto,
		To:   tournaments_mapper.OrganizerMapper{}.Model,
	})
	// Mapper para Proto -> OrganizerModel
	Register(Mapper[*pb_tournaments.Organizer, model.OrganizerModel]{
		From: tournaments_mapper.OrganizerMapper{}.Model,
		To:   tournaments_mapper.OrganizerMapper{}.Proto,
	})

}
