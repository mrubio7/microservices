package microservice_players

import (
	"context"
	"errors"
	"ibercs/internal/model"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	pb "ibercs/proto/teams"
	"reflect"
)

type Server struct {
	pb.UnimplementedTeamServiceServer
	TeamsService  *service.Teams
	FaceitService *faceit.FaceitClient
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)
	teamsService := service.NewTeamsService(db)
	faceitClient := faceit.New(cfg.FaceitApiToken)

	return &Server{
		TeamsService:  teamsService,
		FaceitService: faceitClient,
	}
}

func (s *Server) NewTeam(ctx context.Context, teamRequest *pb.NewTeamRequest) (*pb.Team, error) {
	team := s.FaceitService.GetTeamById(teamRequest.FaceitId)
	if team == nil {
		err := errors.New("team is empty")
		logger.Error(err.Error())
		return nil, err
	}

	t := s.TeamsService.NewTeam(*team)

	mapStats := make(map[string]*pb.TeamMapStats, len(t.Stats.MapStats))
	for _, m := range t.Stats.MapStats {
		mapStats[m.MapName] = &pb.TeamMapStats{
			MapName: m.MapName,
			Winrate: m.WinRate,
			Matches: m.Matches,
		}
	}

	pbTeam := &pb.Team{
		Id:          t.ID,
		FaceitId:    t.FaceitId,
		Name:        t.Name,
		Nickname:    t.Nickname,
		Avatar:      t.Avatar,
		Active:      t.Active,
		PlayersId:   t.PlayersId,
		Twitter:     t.Twitter,
		Instagram:   t.Instagram,
		Web:         t.Web,
		Tournaments: t.Tournaments, Stats: &pb.TeamStats{
			TotalMatches:  t.Stats.TotalMatches,
			Wins:          t.Stats.Wins,
			Winrate:       t.Stats.Winrate,
			RecentResults: t.Stats.RecentResults,
			MapStats:      mapStats,
		},
	}

	return pbTeam, nil
}

func (s *Server) GetTeamById(ctx context.Context, teamRequest *pb.NewTeamRequest) (*pb.Team, error) {
	t := s.TeamsService.GetTeam(teamRequest.FaceitId)
	if t == nil {
		err := errors.New("team not found")
		logger.Error(err.Error())
		return nil, err
	}

	mapStats := make(map[string]*pb.TeamMapStats, len(t.Stats.MapStats))
	for _, m := range t.Stats.MapStats {
		mapStats[m.MapName] = &pb.TeamMapStats{
			MapName: m.MapName,
			Winrate: m.WinRate,
			Matches: m.Matches,
		}
	}

	pbTeam := &pb.Team{
		Id:          t.ID,
		FaceitId:    t.FaceitId,
		Name:        t.Name,
		Nickname:    t.Nickname,
		Avatar:      t.Avatar,
		Active:      t.Active,
		PlayersId:   t.PlayersId,
		Twitter:     t.Twitter,
		Instagram:   t.Instagram,
		Web:         t.Web,
		Tournaments: t.Tournaments,
		Stats: &pb.TeamStats{
			TotalMatches:  t.Stats.TotalMatches,
			Wins:          t.Stats.Wins,
			Winrate:       t.Stats.Winrate,
			RecentResults: t.Stats.RecentResults,
			MapStats:      mapStats,
		},
	}

	return pbTeam, nil
}

func (s *Server) GetTeamByNickname(ctx context.Context, teamRequest *pb.NewTeamRequest) (*pb.Team, error) {
	var t *model.TeamModel

	t = s.TeamsService.GetTeamByNickname(teamRequest.FaceitId)
	if t == nil {
		err := errors.New("team not found")
		logger.Error(err.Error())
		return nil, err
	}

	teamUpdated := s.FaceitService.GetTeamById(t.FaceitId)
	if teamUpdated == nil {
		err := errors.New("team is empty")
		logger.Error(err.Error())
	}

	teamUpdated.ID = t.ID
	if !reflect.DeepEqual(t, teamUpdated) {
		t = teamUpdated
		res := s.TeamsService.UpdateTeam(*teamUpdated)
		if res == nil {
			logger.Error("Unable to update the team")
		}
	}

	mapStats := make(map[string]*pb.TeamMapStats, len(t.Stats.MapStats))
	for _, m := range t.Stats.MapStats {
		mapStats[m.MapName] = &pb.TeamMapStats{
			MapName: m.MapName,
			Winrate: m.WinRate,
			Matches: m.Matches,
		}
	}

	pbTeam := &pb.Team{
		Id:          t.ID,
		FaceitId:    t.FaceitId,
		Name:        t.Name,
		Nickname:    t.Nickname,
		Avatar:      t.Avatar,
		Active:      t.Active,
		PlayersId:   t.PlayersId,
		Twitter:     t.Twitter,
		Instagram:   t.Instagram,
		Web:         t.Web,
		Tournaments: t.Tournaments,
		Stats: &pb.TeamStats{
			TotalMatches:  t.Stats.TotalMatches,
			Wins:          t.Stats.Wins,
			Winrate:       t.Stats.Winrate,
			RecentResults: t.Stats.RecentResults,
			MapStats:      mapStats,
		},
	}

	return pbTeam, nil
}

