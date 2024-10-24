package microservice_players

import (
	"context"
	"errors"
	"ibercs/pkg/config"
	"ibercs/pkg/consts"
	"ibercs/pkg/database"
	"ibercs/pkg/faceit"
	"ibercs/pkg/logger"
	"ibercs/pkg/service"
	pb "ibercs/proto/players"
	"time"
)

type Server struct {
	pb.UnimplementedPlayerServiceServer
	PlayersService *service.Players
	FaceitService  *faceit.FaceitClient
}

func New() *Server {
	cfg, err := config.Load()
	if err != nil {
		logger.Error("Unable to create grpc player server")
		return nil
	}
	db := database.New(cfg.Database)
	playerService := service.NewPlayersService(db)
	faceit := faceit.New(cfg.FaceitApiToken)

	return &Server{
		PlayersService: playerService,
		FaceitService:  faceit,
	}
}

func (s *Server) GetPlayer(ctx context.Context, playerReq *pb.GetPlayerRequest) (*pb.PlayerList, error) {
	playersRes := make([]*pb.Player, len(playerReq.FaceitId))

	for i, p := range playerReq.FaceitId {
		playerUpdated := s.FaceitService.GetPlayerAverageDetails(p, consts.LAST_MATCHES_NUMBER)
		err := s.PlayersService.UpdatePlayer(*playerUpdated)
		if err != nil {
			logger.Warning("Error updating player: %s", err.Error())
		}

		p := s.PlayersService.GetPlayer(p)
		if p == nil {
			return nil, errors.New("unable to get player")
		}

		player := &pb.Player{
			Id:       p.ID,
			Nickname: p.Nickname,
			FaceitId: p.FaceitId,
			SteamId:  p.SteamId,
			Avatar:   p.Avatar,
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

		playersRes[i] = player
	}

	return &pb.PlayerList{Players: playersRes}, nil
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
			Avatar:   p.Avatar,
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

	if latestWeek == nil || latestWeek.Year < int16(currentYear) || latestWeek.Week < int16(currentWeek) {
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

	return &pb.ProminentPlayerList{
		Week:    int32(currentWeek),
		Year:    int32(currentYear),
		Players: players,
	}, nil
}

func (s *Server) NewPlayer(ctx context.Context, req *pb.NewPlayerRequest) (*pb.Player, error) {
	player := s.FaceitService.GetPlayerAverageDetails(req.FaceitId, consts.LAST_MATCHES_NUMBER)

	err := s.PlayersService.UpdatePlayer(*player)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	res := &pb.Player{
		Id:       player.ID,
		Nickname: player.Nickname,
		FaceitId: player.FaceitId,
		SteamId:  player.SteamId,
		Avatar:   player.Avatar,
		Stats: &pb.PlayerStats{
			PlayerId:               player.Stats.ID,
			KdRatio:                player.Stats.KdRatio,
			KrRatio:                player.Stats.KrRatio,
			KillsAverage:           player.Stats.KillsAverage,
			DeathsAverage:          player.Stats.DeathsAverage,
			HeadshotPercentAverage: player.Stats.HeadshotPercentAverage,
			MVPAverage:             player.Stats.MVPAverage,
			AssistAverage:          player.Stats.AssistAverage,
			Elo:                    player.Stats.Elo,
		},
	}

	return res, nil
}
