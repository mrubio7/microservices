package microservice_players

import (
	"context"
	"errors"
	"ibercs/pkg/config"
	"ibercs/pkg/database"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	pb "ibercs/proto/players"
	"time"
)

type Server struct {
	pb.UnimplementedPlayerServiceServer
	PlayersService *service.Players
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)
	playerService := service.NewPlayersService(db)

	return &Server{
		PlayersService: playerService,
	}
}

func (s *Server) GetPlayers(context.Context, *pb.Empty) (*pb.PlayerList, error) {
	playerModels := s.PlayersService.GetPlayers()
	if playerModels == nil {
		return nil, errors.New("unable to get any player")
	}

	var players []*pb.Player
	for _, p := range playerModels {
		player := &pb.Player{
			Id:       p.ID,
			Nickname: p.Nickname,
			FaceitId: p.FaceitId,
			SteamId:  p.SteamId,
			Stats: &pb.PlayerStats{
				PlayerId:               p.Stats.ID,
				KdRatio:                p.Stats.KdRatio,
				KrRatio:                p.Stats.KrRatio,
				KillsAverage:           p.Stats.KillsAverage,
				DeathsAverage:          p.Stats.DeathsAverage,
				HeadshotPercentAverage: p.Stats.HeadshotPercentAverage,
				MVPAverage:             p.Stats.MVPAverage,
				AssistAverage:          p.Stats.AssistAverage,
				Elo:                    p.Stats.Elo,
			},
		}
		players = append(players, player)
	}

	return &pb.PlayerList{Players: players}, nil
}

func (s *Server) GetProminentPlayers(ctx context.Context, req *pb.Empty) (*pb.ProminentPlayerList, error) {
	latestWeek := s.PlayersService.GetProminentPlayers()
	currentYear, currentWeek := time.Now().ISOWeek()

	if latestWeek == nil || latestWeek.Year < int16(currentYear) || latestWeek.Week < int16(currentWeek-1) {
		logger.Info("Generating a new prominent week since no valid week was found or the week is outdated.")
		latestWeek = s.PlayersService.GetNewProminentPlayers()
	}

	if latestWeek == nil {
		err := errors.New("failed to retrieve or generate prominent players")
		logger.Error(err.Error())
		return nil, err
	}

	var players []*pb.ProminentPlayer
	for _, p := range latestWeek.Players {
		player := &pb.ProminentPlayer{
			Id:       p.ID,
			Nickname: p.Nickname,
			FaceitId: p.FaceitId,
			SteamId:  p.SteamId,
			Avatar:   p.Avatar,
			Score:    p.Score,
		}
		players = append(players, player)
	}

	return &pb.ProminentPlayerList{Players: players}, nil
}

func (s *Server) NewPlayer(context.Context, *pb.NewPlayerRequest) (*pb.PlayerResponse, error) {
	return nil, nil
}