func (s *Server) GetTeams(ctx context.Context, teamRequest *pb.GetTeamsRequest) (*pb.TeamList, error) {
	teams := s.TeamsService.GetAll(teamRequest.Active)
	if teams == nil {
		err := errors.New("team not found")
		logger.Error(err.Error())
		return nil, err
	}

	var pbTeams []*pb.Team

	for _, t := range teams {
		mapStats := make(map[string]*pb.TeamMapStats, len(t.Stats.MapStats))
		for _, m := range t.Stats.MapStats {
			mapStats[m.MapName] = &pb.TeamMapStats{
				MapName: m.MapName,
				Winrate: m.WinRate,
				Matches: m.Matches,
			}
		}

		pbTeams = append(pbTeams, &pb.Team{
			Id:          t.ID,
			FaceitId:    t.FaceitId,
			Name:        t.Name,
			Nickname:    t.Nickname,
			Avatar:      t.Avatar,
			Active:      t.Active,
			PlayersId:   t.PlayersId,
			Twitter:     t.Twitter,
			Instagram:   t.Instagram,
			Web:         t.Web,
			Tournaments: t.Tournaments,
			Stats: &pb.TeamStats{
				TotalMatches:  t.Stats.TotalMatches,
				Wins:          t.Stats.Wins,
				Winrate:       t.Stats.Winrate,
				RecentResults: t.Stats.RecentResults,
				MapStats:      mapStats,
			},
		})
	}

	return &pb.TeamList{Teams: pbTeams}, nil
}

func (s *Server) FindTeamByPlayerId(ctx context.Context, request *pb.NewTeamRequest) (*pb.TeamList, error) {
	teams, err := s.TeamsService.FindTeamsByPlayerId(request.FaceitId)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var pbTeams []*pb.Team

	for _, t := range teams {
		pbTeams = append(pbTeams, &pb.Team{
			Id:          t.ID,
			FaceitId:    t.FaceitId,
			Name:        t.Name,
			Nickname:    t.Nickname,
			Avatar:      t.Avatar,
			Active:      t.Active,
			PlayersId:   t.PlayersId,
			Twitter:     t.Twitter,
			Instagram:   t.Instagram,
			Web:         t.Web,
			Tournaments: t.Tournaments,
		})
	}

	return &pb.TeamList{Teams: pbTeams}, nil
}

func (s *Server) GetTeamFromFaceit(ctx context.Context, request *pb.NewTeamRequest) (*pb.Team, error) {
	t := s.FaceitService.GetTeamById(request.FaceitId)

	mapStats := make(map[string]*pb.TeamMapStats, len(t.Stats.MapStats))
	for _, m := range t.Stats.MapStats {
		mapStats[m.MapName] = &pb.TeamMapStats{
			MapName: m.MapName,
			Winrate: m.WinRate,
			Matches: m.Matches,
		}
	}

	pbTeam := &pb.Team{
		Id:          t.ID,
		FaceitId:    t.FaceitId,
		Name:        t.Name,
		Nickname:    t.Nickname,
		Avatar:      t.Avatar,
		Active:      t.Active,
		PlayersId:   t.PlayersId,
		Twitter:     t.Twitter,
		Instagram:   t.Instagram,
		Web:         t.Web,
		Tournaments: t.Tournaments,
		Stats: &pb.TeamStats{
			TotalMatches:  t.Stats.TotalMatches,
			Wins:          t.Stats.Wins,
			Winrate:       t.Stats.Winrate,
			RecentResults: t.Stats.RecentResults,
			MapStats:      mapStats,
		},
	}

	return pbTeam, nil
}

func (s *Server) GetTeamWithEseaStanding(ctx context.Context, request *pb.NewTeamRequest) (*pb.Team, error) {
	t := s.TeamsService.GetTeam(request.FaceitId)
	if t == nil {
		err := errors.New("team not found")
		logger.Error(err.Error())
		return nil, err
	}

	mapStats := make(map[string]*pb.TeamMapStats, len(t.Stats.MapStats))
	for _, m := range t.Stats.MapStats {
		mapStats[m.MapName] = &pb.TeamMapStats{
			MapName: m.MapName,
			Winrate: m.WinRate,
			Matches: m.Matches,
		}
	}

	standing, err := s.TeamsService.GetEseaStanding(t.FaceitId)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	pbTeam := &pb.Team{
		Id:          t.ID,
		FaceitId:    t.FaceitId,
		Name:        t.Name,
		Nickname:    t.Nickname,
		Avatar:      t.Avatar,
		Active:      t.Active,
		PlayersId:   t.PlayersId,
		Twitter:     t.Twitter,
		Instagram:   t.Instagram,
		Web:         t.Web,
		Tournaments: t.Tournaments,
		Stats: &pb.TeamStats{
			TotalMatches:  t.Stats.TotalMatches,
			Wins:          t.Stats.Wins,
			Winrate:       t.Stats.Winrate,
			RecentResults: t.Stats.RecentResults,
			MapStats:      mapStats,
		},
		Standing: &pb.Standing{
			FaceitId:       standing.FaceitId,
			TournamentId:   standing.TournamentId,
			IsDisqualified: standing.IsDisqualified,
			RankStart:      int32(standing.RankStart),
			RankEnd:        int32(standing.RankEnd),
			MatchesPlayed:  int32(standing.MatchesPlayed),
			MatchesWon:     int32(standing.MatchesWon),
			MatchesLost:    int32(standing.MatchesLost),
			MatchesTied:    int32(standing.MatchesTied),
			Points:         int32(standing.Points),
			BuchholzScore:  int32(standing.BuchholzScore),
		},
	}

	return pbTeam, nil
}
