package microservice_players

import (
	"context"
	"errors"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	pb "ibercs/proto/teams"
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
		Id:        t.ID,
		FaceitId:  t.FaceitId,
		Name:      t.Name,
		Nickname:  t.Nickname,
		Avatar:    t.Avatar,
		Active:    t.Active,
		PlayersId: t.PlayersId,
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
		Id:        t.ID,
		FaceitId:  t.FaceitId,
		Name:      t.Name,
		Nickname:  t.Nickname,
		Avatar:    t.Avatar,
		Active:    t.Active,
		PlayersId: t.PlayersId,
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
	t := s.TeamsService.GetTeamByNickname(teamRequest.FaceitId)
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
		Id:        t.ID,
		FaceitId:  t.FaceitId,
		Name:      t.Name,
		Nickname:  t.Nickname,
		Avatar:    t.Avatar,
		Active:    t.Active,
		PlayersId: t.PlayersId,
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
			Id:        t.ID,
			FaceitId:  t.FaceitId,
			Name:      t.Name,
			Nickname:  t.Nickname,
			Avatar:    t.Avatar,
			Active:    t.Active,
			PlayersId: t.PlayersId,
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
